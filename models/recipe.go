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

// GetAll Recipes and related data
func (*Recipe) GetAll(db *sql.DB, dot *dotsql.DotSql) []*Recipe {
	// Execute the query for the data
	rows, err := dot.Query(db, "get-all-recipes")
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for GetRecipes")
	}

	return scanRecipeRows(rows)
}

func (r *Recipe) GetByID(db *sql.DB, dot *dotsql.DotSql, id int) *Recipe {
	rows, _ := dot.Query(db, "find-one-recipe-by-id", id)
	return scanRecipeRows(rows)[0]
}

func scanRecipeRows(rows *sql.Rows) []*Recipe {
	// Initialize the data structures to store the recipes
	recipes := make([]*Recipe, 0)
	recipesMap := make(map[int]*Recipe)

	// Loop through the results
	for rows.Next() {
		var recipe Recipe
		var ri StepIngredient
		var ingredient Ingredient
		var step Step
		var triggerGroup TriggerGroup
		var trigger Trigger
		var triggerType TriggerType
		var utensil Utensil

		// Scan the row for all the data
		err := rows.Scan(
			&recipe.ID, &recipe.Title, &recipe.CreatedAt,
			&ri.ID, &ri.Quantity, &ri.Unit, &ri.CreatedAt,
			&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt,
			&step.ID, &step.Data, &step.StepNumber, &step.CreatedAt,
			&triggerGroup.ID, &triggerGroup.ActionParams, &triggerGroup.ActionKey, &triggerGroup.Service,
			&trigger.ID, &trigger.ActionParams, &trigger.Action, &trigger.Service, &trigger.TriggerParams, &trigger.CreatedAt,
			&triggerType.ID, &triggerType.CreatedAt, &triggerType.Key, &triggerType.SensorType,
			&utensil.ID, &utensil.Name, &utensil.CreatedAt,
		)

		// Continue to the next row if there is an error fetching the data
		if err != nil {
			log.Error(err.Error())
			continue
		}

		r, ok := recipesMap[recipe.ID]
		ri.Recipe = nil
		ri.Ingredient = &ingredient
		trigger.TriggerType = &triggerType

		if !ok {
			r = &recipe

			// Initialize the maps
			r.stepsMap = make(map[int]*Step)
			r.recipeIngredientMap = make(map[int]*StepIngredient)
			step.triggerGroupMap = make(map[null.Int]*TriggerGroup)
			step.utensilMap = make(map[null.Int]*Utensil)
			step.triggerGroupMap = make(map[null.Int]*TriggerGroup)
			triggerGroup.triggerMap = make(map[null.Int]*Trigger)

			// Add triggerGroup + trigger if valid trigger
			if triggerGroup.ID.Valid && trigger.ID.Valid {
				triggerGroup.triggerMap[trigger.ID] = &trigger
				triggerGroup.Triggers = []*Trigger{&trigger}
				step.triggerGroupMap[triggerGroup.ID] = &triggerGroup
				step.TriggerGroups = append(step.TriggerGroups, &triggerGroup)
			}

			// Add utensil if valid
			if utensil.ID.Valid {
				step.utensilMap[utensil.ID] = &utensil
				step.Utensils = []*Utensil{&utensil}
			}

			// Add the Recipe Ingredient
			r.RecipeIngredients = []*StepIngredient{&ri}
			r.recipeIngredientMap[ri.ID] = &ri

			// Add the step
			r.Steps = []*Step{&step}
			r.stepsMap[step.ID] = &step

			// Add the recipe to the list and map of recipes
			recipesMap[r.ID] = r
			recipes = append(recipes, r)
		} else {
			// If the recipe has already initialized, add any new data to the recipe
			// Add if a new recipeIngredient
			_, ok := r.recipeIngredientMap[ri.ID]
			if !ok {
				r.recipeIngredientMap[ri.ID] = &ri
				r.RecipeIngredients = append(r.RecipeIngredients, &ri)
			}

			// Add if a new step
			_, ok = r.stepsMap[step.ID]
			if !ok {
				r.stepsMap[step.ID] = &step
				r.Steps = append(r.Steps, &step)
				r.stepsMap[step.ID].triggerGroupMap = make(map[null.Int]*TriggerGroup)
				r.stepsMap[step.ID].utensilMap = make(map[null.Int]*Utensil)
			}

			// Add if a new TriggerGroup
			_, ok = r.stepsMap[step.ID].triggerGroupMap[triggerGroup.ID]
			if !ok && triggerGroup.ID.Valid {
				r.stepsMap[step.ID].triggerGroupMap[triggerGroup.ID] = &triggerGroup
				r.stepsMap[step.ID].TriggerGroups = append(r.stepsMap[step.ID].TriggerGroups, &triggerGroup)
				r.stepsMap[step.ID].triggerGroupMap[triggerGroup.ID].triggerMap = make(map[null.Int]*Trigger)
			}

			// Add if a new trigger and a valid TriggerGroup
			if triggerGroup.ID.Valid {
				_, ok = r.stepsMap[step.ID].triggerGroupMap[triggerGroup.ID].triggerMap[trigger.ID]
				if !ok && trigger.ID.Valid {
					tG, _ := r.stepsMap[step.ID].triggerGroupMap[triggerGroup.ID]
					tG.triggerMap[trigger.ID] = &trigger
					tG.Triggers = append(tG.Triggers, &trigger)
				}
			}

			// Add if a new Utensil
			_, ok = r.stepsMap[step.ID].utensilMap[utensil.ID]
			if !ok && utensil.ID.Valid {
				r.stepsMap[step.ID].utensilMap[utensil.ID] = &utensil
				r.stepsMap[step.ID].Utensils = append(r.stepsMap[step.ID].Utensils, &utensil)
			}
		}
	}
	return recipes
}
