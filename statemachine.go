package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"recipe-walkthrough/models"
	"strconv"

	"gopkg.in/guregu/null.v3"

	log "github.com/sirupsen/logrus"
)

var CurrentRecipe *RecipeInfo

type RecipeInfo struct {
	ID          null.Int       `json:"recipe_id"`
	JobIDs      []int          `json:"job_id"`
	CurrentStep *models.Step   `json:"current_step"`
	PrevStep    *models.Step   `json:"prev_step"`
	NextStep    *models.Step   `json:"next_step"`
	TotalSteps  int            `json:"total_steps"`
	recipe      *models.Recipe `json:"-"`
}

type JobPayload struct {
	Service       string      `json:"service"`
	ActionKey     interface{} `json:"action_key"`
	ActionParams  interface{} `json:"action_params"`
	TriggerKeys   []string    `json:"trigger_keys"`
	TriggerParams []int       `json:"trigger_params"`
}

type JobResponse struct {
	Status string `json:"status"`
	JobID  int    `json:"job_id"`
}

func (j *JobPayload) SendJob() int {
	url := os.Getenv("TRIGGER_QUEUE_API") + "/add/nlp"
	payload, _ := json.Marshal(j)

	log.Info(string(payload))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error(err.Error())
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var resp JobResponse
	err = decoder.Decode(&resp)
	if err != nil {
		panic(err)
	}
	return resp.JobID
}

func (r *RecipeInfo) incrementNSteps(n int) {
	log.Info("Incrementing " + strconv.Itoa(n) + " step(s)")
	r.initStep(int(r.CurrentStep.StepNumber.Int64 + int64(n-1)))
}

func (r *RecipeInfo) newRecipe(id int) {
	// Get recipe from the database
	recipe := new(models.Recipe).GetByID(database, queries, id)

	// Setup Recipe info
	r.ID = recipe.ID
	r.TotalSteps = len(recipe.Steps)
	r.CurrentStep = recipe.Steps[0]
	r.PrevStep = nil
	r.NextStep = nil

	if r.TotalSteps > 1 {
		r.NextStep = recipe.Steps[1]
	}

	r.recipe = recipe
	r.initStep(0)
}

func (r *RecipeInfo) clear() {
	r.ID = null.IntFrom(-1)
	r.TotalSteps = 0
	r.CurrentStep = nil
	r.NextStep = nil
	r.PrevStep = nil
}

func (r *RecipeInfo) initStep(step int) (bool, error) {
	if step == r.TotalSteps || r.ID.ValueOrZero() == -1 {
		r.clear()
		return true, nil
	}

	if step < 0 || step > r.TotalSteps {
		log.Error("invalid step")
		return false, errors.New("invalid step")
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
			l, _ := strconv.Atoi(trigger.TriggerParams.String)
			job.TriggerParams = append(job.TriggerParams, l)
			job.TriggerKeys = append(job.TriggerKeys, trigger.TriggerType.Key.String)
		}
		jobs = append(jobs, &job)
	}

	// Send the jobs to the trigger-queue
	//jobIDs := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(len(jobs))

	for _, j := range jobs {
		func(l *JobPayload) {
			//defer wg.Done()
			//jobIDs <- l.SendJob()
			r.JobIDs = append(r.JobIDs, l.SendJob())
		}(j)
	}

	//wg.Wait()

	//for id := range jobIDs {
	//	r.JobIDs = append(r.JobIDs, id)
	//}

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

func init() {
	// Initialize to a non-existent step
	CurrentRecipe = new(RecipeInfo)
	//CurrentRecipe.newRecipe(1)

	//url := os.Getenv("TRIGGER_QUEUE_API") + "/add/nlp"
	//
	//payload := strings.NewReader("{\"service\":\"nlp\",\"action_key\":\"changeStep\",\"action_params\":{\"increment_steps\": 1},\"trigger_keys\":[\"temp_>\"],\"trigger_params\": [50]}")
	//
	//req, _ := http.NewRequest("POST", url, payload)
	//
	//req.Header.Add("Content-Type", "application/json")
	//
	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))
}
