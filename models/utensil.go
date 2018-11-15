package models

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
)

type Utensil struct {
	ID        null.Int    `json:"id,omitempty"`
	Name      null.String `json:"name,omitempty"`
	CreatedAt null.Time   `json:"created_at,omitempty"`
}

func (*Utensil) GetAll(db *sql.DB, dot *dotsql.DotSql) []*Utensil {
	rows, err := dot.Query(db, "get-all-utensils")
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for GetRecipes")
	}

	utensils := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			utensils = append(utensils, &utensil)
		}
	}
	return utensils
}

func (i *Utensil) GetAllByRecipe(db *sql.DB, dot *dotsql.DotSql, recipeID string) []*Utensil {
	rows, err := dot.Query(db, "get-recipe-utensil-by-recipe-id", recipeID)
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for Utensil.GetAllByRecipe")
	}

	utensils := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			utensils = append(utensils, &utensil)
		}
	}
	return utensils
}

func (i *Utensil) GetAllByRecipeStep(db *sql.DB, dot *dotsql.DotSql, recipeID string, stepNumber string) []*Utensil {
	rows, err := dot.Query(db, "get-recipe-utensils-by-recipe-step-number", recipeID, stepNumber)
	defer rows.Close()

	if err != nil {
		log.Error("Error with query for Utensil.GetAllByRecipeStep")
	}

	ingredients := make([]*Utensil, 0)
	for rows.Next() {
		var utensil Utensil

		err := rows.Scan(&utensil.ID, &utensil.Name, &utensil.CreatedAt)

		if err != nil {
			log.Error(err.Error())
		} else {
			ingredients = append(ingredients, &utensil)
		}
	}
	return ingredients
}
