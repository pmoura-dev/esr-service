package boltdb

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore"

	"go.etcd.io/bbolt"
)

// BoltDBDataStore represents a Bolt datastore
type BoltDBDataStore struct {
	db *bbolt.DB
}

func NewBoltDBDataStore(config config.DataStoreConfig) (*BoltDBDataStore, error) {
	path := fmt.Sprintf("%s.db", config.Name)
	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		return nil, datastore.ErrConnectionFailed
	}

	return &BoltDBDataStore{db: db}, nil
}
