package services

import (
	"github.com/pmoura-dev/esr-service/internal/types"
)

type EntityService interface {
	GetEntityByID(id string) (types.Entity, error)
	ListEntities() ([]types.Entity, error)
	AddEntity(entity types.Entity) error
	DeleteEntity(id string) error

	ProcessCommand(entityID string, desiredState map[string]any) (string, error)
}

type CommandService interface {
	GetCommandByID(id string)
}

type ReportSubscriptionService interface{}
