package models

import (
	"gopkg.in/guregu/null.v3"
)

type Utensil struct {
	ID        null.Int    `json:"id,omitempty"`
	Name      null.String `json:"name,omitempty"`
	CreatedAt null.Time   `json:"created_at,omitempty"`
}
