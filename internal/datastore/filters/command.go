package filters

import (
	"time"

	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

type CommandFilter struct {
	EntityID     *string
	Status       *models.CommandStatus
	IssuedAfter  *time.Time
	IssuedBefore *time.Time
}

func NewCommandFilter() *CommandFilter {
	return &CommandFilter{}
}

func (f *CommandFilter) ByEntityID(entityID string) *CommandFilter {
	f.EntityID = &entityID
	return f
}

func (f *CommandFilter) ByStatus(status models.CommandStatus) *CommandFilter {
	f.Status = &status
	return f
}

func (f *CommandFilter) ByTimeAfterIssuing(threshold time.Time) *CommandFilter {
	f.IssuedAfter = &threshold
	return f
}

func (f *CommandFilter) ByTimeBeforeIssuing(threshold time.Time) *CommandFilter {
	f.IssuedBefore = &threshold
	return f
}

func (f *CommandFilter) Check(command models.Command) bool {
	if f.EntityID != nil && *f.EntityID != command.EntityID {
		return false
	}

	if f.Status != nil && *f.Status != command.Status {
		return false
	}

	if f.IssuedAfter != nil && !command.IssuedAt.After(*f.IssuedAfter) {
		return false
	}

	if f.IssuedBefore != nil && !command.IssuedAt.Before(*f.IssuedBefore) {
		return false
	}

	return true
}
