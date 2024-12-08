package models

import (
	"encoding/json"
	"errors"
	"time"
)

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

func (cs *CommandStatus) UnmarshalJSON(data []byte) error {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}

	switch CommandStatus(status) {
	case CommandStatusPending, CommandStatusSuccess, CommandStatusFailed:
		*cs = CommandStatus(status)
		return nil
	default:
		return errors.New("invalid CommandStatus value")
	}
}

type Command struct {
	ID           string         `json:"id"`
	EntityID     string         `json:"entity_id"`
	DesiredState map[string]any `json:"desired_state"`
	Status       CommandStatus  `json:"status"`
	IssuedAt     time.Time      `json:"issued_at"`
	ResolvedAt   *time.Time     `json:"resolved_at"`
}
