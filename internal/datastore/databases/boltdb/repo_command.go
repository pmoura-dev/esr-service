package boltdb

import (
	"encoding/json"

	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/filters"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"

	"go.etcd.io/bbolt"
)

func (s *BoltDBDataStore) GetCommandByID(id string) (models.Command, error) {
	var command models.Command

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketCommand))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data := bucket.Get([]byte(id))
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		if err := json.Unmarshal(data, &command); err != nil {
			return datastore.ErrInvalidData
		}

		return nil
	})

	if err != nil {
		return models.Command{}, err
	}

	return command, nil
}

func (s *BoltDBDataStore) ListCommands(filter filters.CommandFilter) ([]models.Command, error) {
	var commandList []models.Command

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketCommand))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		return bucket.ForEach(func(_, data []byte) error {
			var command models.Command

			if err := json.Unmarshal(data, &command); err != nil {
				return datastore.ErrInvalidData
			}

			if !filter.Check(command) {
				return nil
			}

			commandList = append(commandList, command)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return commandList, nil
}
