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
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		expected    datastore.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityValid1,
			},
			inputID:  "1",
			expected: datastore.Entity{ID: "1", Name: "TestEntity1"},
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			inputID:     "1",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Invalid Data",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityInvalid,
			},
			inputID:     "1",
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
		{
			name:        "Error - Record Not Found",
			bucket:      bucketEntity,
			inputID:     "1",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BBoltDataStore{db: db}

			got, err := store.GetEntityByID(tt.inputID)

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
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     int
		expected    []datastore.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityValid1,
				"2": _data.MockEntityValid2,
			},
			expected: []datastore.Entity{
				{ID: "1", Name: "TestEntity1"},
				{ID: "2", Name: "TestEntity2"},
			},
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Invalid Data",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityInvalid,
			},
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BBoltDataStore{db: db}

			got, err := store.GetAllEntities()

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

func TestAddEntity(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		inputName   string
		wantErr     bool
		expectedErr error
	}{
		{
			name:      "Success",
			bucket:    bucketEntity,
			inputID:   "1",
			inputName: "TestEntity",
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Duplicate Record",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityValid1,
			},
			inputID:     "1",
			wantErr:     true,
			expectedErr: datastore.ErrDuplicateRecord,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BBoltDataStore{db: db}

			err := store.AddEntity(tt.inputID, tt.inputName)

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
		})
	}
}

func TestDeleteEntity(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityValid1,
			},
			inputID: "1",
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntityValid1,
			},
			inputID:     "2",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BBoltDataStore{db: db}

			err := store.DeleteEntity(tt.inputID)

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
		})
	}
}

func setupMockDB(t *testing.T, bucket string, pairs map[string]string) *bolt.DB {
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
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		for k, v := range pairs {
			bucket.Put([]byte(k), []byte(v))
		}

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
