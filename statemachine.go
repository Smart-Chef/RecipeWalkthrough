package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"recipe-walkthrough/models"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
)

type RecipeInfo struct {
	ID          null.Int       `json:"recipe_id"`
	JobIDs      []int64          `json:"job_ids"`
	CurrentStep *models.Step   `json:"current_step"`
	PrevStep    *models.Step   `json:"prev_step"`
	NextStep    *models.Step   `json:"next_step"`
	TotalSteps  int            `json:"total_steps"`
	recipe      *models.Recipe `json:"-"`
}

type JobPayload struct {
	Service       string        `json:"service"`
	ActionKey     interface{}   `json:"action_key"`
	ActionParams  interface{}   `json:"action_params"`
	TriggerKeys   []string      `json:"trigger_keys"`
	TriggerParams []interface{} `json:"trigger_params"`
}

type JobResponse struct {
	Status string `json:"status"`
	JobID  int64    `json:"job_id"`
}

type ClearJob struct {
	ID int64 `json:"id"`
}

type ClearJobResponse struct {
	ID int64 `json:"id"`
}

// Any special setup for triggers/actions before sending to trigger-queue
// i.e dynamic triggers or actions
func (j *JobPayload) specialJobSetup() error {
	for i, trigger := range j.TriggerKeys {
		// Convert the delay to a timestamp
		if trigger == "timer" {
			loc, _ := time.LoadLocation("UTC")
			j.TriggerParams[i] = time.Now().In(loc).Add(time.Second * time.Duration(j.TriggerParams[i].(int)))
		}
	}
	return nil
}

// SendJob to the trigger queue
// returns JobID, error
func (j *JobPayload) SendJob() (int64, error) {
	j.specialJobSetup()
	url := os.Getenv("TRIGGER_QUEUE_API") + "/add"
	payload, _ := json.Marshal(j)

	log.Infof("Sending to trigger-queue: %s", payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return -1, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var resp JobResponse
	err = decoder.Decode(&resp)
	if err != nil {
		return -1, err
	}
	return resp.JobID, nil
}

func (j *ClearJob) clear() (bool, error) {
	url := os.Getenv("TRIGGER_QUEUE_API") + "/delete/walk-through/" + strconv.FormatInt(j.ID, 10)
	req, _ := http.NewRequest("DELETE", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var resp ClearJobResponse
	err := decoder.Decode(&resp)
	if err != nil {
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil
}

// incrementNSteps forwards
// returns recipeCompleted, error
func (r *RecipeInfo) incrementNSteps(n int) (bool, error) {
	if r.CurrentStep == nil {
		return false, errors.New("cannot increment step: no recipe currently active")
	}
	log.Info("Incrementing " + strconv.Itoa(n) + " step(s)")
	return r.initStep(int(r.CurrentStep.StepNumber.Int64 + int64(n-1)))
}

// Setup a newRecipe given a Recipe ID
func (r *RecipeInfo) newRecipe(id int) error {
	// Get recipe from the database
	recipe, err := new(models.Recipe).GetByID(database, queries, id)
	if err != nil {
		return err
	}

	// Setup Recipe info
	r.ID = recipe.ID
	r.TotalSteps = len(recipe.Steps)
	r.CurrentStep = recipe.Steps[0]
	r.PrevStep = nil
	r.NextStep = nil
	r.JobIDs = make([]int64, 0)

	if r.TotalSteps > 1 {
		r.NextStep = recipe.Steps[1]
	}

	r.recipe = recipe
	done, err := r.initStep(0)

	if err != nil {
		return err
	}

	if done {
		log.Warn("Just setup a newRecipe that is already done")
	}
	//err = r.SayCurrentStep()
	return err
}

// clear a Recipe from the state machine and clean up any triggers
// return (successful cleaning up triggers)
func (r *RecipeInfo) clear() error {
	r.TotalSteps = 0
	r.CurrentStep = nil
	r.NextStep = nil
	r.PrevStep = nil

	log.Info("Clearing Jobs")
	_, err := r.clearJobs()
	r.JobIDs = make([]int64, 0)
	return err
}

func (r *RecipeInfo) clearJobs() ([]int64, error) {
	// Clear the jobs from the trigger queue
	nonSuccessfulIDs := make([]int64, 0)
	erred := false
	for _, jobID := range CurrentRecipe.JobIDs {
		j := &ClearJob{ID: jobID}
		success, err := j.clear()
		if err != nil || !success {
			erred = true
			nonSuccessfulIDs = append(nonSuccessfulIDs, jobID)
			if err != nil {
				log.Error(err)
			}
		}
	}
	if erred {
		msg := "could not clear jobs ["
		for i, id := range nonSuccessfulIDs {
			msg += strconv.FormatInt(id, 10)
			if i != len(nonSuccessfulIDs)-1 {
				msg += ", "
			} else {
				msg += "]"
			}
		}
		return nonSuccessfulIDs, errors.New(msg)
	}
	return nonSuccessfulIDs, nil
}

// Setup a new step in the recipe
// returns recipeCompleted, error
func (r *RecipeInfo) initStep(step int) (bool, error) {
	// Check if the recipe is done
	if step == r.TotalSteps || r.ID.ValueOrZero() == -1 {
		return true, r.clear()
	}

	// Check if we are given a valid step in our recipe
	if step < 0 || step > r.TotalSteps {
		log.Errorf("invalid step (%d) when there are only (%d) steps", step, r.TotalSteps)
		return false, errors.New("invalid step")
	}

	// Clear past step jobs
	var err error
	_, err = CurrentRecipe.clearJobs()
	CurrentRecipe.JobIDs = make([]int64, 0)
	if err != nil {
		log.Error(err.Error())
	}

	// Setup triggers/jobs
	jobs := make([]*JobPayload, 0)

	// Construct each JobPayload from each TriggerGroup
	for _, triggerGroup := range r.recipe.Steps[step].TriggerGroups {
		var job JobPayload

		if triggerGroup.ActionParams.Valid {
			job.ActionParams = triggerGroup.ActionParams
		}

		if triggerGroup.ActionKey.Valid {
			job.ActionKey = triggerGroup.ActionKey
		}

		if triggerGroup.Service.Valid {
			job.Service = triggerGroup.Service.String
		}

		// Loop through all the triggers in the trigger group
		for _, trigger := range triggerGroup.Triggers {
			l, err := strconv.Atoi(trigger.TriggerParams.String)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			job.TriggerParams = append(job.TriggerParams, l)
			job.TriggerKeys = append(job.TriggerKeys, trigger.TriggerType.Key.String)
		}
		jobs = append(jobs, &job)
	}

	// Send the jobs to the trigger-queue
	for _, j := range jobs {
		var id int64
		var err error
		id, err = j.SendJob()
		if err != nil {
			log.Error("error sending job %+v", j)
			log.Error(err.Error())
		} else {
			r.JobIDs = append(r.JobIDs, id)
		}
	}

	// Update self
	if step > 0 {
		r.PrevStep = r.recipe.Steps[step-1]
	} else {
		r.PrevStep = nil
	}

	r.CurrentStep = r.recipe.Steps[step]

	if step+1 < r.TotalSteps {
		r.NextStep = r.recipe.Steps[step+1]
	} else {
		r.NextStep = nil
	}
	return false, nil
}

// Send the current step info to the NLP to say to the user
func (r *RecipeInfo) SayCurrentStep() error {
	url := os.Getenv("NLP_API") + "/send_message/" + r.CurrentStep.Data.String
	log.Infof("Sending to NLP: %s", r.CurrentStep.Data.String)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	log.Infof("NLP Response: %s", body)
	return err
}
