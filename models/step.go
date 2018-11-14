package models

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Step struct {
	ID         int                   `json:"id"`
	Recipe     *Recipe               `json:"recipe,omitempty"`
	Data       string                `json:"data"`
	Utensils   []*Utensil            `json:"utensil,omitempty"`
	utensilMap map[null.Int]*Utensil `json:"-"`
	Triggers   []*Trigger            `json:"triggers,omitempty"`
	triggerMap map[null.Int]*Trigger `json:"-"`
	StepNumber int                   `json:"step_number"`
	CreatedAt  time.Time             `json:"created_at"`
}

type StepIngredient struct {
	ID         int         `json:"id"`
	Recipe     *Recipe     `json:"recipe,omitempty"`
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	Quantity   int         `json:"quantity"`
	Unit       string      `json:"unit"`
	CreatedAt  time.Time   `json:"created_at"`
}
