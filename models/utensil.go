package models

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	"gopkg.in/guregu/null.v3"
)

type Utensil struct {
	ID        null.Int    `json:"id,omitempty"`
	Name      null.String `json:"name,omitempty"`
	CreatedAt null.Time   `json:"created_at,omitempty"`
}

// GetAll utensils in the database
func (*Utensil) GetAll(db *sql.DB, dot *dotsql.DotSql) ([]*Utensil, error) {
	rows, err := dot.Query(db, "get-all-utensils")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	utensils := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			utensils = append(utensils, &utensil)
		}
	}
	return utensils, nil
}

// GetAllByRecipe - get all utensils by recipe id
func (i *Utensil) GetAllByRecipe(db *sql.DB, dot *dotsql.DotSql, recipeID string) ([]*Utensil, error) {
	rows, err := dot.Query(db, "get-recipe-utensil-by-recipe-id", recipeID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	utensils := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			utensils = append(utensils, &utensil)
		}
	}
	return utensils, nil
}

// GetAllByRecipeStep - get all utensils by recipe step id and recipe id
func (i *Utensil) GetAllByRecipeStep(db *sql.DB, dot *dotsql.DotSql, recipeID string, stepNumber string) ([]*Utensil, error) {
	rows, err := dot.Query(db, "get-recipe-utensils-by-recipe-step-number", recipeID, stepNumber)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	ingredients := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			return nil, err
		} else {
			ingredients = append(ingredients, &utensil)
		}
	}
	return ingredients, nil
}
