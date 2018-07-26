package store

import (
	"github.com/outsideris/learning-blockchain/bitcoin/pkg/store/leveldb"
	log "github.com/sirupsen/logrus"
)

type DB struct {
	block *leveldb.DBStore
}

// Create a database for blocks
func NewBlockStore(path string) (*DB, error) {
	log.Infof("Creating BlockStore from '%s'", path+"/blocks")
	store, err := leveldb.NewDBStore(path + "/blocks")
	if err != nil {
		return nil, err
	}

	log.Info("BlockStore created")
	return &DB{
		store,
	}, nil
}
