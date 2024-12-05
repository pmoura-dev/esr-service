package bbolt

import (
	"encoding/json"
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore"

	bolt "go.etcd.io/bbolt"
)

const (
	bucketEntity = "Entity"
)

// BBoltDataStore represents a Bolt datastore
type BBoltDataStore struct {
	db *bolt.DB
}

func NewBBoltDataStore(config config.DataStoreConfig) (*BBoltDataStore, error) {
	path := fmt.Sprintf("%s.db", config.Name)
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, datastore.ErrConnectionFailed
	}

	return &BBoltDataStore{db: db}, nil
}

func (s *BBoltDataStore) CreateTables() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketEntity)); err != nil {
			return err
		}

		return nil
	})
}

func (s *BBoltDataStore) GetEntityByID(id int) (datastore.Entity, error) {
	var entity datastore.Entity

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketEntity))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		key := []byte(fmt.Sprintf("%d", id))
		data := bucket.Get(key)
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

func (s *BBoltDataStore) GetAllEntities() ([]datastore.Entity, error) {
	var entityList []datastore.Entity

	err := s.db.View(func(tx *bolt.Tx) error {
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
