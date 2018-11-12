package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var INVALIDSTATUS = "INVALID"

// handleBadRequest for the application
func handleBadRequest(w http.ResponseWriter, msgAndArgs ...interface{}) {
	//log.Warn(msgAndArgs)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&struct {
		Status string      `json:"status"`
		Msg    interface{} `json:"msg"`
	}{INVALIDSTATUS, msgAndArgs})
}

// Server testing controllers
var NewRecipe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
})

var GetCurrentStep = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CurrentStep)
})

var GotToNStep = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		IncrementSteps int `json:"increment_steps"`
	}
	type Response struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	}

	req := &Request{}
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")

	CurrentStep.incrementNSteps(req.IncrementSteps)

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Msg:    "Go forward " + strconv.Itoa(req.IncrementSteps) + " step(s)",
	})
})

var InitRecipe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
})

// Server testing controllers
var Pong = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Info("Pong!")
	w.Write([]byte("Pong!\n"))
})
