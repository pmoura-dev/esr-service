package datastore

type DataStore interface {
	CreateTables() error

	EntityRepository
}

type EntityRepository interface {
	GetEntityByID(id string) (Entity, error)
	GetAllEntities() ([]Entity, error)
	AddEntity(id string, name string) error
	DeleteEntity(id string) error
}

type CommandRepository interface{}

type ReportSubscriptionRepository interface{}

type MetricRepository interface{}

type StateRepository interface{}
