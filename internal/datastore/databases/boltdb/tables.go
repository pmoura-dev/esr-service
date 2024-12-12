package boltdb

import (
	"go.etcd.io/bbolt"
)

const (
	bucketEntity             = "Entity"
	bucketCommand            = "Command"
	bucketReportSubscription = "ReportSubscription"
)

func (s *DataStore) CreateTables() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketEntity)); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists([]byte(bucketCommand)); err != nil {
			return err
		}

		return nil
	})
}
