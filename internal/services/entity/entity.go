package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/pmoura-dev/esr-service/internal/broker"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/services"
	"github.com/pmoura-dev/esr-service/internal/types"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type BaseEntityService struct {
	datastore datastore.DataStore
	broker    broker.Broker
}

func NewBaseEntityService(datastore datastore.DataStore, broker broker.Broker) *BaseEntityService {
	return &BaseEntityService{
		datastore: datastore,
		broker:    broker,
	}
}

func (s *BaseEntityService) GetEntityByID(id string) (types.Entity, error) {
	entity, err := s.datastore.GetEntityByID(id)
	if err != nil {
		switch {
		case errors.Is(err, datastore.ErrRecordNotFound):
			return types.Entity{}, services.ErrEntityNotFound
		default:
			return types.Entity{}, services.ErrInternalError
		}
	}

	return entity, nil
}

func (s *BaseEntityService) ListEntities() ([]types.Entity, error) {
	entityList, err := s.datastore.ListEntities()
	if err != nil {
		return nil, services.ErrInternalError
	}

	return entityList, nil
}

func (s *BaseEntityService) AddEntity(entity types.Entity) error {
	if err := s.datastore.AddEntity(entity); err != nil {
		switch {
		case errors.Is(err, datastore.ErrDuplicateRecord):
			return services.ErrEntityAlreadyExists
		default:
			return services.ErrInternalError
		}
	}

	return nil
}

func (s *BaseEntityService) DeleteEntity(id string) error {
	if err := s.datastore.DeleteEntity(id); err != nil {
		switch {
		case errors.Is(err, datastore.ErrRecordNotFound):
			return services.ErrEntityNotFound
		default:
			return services.ErrInternalError
		}
	}

	return nil
}

func (s *BaseEntityService) ProcessCommand(entityID string, desiredState map[string]any) (string, error) {
	// check if entity exists
	_, err := s.datastore.GetEntityByID(entityID)
	if err != nil {
		switch {
		case errors.Is(err, datastore.ErrRecordNotFound):
			return "", services.ErrEntityNotFound
		default:
			return "", services.ErrInternalError
		}
	}

	commandID := generateCommandID()

	command := types.Command{
		ID:           commandID,
		EntityID:     entityID,
		DesiredState: desiredState,
		Status:       types.CommandStatusPending,
		IssuedAt:     time.Now(),
	}

	if err := s.datastore.AddCommand(command); err != nil {
		return "", services.ErrInternalError
	}

	topic := s.broker.Format(fmt.Sprintf("entities/%s/update", entityID))
	payload, err := json.Marshal(desiredState)
	if err != nil {
		return "", services.ErrInternalError
	}

	if err := s.broker.GetPublisher().Publish(topic, message.NewMessage(commandID, payload)); err != nil {
		return "", services.ErrInternalError
	}

	return commandID, nil
}

func generateCommandID() string {
	return uuid.NewString()
}
