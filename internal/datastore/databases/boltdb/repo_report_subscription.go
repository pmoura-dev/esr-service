package boltdb

import (
	"encoding/json"
	"strconv"

	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/types"

	"go.etcd.io/bbolt"
)

func (s *DataStore) GetReportSubscriptionByID(id int) (types.ReportSubscription, error) {
	var subscription types.ReportSubscription

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		key := []byte(strconv.Itoa(id))
		data := bucket.Get(key)
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		if err := json.Unmarshal(data, &subscription); err != nil {
			return datastore.ErrInvalidData
		}

		return nil
	})

	if err != nil {
		return types.ReportSubscription{}, err
	}

	return subscription, nil
}

func (s *DataStore) ListReportSubscriptions(filter datastore.Filter[types.ReportSubscription]) ([]types.ReportSubscription, error) {
	var subscriptionList []types.ReportSubscription

	err := s.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		return bucket.ForEach(func(_, data []byte) error {
			var subscription types.ReportSubscription

			if err := json.Unmarshal(data, &subscription); err != nil {
				return datastore.ErrInvalidData
			}

			if filter.Check(subscription) {
				subscriptionList = append(subscriptionList, subscription)
			}

			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return subscriptionList, nil
}

func (s *DataStore) AddReportSubscription(reportSubscription types.ReportSubscription) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		id, _ := bucket.NextSequence()
		reportSubscription.ID = int(id)

		data, err := json.Marshal(reportSubscription)
		if err != nil {
			return datastore.ErrInvalidData
		}

		if err := bucket.Put([]byte(strconv.Itoa(int(id))), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) DeleteReportSubscription(id int) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		if bucket.Get([]byte(strconv.Itoa(id))) == nil {
			return datastore.ErrRecordNotFound
		}

		if err := bucket.Delete([]byte(strconv.Itoa(id))); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) ActivateReportSubscription(id int) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data := bucket.Get([]byte(strconv.Itoa(id)))
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		var subscription types.ReportSubscription

		if err := json.Unmarshal(data, &subscription); err != nil {
			return datastore.ErrInvalidData
		}

		subscription.IsActive = true
		data, err := json.Marshal(subscription)
		if err != nil {
			return datastore.ErrInvalidData
		}

		if err := bucket.Put([]byte(strconv.Itoa(id)), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}

func (s *DataStore) DeactivateReportSubscription(id int) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketReportSubscription))
		if bucket == nil {
			return datastore.ErrTableDoesNotExist
		}

		data := bucket.Get([]byte(strconv.Itoa(id)))
		if data == nil {
			return datastore.ErrRecordNotFound
		}

		var subscription types.ReportSubscription

		if err := json.Unmarshal(data, &subscription); err != nil {
			return datastore.ErrInvalidData
		}

		subscription.IsActive = false
		data, err := json.Marshal(subscription)
		if err != nil {
			return datastore.ErrInvalidData
		}

		if err := bucket.Put([]byte(strconv.Itoa(id)), data); err != nil {
			return datastore.ErrTransactionFailed
		}

		return nil
	})
}
