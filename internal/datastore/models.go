package datastore

import "time"

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CommandStatus string

const (
	CommandStatusPending CommandStatus = "pending"
	CommandStatusSuccess CommandStatus = "success"
	CommandStatusFailed  CommandStatus = "failed"
)

type Command struct {
	ID           int            `json:"id"`
	EntityID     string         `json:"entity_id"`
	DesiredState map[string]any `json:"desired_state"`
	Status       CommandStatus  `json:"status"`
	IssuedAt     time.Time      `json:"issued_at"`
	ResolvedAt   *time.Time     `json:"resolved_at"`
}
