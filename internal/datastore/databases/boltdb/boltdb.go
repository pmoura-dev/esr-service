package boltdb

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"go.etcd.io/bbolt"
)

const (
	Name = "boltdb"
)

// DataStore represents a Bolt datastore
type DataStore struct {
	db *bbolt.DB
}

func NewBoltDBDataStore(config config.DataStoreConfig) (*DataStore, error) {
	path := fmt.Sprintf("%s.db", config.Name)
	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		return nil, datastore.ErrConnectionFailed
	}

	return &DataStore{db: db}, nil
}

func (s *DataStore) Close() {
	_ = s.db.Close()
}
