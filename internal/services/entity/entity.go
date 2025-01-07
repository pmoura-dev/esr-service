package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pmoura-dev/esr-service/internal/broker"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"

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

func (s *BaseEntityService) ListEntities() ([]models.Entity, error) {
	return s.datastore.ListEntities()
}

func (s *BaseEntityService) AddEntity(entity models.Entity) error {
	return s.datastore.AddEntity(entity.ID, entity.Name)
}

func (s *BaseEntityService) ProcessCommand(entityID string, desiredState map[string]any) (string, error) {
	commandID := generateCommandID()

	command := models.Command{
		ID:           commandID,
		EntityID:     entityID,
		DesiredState: desiredState,
		Status:       models.CommandStatusPending,
		IssuedAt:     time.Now(),
	}

	if err := s.datastore.AddCommand(command); err != nil {
		return "", err
	}

	topic := fmt.Sprintf("entities/%s/update", entityID)
	payload, err := json.Marshal(desiredState)
	if err != nil {
		return "", err
	}

	if err := s.broker.GetPublisher().Publish(topic, message.NewMessage(commandID, payload)); err != nil {
		return "", err
	}

	return commandID, nil
}

func generateCommandID() string {
	return uuid.NewString()
}
