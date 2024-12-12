package datastore

import (
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

type DataStore interface {
	CreateTables() error

	EntityRepository
	CommandRepository
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

type StateRepository interface{}

type Filter[T any] interface {
	Check(T) bool
}
