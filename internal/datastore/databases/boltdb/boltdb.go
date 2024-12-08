package boltdb

import (
	"encoding/json"
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore"

	"go.etcd.io/bbolt"
)

const (
	bucketEntity = "Entity"
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

func (s *BoltDBDataStore) CreateTables() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketEntity)); err != nil {
			return err
		}

		return nil
	})
}

func (s *BoltDBDataStore) GetEntityByID(id string) (datastore.Entity, error) {
	var entity datastore.Entity

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data := bucket.Get([]byte(id))
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		if err := json.Unmarshal(data, &entity); err != nil {
			return datastore.ErrInvalidData
		}

		return nil
	})

	if err != nil {
		return datastore.Entity{}, err
	}

	return entity, nil
}

func (s *BoltDBDataStore) GetAllEntities() ([]datastore.Entity, error) {
	var entityList []datastore.Entity

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		return bucket.ForEach(func(_, data []byte) error {
			var entity datastore.Entity

			if err := json.Unmarshal(data, &entity); err != nil {
				return datastore.ErrInvalidData
			}

			entityList = append(entityList, entity)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return entityList, nil
}

func (s *BoltDBDataStore) AddEntity(id string, name string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		if bucket.Get([]byte(id)) != nil {
			return datastore.ErrDuplicateRecord
		}

		newEntity := datastore.Entity{
			ID:   id,
			Name: name,
		}

		data, err := json.Marshal(newEntity)
		if err != nil {
			return datastore.ErrTransactionFailed
		}

		if err := bucket.Put([]byte(id), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *BoltDBDataStore) DeleteEntity(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		if bucket.Get([]byte(id)) == nil {
			return datastore.ErrRecordNotFound
		}

		if err := bucket.Delete([]byte(id)); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}
