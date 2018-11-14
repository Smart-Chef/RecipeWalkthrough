package main

import (
	"encoding/json"
	"net/http"
	"recipe-walkthrough/models"
	"strconv"

	"github.com/gorilla/mux"

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

var GetRecipes = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string           `json:"status"`
		Data   []*models.Recipe `json:"recipes"`
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   new(models.Recipe).GetAll(database, queries),
	})
})

var GetRecipeByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string         `json:"status"`
		Data   *models.Recipe `json:"data"`
	}

	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	recipe, _ := new(models.Recipe).GetByID(database, queries, id)

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   recipe,
	})
})

var GetIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   new(models.Ingredient).GetAll(database, queries),
	})
})

var GetRecipeIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   new(models.Ingredient).GetAllByRecipe(database, queries, id),
	})
})

var GetRecipeStepIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	stepNumber := mux.Vars(r)["step_number"]

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   new(models.Ingredient).GetAllByRecipeStep(database, queries, id, stepNumber),
	})
})

// Server testing controllers
var Pong = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Info("Pong!")
	w.Write([]byte("Pong!\n"))
})
