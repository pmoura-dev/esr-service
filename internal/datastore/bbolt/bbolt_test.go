package bbolt

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/pmoura-dev/esr-service/internal/_data"
	"github.com/pmoura-dev/esr-service/internal/datastore"

	bolt "go.etcd.io/bbolt"
)

func TestGetEntityByID(t *testing.T) {
	db := setupMockDB(t)
	store := BBoltDataStore{db: db}

	tests := []struct {
		name        string
		id          int
		expected    datastore.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:     "Success",
			id:       1,
			expected: datastore.Entity{ID: 1, Name: "TestEntity"},
		},
		{
			name:        "Error - Invalid Data",
			id:          3,
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
		{
			name:        "Error - Record Not Found",
			id:          4,
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetEntityByID(tt.id)

			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Test failed. Expected error: %v, Got: %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Test failed. Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("Test failed. Expected: %+v, Got: %+v", tt.expected, got)
			}
		})
	}
}

func TestGetAllEntities(t *testing.T) {

}

func setupMockDB(t *testing.T) *bolt.DB {

	// Create a temporary file for the database
	tempFile, err := os.CreateTemp("", "mock.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tempFile.Close()

	// Open the BoltDB database
	db, err := bolt.Open(tempFile.Name(), 0600, nil)
	if err != nil {
		t.Fatalf("failed to open BoltDB: %v", err)
	}

	// Preload data
	err = db.Update(func(tx *bolt.Tx) error {
		// entities
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketEntity))
		if err != nil {
			return err
		}

		bucket.Put([]byte("1"), _data.MockEntityValid1)
		bucket.Put([]byte("2"), _data.MockEntityValid2)
		bucket.Put([]byte("3"), _data.MockEntityInvalid)

		return nil
	})

	if err != nil {
		t.Fatalf("failed to preload mock db: %v", err)
	}

	// Cleanup after the test
	t.Cleanup(func() {
		db.Close()
		os.Remove(tempFile.Name())
	})

	return db
}
