package databases

import (
	"fmt"

	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore"
	"github.com/pmoura-dev/esr-service/internal/datastore/databases/boltdb"
)

func GetDataStore(config config.DataStoreConfig) (datastore.DataStore, error) {
	switch config.DataStoreType {
	case boltdb.Name:
		return boltdb.NewBoltDBDataStore(config)
	default:
		return nil, fmt.Errorf("unknown datastore type: %s", config.DataStoreType)
	}
}
