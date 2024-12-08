package boltdb

import (
	"os"
	"testing"

	"go.etcd.io/bbolt"
)

func setupMockDB(t *testing.T, bucket string, pairs map[string]string) *bbolt.DB {
	// Create a temporary file for the database
	tempFile, err := os.CreateTemp("", "mock.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tempFile.Close()

	// Open the bboltDB database
	db, err := bbolt.Open(tempFile.Name(), 0600, nil)
	if err != nil {
		t.Fatalf("failed to open bboltDB: %v", err)
	}

	// Preload data
	err = db.Update(func(tx *bbolt.Tx) error {

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
