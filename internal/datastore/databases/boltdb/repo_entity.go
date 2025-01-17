package boltdb

import (
	"encoding/json"

	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/types"

	"go.etcd.io/bbolt"
)

func (s *DataStore) GetEntityByID(id string) (types.Entity, error) {
	var entity types.Entity

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
		return types.Entity{}, err
	}

	return entity, nil
}

func (s *DataStore) ListEntities() ([]types.Entity, error) {
	var entityList []types.Entity

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		return bucket.ForEach(func(_, data []byte) error {
			var entity types.Entity

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

func (s *DataStore) AddEntity(entity types.Entity) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		if bucket.Get([]byte(entity.ID)) != nil {
			return datastore.ErrDuplicateRecord
		}

		data, err := json.Marshal(entity)
		if err != nil {
			return datastore.ErrTransactionFailed
		}

		if err := bucket.Put([]byte(entity.ID), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) DeleteEntity(id string) error {
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
