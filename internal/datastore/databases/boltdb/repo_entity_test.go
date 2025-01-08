package boltdb

import (
	"errors"
	"reflect"
	"testing"

	"github.com/pmoura-dev/esr-service/internal/_data"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

func TestGetEntityByID(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		expected    models.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntity1,
			},
			inputID:  "1",
			expected: mockEntity1,
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
			store := DataStore{db: db}

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
		expected    []models.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			mocks: map[string]string{
				"1": _data.MockEntity1,
				"2": _data.MockEntity2,
			},
			expected: []models.Entity{
				mockEntity1,
				mockEntity2,
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
			store := DataStore{db: db}

			got, err := store.ListEntities()

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

		inputEntity models.Entity
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketEntity,
			inputEntity: models.Entity{
				ID:   "1",
				Name: "TestEntity",
			},
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
				"1": _data.MockEntity1,
			},
			inputEntity: models.Entity{
				ID: "1",
			},
			wantErr:     true,
			expectedErr: datastore.ErrDuplicateRecord,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.AddEntity(tt.inputEntity)

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
				"1": _data.MockEntity1,
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
				"1": _data.MockEntity1,
			},
			inputID:     "2",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

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

var (
	mockEntity1 = models.Entity{ID: "1", Name: "TestEntity1"}
	mockEntity2 = models.Entity{ID: "2", Name: "TestEntity2"}
)
