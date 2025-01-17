package boltdb

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/pmoura-dev/esr-service/internal/_data"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/filters"
	"github.com/pmoura-dev/esr-service/internal/types"
)

func TestGetCommandByID(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		expected    types.Command
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
			store := DataStore{db: db}

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

		inputFilter datastore.Filter[types.Command]
		expected    []types.Command
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
			inputFilter: filters.NewCommandFilter(),
			expected: []types.Command{
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
			inputFilter: filters.NewCommandFilter().ByEntityID("1"),
			expected: []types.Command{
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
			inputFilter: filters.NewCommandFilter().ByStatus(types.CommandStatusPending),
			expected: []types.Command{
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
			inputFilter: filters.NewCommandFilter().ByTimeAfterIssuing(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []types.Command{
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
			inputFilter: filters.NewCommandFilter().ByTimeBeforeIssuing(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []types.Command{
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
			store := DataStore{db: db}

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

func TestAddCommand(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputCommand types.Command
		wantErr      bool
		expectedErr  error
	}{
		{
			name:         "Success",
			bucket:       bucketCommand,
			inputCommand: mockCommand1Pending,
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.AddCommand(tt.inputCommand)

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

func TestResolveCommand(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     string
		inputStatus types.CommandStatus
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
			},
			inputID:     "cmd1",
			inputStatus: types.CommandStatusSuccess,
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
			},
			inputID:     "cmd2",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.ResolveCommand(tt.inputID, tt.inputStatus)

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

func TestDeleteCommand(t *testing.T) {
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
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
			},
			inputID: "cmd1",
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketCommand,
			mocks: map[string]string{
				"cmd1": _data.MockCommand1Pending,
			},
			inputID:     "cmd2",
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.DeleteCommand(tt.inputID)

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
	mockCommand1Pending = types.Command{
		ID:           "cmd1",
		EntityID:     "1",
		DesiredState: map[string]any{"power": "on"},
		Status:       types.CommandStatusPending,
		IssuedAt:     time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
	}

	mockCommand1Success = types.Command{
		ID:           "cmd2",
		EntityID:     "1",
		DesiredState: map[string]any{"power": "off"},
		Status:       types.CommandStatusSuccess,
		IssuedAt:     time.Date(2010, 11, 10, 23, 0, 0, 0, time.UTC),
		ResolvedAt:   _data.Ptr(time.Date(2010, 11, 10, 23, 0, 10, 0, time.UTC)),
	}

	mockCommand2Failed = types.Command{
		ID:           "cmd3",
		EntityID:     "2",
		DesiredState: map[string]any{"power": "off"},
		Status:       types.CommandStatusFailure,
		IssuedAt:     time.Date(2011, 11, 10, 23, 0, 0, 0, time.UTC),
		ResolvedAt:   _data.Ptr(time.Date(2011, 11, 10, 23, 0, 10, 0, time.UTC)),
	}
)
