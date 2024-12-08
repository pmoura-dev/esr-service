package datastore

import (
	"github.com/pmoura-dev/esr-service/internal/datastore/filters"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

type DataStore interface {
	CreateTables() error

	EntityRepository
}

type EntityRepository interface {
	GetEntityByID(id string) (models.Entity, error)
	ListEntities() ([]models.Entity, error)
	AddEntity(id string, name string) error
	DeleteEntity(id string) error
}

type CommandRepository interface {
	GetCommandByID(id string) (models.Command, error)
	ListCommands(filter filters.CommandFilter) ([]models.Command, error)
	AddCommand(command models.Command) error
	DeleteCommand(id string) error
}

type ReportSubscriptionRepository interface{}

type MetricRepository interface{}

type StateRepository interface{}
