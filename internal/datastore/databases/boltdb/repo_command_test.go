package boltdb

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/pmoura-dev/esr-service/internal/_data"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/filters"
	"github.com/pmoura-dev/esr-service/internal/datastore/models"
)

func TestGetCommandByID(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		expected    models.Command
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
			},
			inputID:  "cmd1",
			expected: mockCommand1Pending,
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Invalid Data",
			bucket: bucketCommand,
			mocks: map[string]string{
				"invalid": _data.MockCommandInvalid,
			},
			inputID:     "invalid",
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
		{
			name:        "Error - Record Not Found",
			bucket:      bucketCommand,
			inputID:     "cmd1",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BoltDBDataStore{db: db}

			got, err := store.GetCommandByID(tt.inputID)

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

func TestGetAllCommands(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputFilter filters.CommandFilter
		expected    []models.Command
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success - No filter",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
				"cmd2": _data.MockCommand1Success,
				"cmd3": _data.MockCommand2Failed,
			},
			expected: []models.Command{
				mockCommand1Pending,
				mockCommand1Success,
				mockCommand2Failed,
			},
		},
		{
			name:   "Success - Filter by: EntityID",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
				"cmd2": _data.MockCommand1Success,
				"cmd3": _data.MockCommand2Failed,
			},
			inputFilter: *filters.NewCommandFilter().ByEntityID("1"),
			expected: []models.Command{
				mockCommand1Pending,
				mockCommand1Success,
			},
		},
		{
			name:   "Success - Filter by: Status",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
				"cmd2": _data.MockCommand1Success,
				"cmd3": _data.MockCommand2Failed,
			},
			inputFilter: *filters.NewCommandFilter().ByStatus(models.CommandStatusPending),
			expected: []models.Command{
				mockCommand1Pending,
			},
		},
		{
			name:   "Success - Filter by: Time After Issuing",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
				"cmd2": _data.MockCommand1Success,
				"cmd3": _data.MockCommand2Failed,
			},
			inputFilter: *filters.NewCommandFilter().ByTimeAfterIssuing(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []models.Command{
				mockCommand1Success,
				mockCommand2Failed,
			},
		},
		{
			name:   "Success - Filter by: Time Before Issuing",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
				"cmd2": _data.MockCommand1Success,
				"cmd3": _data.MockCommand2Failed,
			},
			inputFilter: *filters.NewCommandFilter().ByTimeBeforeIssuing(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []models.Command{
				mockCommand1Pending,
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
			bucket: bucketCommand,
			mocks: map[string]string{
				"invalid": _data.MockCommandInvalid,
			},
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := BoltDBDataStore{db: db}

			got, err := store.ListCommands(tt.inputFilter)

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

var (
	mockCommand1Pending = models.Command{
		ID:           "cmd1",
		EntityID:     "1",
		DesiredState: map[string]any{"power": "on"},
		Status:       models.CommandStatusPending,
		IssuedAt:     time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
	}

	mockCommand1Success = models.Command{
		ID:           "cmd2",
		EntityID:     "1",
		DesiredState: map[string]any{"power": "off"},
		Status:       models.CommandStatusSuccess,
		IssuedAt:     time.Date(2010, 11, 10, 23, 0, 0, 0, time.UTC),
		ResolvedAt:   _data.Ptr(time.Date(2010, 11, 10, 23, 0, 10, 0, time.UTC)),
	}

	mockCommand2Failed = models.Command{
		ID:           "cmd3",
		EntityID:     "2",
		DesiredState: map[string]any{"power": "off"},
		Status:       models.CommandStatusFailed,
		IssuedAt:     time.Date(2011, 11, 10, 23, 0, 0, 0, time.UTC),
		ResolvedAt:   _data.Ptr(time.Date(2011, 11, 10, 23, 0, 10, 0, time.UTC)),
	}
)
