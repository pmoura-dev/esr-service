package boltdb

import (
	"encoding/json"
	"time"

	"github.com/pmoura-dev/esr-service/internal/_data"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/filters"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"

	"go.etcd.io/bbolt"
)

func (s *DataStore) GetCommandByID(id string) (models.Command, error) {
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

func (s *DataStore) ListCommands(filter filters.CommandFilter) ([]models.Command, error) {
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

func (s *DataStore) AddCommand(command models.Command) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketCommand))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data, err := json.Marshal(command)
		if err != nil {
			return datastore.ErrInvalidData
		}

		if err := bucket.Put([]byte(command.ID), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) ResolveCommand(id string, status models.CommandStatus) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketCommand))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data := bucket.Get([]byte(id))
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		var command models.Command
		if err := json.Unmarshal(data, &command); err != nil {
			return datastore.ErrInvalidData
		}

		command.Status = status
		command.ResolvedAt = _data.Ptr(time.Now())

		data, err := json.Marshal(command)
		if err != nil {
			return datastore.ErrInvalidData
		}

		if err := bucket.Put([]byte(command.ID), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) DeleteCommand(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketCommand))
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
