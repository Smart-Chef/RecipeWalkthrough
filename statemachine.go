package main

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

var CurrentStep *Step

type Step struct {
	RecipeID int   `json:"recipe_id"`
	StepNum  int   `json:"step_num"`
	PrevStep *Step `json:"prev_step"`
	NextStep *Step `json:"next_step"`
}

func (s *Step) incrementNSteps(n int) {
	// TODO: implement real step changing
	log.Info("Incrementing " + strconv.Itoa(n) + " step(s)")
	s.StepNum += n
}

func init() {
	// Initialize to a non-existent step
	CurrentStep = new(Step)
	CurrentStep.RecipeID = -1
	CurrentStep.StepNum = 1

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
