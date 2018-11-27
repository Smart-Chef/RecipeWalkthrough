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
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&struct {
		Status string      `json:"status"`
		Msg    interface{} `json:"msg"`
	}{INVALIDSTATUS, msgAndArgs})
}

// NewRecipe initializes recipe by RECIPE_ID
// Payload: { "id": RECIPE_ID }
// Response: { "status": msg }
var NewRecipe = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	}

	req := &Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error("Invalid Request in NewRecipe")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	// Setup the newRecipe
	err = CurrentRecipe.newRecipe(req.ID)

	if err != nil {
		log.Error("Error setting up new recipe")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Msg:    CurrentRecipe.CurrentStep.Data.String,
	})
})

// Get all info about the CurrentStep, NextStep, and PrevStep
var GetCurrentStep = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(CurrentRecipe)
	if err != nil {
		log.Error("Error marshalling CurrentRecipe")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
	}
})

// Increment N steps forward
// Payload: {"increment_steps": 1, "send_to_nlp": DEFAULT_FALSE}
var GotToNStep = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Request struct {
		IncrementSteps int  `json:"increment_steps"`
		SendToNLP      bool `json:"send_to_nlp"`
	}
	type Response struct {
		Status     string `json:"status"`
		RecipeDone bool   `json:"recipe_done"`
		Msg        string `json:"msg"`
		StepInfo   string `json:"step_info"`
	}

	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Error("Invalid Request ")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	log.Warn(req.SendToNLP)

	recipeDone, err := CurrentRecipe.incrementNSteps(req.IncrementSteps)

	if err != nil {
		log.Errorf("Error incrementing %d steps", req.IncrementSteps)
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	if req.SendToNLP {
		err = CurrentRecipe.SayCurrentStep()
		if err != nil {
			log.Error(err.Error())
		}
	}

	var msg string
	if recipeDone {
		log.Infof("Recipe(%d) is now completed", CurrentRecipe.ID.Int64)
		msg = "Recipe completed"
	} else {
		msg = "Go forward " + strconv.Itoa(req.IncrementSteps) + " step(s)"
	}

	stepInfo := ""

	if CurrentRecipe.CurrentStep != nil {
		stepInfo = CurrentRecipe.CurrentStep.Data.String
	}
	json.NewEncoder(w).Encode(&Response{
		Status:     "success",
		RecipeDone: recipeDone,
		Msg:        msg,
		StepInfo:   stepInfo,
	})
})

var GetRecipes = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string           `json:"status"`
		Data   []*models.Recipe `json:"recipes"`
	}

	recipes, err := new(models.Recipe).GetAll(database, queries)
	if err != nil {
		log.Error("Error getting all recipes")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   recipes,
	})
})

var GetRecipeByID = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string         `json:"status"`
		Data   *models.Recipe `json:"data"`
	}

	id, ok := mux.Vars(r)["id"]
	if !ok {
		msg := "No {ID} sent in route"
		log.Error(msg)
		handleBadRequest(w, msg)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	recipe, err := new(models.Recipe).GetByID(database, queries, idInt)

	if err != nil {
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   recipe,
	})
})

var GetIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}

	ingredients, err := new(models.Ingredient).GetAll(database, queries)
	if err != nil {
		log.Error("Error getting all ingredients")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   ingredients,
	})
})

var GetRecipeIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("Error parsing ID from request route")
		handleBadRequest(w, "Error parsing ID from request route")
		return
	}

	ingredients, err := new(models.Ingredient).GetAllByRecipe(database, queries, id)
	if err != nil {
		log.Errorf("Error getting all ingredients for recipe (%d)", CurrentRecipe.ID)
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   ingredients,
	})
})

var GetRecipeStepIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string               `json:"status"`
		Data   []*models.Ingredient `json:"ingredients"`
	}
	w.Header().Set("Content-Type", "application/json")
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("Error parsing ID from request route")
		handleBadRequest(w, "Error parsing ID from request route")
		return
	}

	stepNumber, ok := mux.Vars(r)["step_number"]
	if !ok {
		log.Error("Error parsing step_number from request route")
		handleBadRequest(w, "Error parsing step_number from request route")
		return
	}

	ingredients, err := new(models.Ingredient).GetAllByRecipeStep(database, queries, id, stepNumber)
	if err != nil {
		log.Error("Error ingredients by step")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   ingredients,
	})
})

var GetUtensils = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string            `json:"status"`
		Data   []*models.Utensil `json:"utensils"`
	}
	utensils, err := new(models.Utensil).GetAll(database, queries)
	if err != nil {
		log.Error("Error getting all utensils")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   utensils,
	})
})

var GetRecipeUtensils = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string            `json:"status"`
		Data   []*models.Utensil `json:"utensils"`
	}
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("Error parsing ID from request route")
		handleBadRequest(w, "Error parsing ID from request route")
		return
	}
	utensils, err := new(models.Utensil).GetAllByRecipe(database, queries, id)
	if err != nil {
		log.Error("Error getting all utensils for recipe")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   utensils,
	})
})

var GetUtensilStepIngredients = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Response struct {
		Status string            `json:"status"`
		Data   []*models.Utensil `json:"utensils"`
	}
	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("Error parsing ID from request route")
		handleBadRequest(w, "Error parsing ID from request route")
		return
	}
	stepNumber, ok := mux.Vars(r)["step_number"]
	if !ok {
		log.Error("Error parsing step_number from request route")
		handleBadRequest(w, "Error parsing step_number from request route")
		return
	}
	utensils, err := new(models.Utensil).GetAllByRecipeStep(database, queries, id, stepNumber)
	if err != nil {
		log.Error("Error getting all utensils for recipe-step")
		log.Error(err.Error())
		handleBadRequest(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&Response{
		Status: "success",
		Data:   utensils,
	})
})

// Server testing controllers
var Pong = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Info("Pong!")
	w.Write([]byte("Pong!\n"))
})
