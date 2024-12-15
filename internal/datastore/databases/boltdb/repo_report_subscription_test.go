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

func TestGetReportSubscriptionByID(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     int
		expected    models.ReportSubscription
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID:  1,
			expected: mockReportSubscription1State,
		},
		{
			name:        "Error - Table Not Found",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Invalid Data",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscriptionInvalid,
			},
			inputID:     1,
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
		{
			name:        "Error - Record Not Found",
			bucket:      bucketReportSubscription,
			inputID:     1,
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			got, err := store.GetReportSubscriptionByID(tt.inputID)

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

func TestListReportSubscriptions(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputFilter datastore.Filter[models.ReportSubscription]
		expected    []models.ReportSubscription
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success - No filter",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter(),
			expected: []models.ReportSubscription{
				mockReportSubscription1State,
				mockReportSubscription1MetricPower,
				mockReportSubscription2State,
			},
		},
		{
			name:   "Success - Filter by: EntityID",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter().ByEntityID("entity_1"),
			expected: []models.ReportSubscription{
				mockReportSubscription1State,
				mockReportSubscription1MetricPower,
			},
		},
		{
			name:   "Success - Filter by: ReportType",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter().ByReportType(models.ReportTypeState),
			expected: []models.ReportSubscription{
				mockReportSubscription1State,
				mockReportSubscription2State,
			},
		},
		{
			name:   "Success - Filter by: Is Active",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter().ByIsActive(true),
			expected: []models.ReportSubscription{
				mockReportSubscription1State,
				mockReportSubscription2State,
			},
		},
		{
			name:   "Success - Filter by: Time After Update",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter().ByTimeAfterUpdated(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []models.ReportSubscription{
				mockReportSubscription2State,
			},
		},
		{
			name:   "Success - Filter by: Time Before Update",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
				"2": _data.MockReportSubscription1MetricPower,
				"3": _data.MockReportSubscription2State,
			},
			inputFilter: filters.NewReportSubscriptionFilter().ByTimeBeforeUpdated(
				time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			expected: []models.ReportSubscription{
				mockReportSubscription1State,
				mockReportSubscription1MetricPower,
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
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"invalid": _data.MockReportSubscriptionInvalid,
			},
			wantErr:     true,
			expectedErr: datastore.ErrInvalidData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			got, err := store.ListReportSubscriptions(tt.inputFilter)

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

func TestAddReportSubscription(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputReportSubscription models.ReportSubscription
		wantErr                 bool
		expectedErr             error
	}{
		{
			name:                    "Success",
			bucket:                  bucketReportSubscription,
			inputReportSubscription: mockReportSubscription1State,
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

			err := store.AddReportSubscription(tt.inputReportSubscription)

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

func TestDeleteReportSubscription(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     int
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID: 1,
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID:     2,
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.DeleteReportSubscription(tt.inputID)

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

func TestActivateReportSubscription(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     int
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID: 1,
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID:     2,
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.ActivateReportSubscription(tt.inputID)

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

func TestDeactivateReportSubscription(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		mocks  map[string]string

		inputID     int
		wantErr     bool
		expectedErr error
	}{
		{
			name:   "Success",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID: 1,
		},
		{
			name:        "Error - Table Does Not Exist",
			bucket:      "test",
			wantErr:     true,
			expectedErr: datastore.ErrTableDoesNotExist,
		},
		{
			name:   "Error - Record Not Found",
			bucket: bucketReportSubscription,
			mocks: map[string]string{
				"1": _data.MockReportSubscription1State,
			},
			inputID:     2,
			wantErr:     true,
			expectedErr: datastore.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			db := setupMockDB(t, tt.bucket, tt.mocks)
			store := DataStore{db: db}

			err := store.DeactivateReportSubscription(tt.inputID)

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
	mockReportSubscription1State = models.ReportSubscription{
		ID:         1,
		EntityID:   "entity_1",
		ReportType: "state",
		IsActive:   true,
		UpdatedAt:  time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
	}

	mockReportSubscription1MetricPower = models.ReportSubscription{
		ID:         2,
		EntityID:   "entity_1",
		ReportType: "metric",
		Metric:     _data.Ptr("power"),
		UpdatedAt:  time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
	}

	mockReportSubscription2State = models.ReportSubscription{
		ID:         3,
		EntityID:   "entity_2",
		ReportType: "state",
		IsActive:   true,
		UpdatedAt:  time.Date(2011, 11, 10, 23, 0, 0, 0, time.UTC),
	}
)
