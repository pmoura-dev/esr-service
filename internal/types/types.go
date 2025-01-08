package types

import (
	"encoding/json"
	"errors"
	"time"
)

type Command struct {
	ID           string         `json:"id"`
	EntityID     string         `json:"entity_id"`
	DesiredState map[string]any `json:"desired_state"`
	Status       CommandStatus  `json:"status"`
	IssuedAt     time.Time      `json:"issued_at"`
	ResolvedAt   *time.Time     `json:"resolved_at"`
}

type CommandStatus string

const (
	CommandStatusPending CommandStatus = "pending"
	CommandStatusSuccess CommandStatus = "success"
	CommandStatusFailure CommandStatus = "failure"
)

func (cs *CommandStatus) UnmarshalJSON(data []byte) error {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}

	switch CommandStatus(status) {
	case CommandStatusPending, CommandStatusSuccess, CommandStatusFailure:
		*cs = CommandStatus(status)
		return nil
	default:
		return errors.New("invalid CommandStatus value")
	}
}

type State struct {
	ID         int
	EntityID   string
	State      map[string]any
	ReportedAt time.Time
}

type ReportSubscription struct {
	ID         int        `json:"id"`
	EntityID   string     `json:"entity_id"`
	ReportType ReportType `json:"report_type"`
	Metric     *string    `json:"metric,omitempty"`
	IsActive   bool       `json:"is_active"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type ReportType string

const (
	ReportTypeState  ReportType = "state"
	ReportTypeMetric ReportType = "metric"
)

func (rt *ReportType) UnmarshalJSON(data []byte) error {
	var reportType string
	if err := json.Unmarshal(data, &reportType); err != nil {
		return err
	}

	switch ReportType(reportType) {
	case ReportTypeState, ReportTypeMetric:
		*rt = ReportType(reportType)
		return nil
	default:
		return errors.New("invalid ReportType value")
	}
}
