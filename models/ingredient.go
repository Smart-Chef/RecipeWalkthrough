package models

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	"gopkg.in/guregu/null.v3"
)

type Ingredient struct {
	ID        null.Int        `json:"id,omitempty"`
	Name      null.String     `json:"name,omitempty"`
	StepInfo  *StepIngredient `json:"step_info,omitempty"`
	CreatedAt null.Time       `json:"created_at,omitempty"`
}

// GetALl Ingredients in the database
func (*Ingredient) GetAll(db *sql.DB, dot *dotsql.DotSql) ([]*Ingredient, error) {
	rows, err := dot.Query(db, "get-all-ingredients")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient

		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients, nil
}

// GetAllByRecipe - get all ingredients by recipe id
func (i *Ingredient) GetAllByRecipe(db *sql.DB, dot *dotsql.DotSql, recipeID string) ([]*Ingredient, error) {
	rows, err := dot.Query(db, "get-recipe-ingredients-by-recipe-id", recipeID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients, nil
}

// GetAllByRecipeStep - get all ingredients by recipe step id and recipe id
func (i *Ingredient) GetAllByRecipeStep(db *sql.DB, dot *dotsql.DotSql, recipeID string, stepNumber string) ([]*Ingredient, error) {
	rows, err := dot.Query(db, "get-recipe-ingredients-by-recipe-step-number", recipeID, stepNumber)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ingredients := make([]*Ingredient, 0)
	for rows.Next() {
		var ingredient Ingredient
		var stepIngredient StepIngredient

		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.CreatedAt, &stepIngredient.ID, &stepIngredient.Unit, &stepIngredient.Quantity, &stepIngredient.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			ingredient.StepInfo = &stepIngredient
			ingredients = append(ingredients, &ingredient)
		}
	}
	return ingredients, nil
}
