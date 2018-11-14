package models

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/gchaincl/dotsql"
	log "github.com/sirupsen/logrus"
)

type Recipe struct {
	ID                  int                     `json:"id"`
	Title               string                  `json:"title"`
	CreatedAt           time.Time               `json:"created_at"`
	RecipeIngredients   []*StepIngredient       `json:"recipe_ingredients,omitempty"`
	recipeIngredientMap map[int]*StepIngredient `json:"-"`
	Steps               []*Step                 `json:"steps,omitempty"`
	stepsMap            map[int]*Step           `json:"-"`
}

func (*Recipe) GetAll(db *sql.DB, dot *dotsql.DotSql) []*Recipe {
	rows, err := dot.Query(db, "get-all-recipes")
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for GetRecipes")
	}

	recipes := make([]*Recipe, 0)
	recipesMap := make(map[int]*Recipe)

	for rows.Next() {
		var recipe Recipe
		var ri StepIngredient
		var ingredient Ingredient
		var step Step
		var trigger Trigger
		var triggerType TriggerType
		var utensil Utensil

		err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.CreatedAt, &ri.ID, &ri.Quantity, &ri.Unit, &ri.CreatedAt, &ingredient.ID, &ingredient.Name, &ingredient.CreatedAt, &step.ID, &step.Data, &step.StepNumber, &step.CreatedAt, &trigger.ID, &trigger.ActionParams, &trigger.Action, &trigger.Service, &trigger.TriggerParams, &trigger.CreatedAt, &triggerType.ID, &triggerType.CreatedAt, &triggerType.Key, &triggerType.SensorType, &utensil.ID, &utensil.Name, &utensil.CreatedAt)
		if err != nil {
			log.Error(err.Error())
		} else {
			r, ok := recipesMap[recipe.ID]
			ri.Recipe = nil
			ri.Ingredient = &ingredient
			trigger.TriggerType = &triggerType

			if !ok {
				r = &recipe
				r.stepsMap = make(map[int]*Step)
				r.recipeIngredientMap = make(map[int]*StepIngredient)
				step.triggerMap = make(map[null.Int]*Trigger)
				step.utensilMap = make(map[null.Int]*Utensil)

				if trigger.ID.Valid {
					step.triggerMap[trigger.ID] = &trigger
					step.Triggers = []*Trigger{&trigger}
				}

				if utensil.ID.Valid {
					step.utensilMap[utensil.ID] = &utensil
					step.Utensils = []*Utensil{&utensil}
				}

				r.RecipeIngredients = []*StepIngredient{&ri}
				r.Steps = []*Step{&step}
				recipesMap[r.ID] = r
				recipes = append(recipes, r)
				r.stepsMap[step.ID] = &step
				r.recipeIngredientMap[ri.ID] = &ri
			} else {
				_, ok := r.recipeIngredientMap[ri.ID]
				if !ok {
					r.recipeIngredientMap[ri.ID] = &ri
					r.RecipeIngredients = append(r.RecipeIngredients, &ri)
				}

				_, ok = r.stepsMap[step.ID]
				if !ok {
					r.stepsMap[step.ID] = &step
					r.Steps = append(r.Steps, &step)
					r.stepsMap[step.ID].triggerMap = make(map[null.Int]*Trigger)
					r.stepsMap[step.ID].utensilMap = make(map[null.Int]*Utensil)
				}

				_, ok = r.stepsMap[step.ID].triggerMap[trigger.ID]
				if !ok && trigger.ID.Valid {
					r.stepsMap[step.ID].triggerMap[trigger.ID] = &trigger
					r.stepsMap[step.ID].Triggers = append(r.stepsMap[step.ID].Triggers, &trigger)
				}

				_, ok = r.stepsMap[step.ID].utensilMap[utensil.ID]

				if !ok && utensil.ID.Valid {
					r.stepsMap[step.ID].utensilMap[utensil.ID] = &utensil
					r.stepsMap[step.ID].Utensils = append(r.stepsMap[step.ID].Utensils, &utensil)
				}
			}
		}
	}
	return recipes
}

func (*Recipe) GetByID(db *sql.DB, dot *dotsql.DotSql, id string) (*Recipe, bool) {
	var recipe Recipe
	row, _ := dot.QueryRow(db, "find-one-recipe-by-id", id)
	if err := row.Scan(&recipe.ID, &recipe.Title, &recipe.CreatedAt); err != nil {
		log.Error(err.Error())
		return nil, false
	}
	return &recipe, true
}
