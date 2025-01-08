package services

import (
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

type EntityService interface {
	AddEntity(entity models.Entity) error
	ListEntities() ([]models.Entity, error)
	ProcessCommand(entityID string, desiredState map[string]any) (string, error)
}
