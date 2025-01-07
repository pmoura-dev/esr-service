package datastore

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore/databases/boltdb"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

type DataStore interface {
	Init() error
	Close()

	EntityRepository
	CommandRepository
	ReportSubscriptionRepository
}

type EntityRepository interface {
	GetEntityByID(id string) (models.Entity, error)
	ListEntities() ([]models.Entity, error)
	AddEntity(id string, name string) error
	DeleteEntity(id string) error
}

type CommandRepository interface {
	GetCommandByID(id string) (models.Command, error)
	ListCommands(filter Filter[models.Command]) ([]models.Command, error)
	AddCommand(command models.Command) error
	ResolveCommand(id string, result models.CommandStatus) error
	DeleteCommand(id string) error
}

type ReportSubscriptionRepository interface {
	GetReportSubscriptionByID(id int) (models.ReportSubscription, error)
	ListReportSubscriptions(filter Filter[models.ReportSubscription]) ([]models.ReportSubscription, error)
	AddReportSubscription(reportSubscription models.ReportSubscription) error
	DeleteReportSubscription(id int) error
	ActivateReportSubscription(id int) error
	DeactivateReportSubscription(id int) error
}

type MetricRepository interface{}

type StateRepository interface {
	GetStateByEntityID(entityID int) (models.State, error)
}

type Filter[T any] interface {
	Check(T) bool
}

func GetDataStore(config config.DataStoreConfig) (DataStore, error) {
	switch config.DataStoreType {
	case boltdb.Name:
		return boltdb.NewBoltDBDataStore(config)
	default:
		return nil, fmt.Errorf("unknown datastore type: %s", config.DataStoreType)
	}
}
