package models

import "gopkg.in/guregu/null.v3"

type Trigger struct {
	ID            null.Int     `json:"id"`
	Parameters    null.String  `json:"parameters"`
	Action        null.String  `json:"action"`
	ActionParams  null.String  `json:"action_params"`
	TriggerParams null.String  `json:"trigger_params"`
	Service       null.String  `json:"service"`
	Step          *Step        `json:"step,omitempty"`
	TriggerType   *TriggerType `json:"trigger_type,omitempty"`
	CreatedAt     null.Time    `json:"created_at"`
}

type TriggerType struct {
	ID         null.Int    `json:"id"`
	Key        null.String `json:"key"`
	SensorType null.String `json:"sensor_type"`
	Triggers   []*Trigger  `json:"triggers,omitempty"`
	CreatedAt  null.Time   `json:"created_at"`
}
