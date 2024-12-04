package datastore

type DataStore interface {
	CreateTables() error

	EntityRepository
}

type EntityRepository interface {
	GetEntityByID(id int) (Entity, error)
	GetAllEntities() ([]Entity, error)
	AddEntity(name string) error
	DeleteEntity(id int) error
}

type CommandRepository interface{}

type ReportSubscriptionRepository interface{}

type MetricRepository interface{}

type StateRepository interface{}
