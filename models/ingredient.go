package models

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
)

type Ingredient struct {
	ID        null.Int        `json:"id,omitempty"`
	Name      null.String     `json:"name,omitempty"`
	StepInfo  *StepIngredient `json:"step_info,omitempty"`
	CreatedAt null.Time       `json:"created_at,omitempty"`
}

func (*Ingredient) GetAll(db *sql.DB, dot *dotsql.DotSql) []*Ingredient {
	rows, err := dot.Query(db, "get-all-ingredients")
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for GetRecipes")
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient

		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients
}

func (i *Ingredient) GetAllByRecipe(db *sql.DB, dot *dotsql.DotSql, recipeID string) []*Ingredient {
	rows, err := dot.Query(db, "get-recipe-ingredients-by-recipe-id", recipeID)
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for StepIngredient.GetAllByRecipe")
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient

		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients
}

func (i *Ingredient) GetAllByRecipeStep(db *sql.DB, dot *dotsql.DotSql, recipeID string, stepNumber string) []*Ingredient {
	rows, err := dot.Query(db, "get-recipe-ingredients-by-recipe-step-number", recipeID, stepNumber)
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for StepIngredient.GetAllByRecipeStep")
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient
		var stepIngredient StepIngredient

		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt, &stepIngredient.ID, &stepIngredient.Unit, &stepIngredient.Quantity, &stepIngredient.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			ingredient.StepInfo = &stepIngredient
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients
}
