package filters

import (
	"time"

	"github.com/pmoura-dev/esr-service/internal/types"
)

type CommandFilter struct {
	entityID     *string
	status       *types.CommandStatus
	issuedAfter  *time.Time
	issuedBefore *time.Time
}

func NewCommandFilter() *CommandFilter {
	return &CommandFilter{}
}

func (f *CommandFilter) ByEntityID(entityID string) *CommandFilter {
	f.entityID = &entityID
	return f
}

func (f *CommandFilter) ByStatus(status types.CommandStatus) *CommandFilter {
	f.status = &status
	return f
}

func (f *CommandFilter) ByTimeAfterIssuing(threshold time.Time) *CommandFilter {
	f.issuedAfter = &threshold
	return f
}

func (f *CommandFilter) ByTimeBeforeIssuing(threshold time.Time) *CommandFilter {
	f.issuedBefore = &threshold
	return f
}

func (f *CommandFilter) Check(command types.Command) bool {
	if f.entityID != nil && *f.entityID != command.EntityID {
		return false
	}

	if f.status != nil && *f.status != command.Status {
		return false
	}

	if f.issuedAfter != nil && !command.IssuedAt.After(*f.issuedAfter) {
		return false
	}

	if f.issuedBefore != nil && !command.IssuedAt.Before(*f.issuedBefore) {
		return false
	}

	return true
}
