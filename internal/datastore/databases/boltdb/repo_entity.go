package boltdb

import (
	"encoding/json"

	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"

	"go.etcd.io/bbolt"
)

func (s *BoltDBDataStore) GetEntityByID(id string) (models.Entity, error) {
	var entity models.Entity

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
		return models.Entity{}, err
	}

	return entity, nil
}

func (s *BoltDBDataStore) ListEntities() ([]models.Entity, error) {
	var entityList []models.Entity

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		return bucket.ForEach(func(_, data []byte) error {
			var entity models.Entity

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

		newEntity := models.Entity{
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
