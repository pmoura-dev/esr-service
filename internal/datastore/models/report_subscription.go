package models

import (
	"encoding/json"
	"errors"
	"time"
)

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
