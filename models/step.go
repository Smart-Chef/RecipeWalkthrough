package models

import (
	"gopkg.in/guregu/null.v3"
)

type Step struct {
	ID              null.Int                   `json:"id"`
	Recipe          *Recipe                    `json:"recipe,omitempty"`
	Data            null.String                `json:"data"`
	Utensils        []*Utensil                 `json:"utensil,omitempty"`
	utensilMap      map[null.Int]*Utensil      `json:"-"`
	TriggerGroups   []*TriggerGroup            `json:"trigger_groups,omitempty"`
	triggerGroupMap map[null.Int]*TriggerGroup `json:"-"`
	StepNumber      null.Int                   `json:"step_number"`
	CreatedAt       null.Time                  `json:"created_at"`
}

type StepIngredient struct {
	ID         null.Int    `json:"id"`
	Recipe     *Recipe     `json:"recipe,omitempty"`
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	Quantity   null.Int    `json:"quantity"`
	Unit       null.String `json:"unit"`
	CreatedAt  null.Time   `json:"created_at"`
}
