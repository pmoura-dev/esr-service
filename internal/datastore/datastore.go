package datastore

import (
	"github.com/pmoura-dev/esr-service/internal/types"
)

type DataStore interface {
	Init() error
	Close()

	EntityRepository
	CommandRepository
	ReportSubscriptionRepository
}

type EntityRepository interface {
	GetEntityByID(id string) (types.Entity, error)
	ListEntities() ([]types.Entity, error)
	AddEntity(entity types.Entity) error
	DeleteEntity(id string) error
}

type CommandRepository interface {
	GetCommandByID(id string) (types.Command, error)
	ListCommands(filter Filter[types.Command]) ([]types.Command, error)
	AddCommand(command types.Command) error
	ResolveCommand(id string, result types.CommandStatus) error
	DeleteCommand(id string) error
}

type ReportSubscriptionRepository interface {
	GetReportSubscriptionByID(id int) (types.ReportSubscription, error)
	ListReportSubscriptions(filter Filter[types.ReportSubscription]) ([]types.ReportSubscription, error)
	AddReportSubscription(reportSubscription types.ReportSubscription) error
	DeleteReportSubscription(id int) error
	ActivateReportSubscription(id int) error
	DeactivateReportSubscription(id int) error
}

type MetricRepository interface{}

type StateRepository interface {
	GetStateByEntityID(entityID int) (types.State, error)
}

type Filter[T any] interface {
	Check(T) bool
}
