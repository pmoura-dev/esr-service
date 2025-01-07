package models

import (
	"time"
)

type State struct {
	ID         int
	EntityID   string
	State      map[string]any
	ReportedAt time.Time
}
