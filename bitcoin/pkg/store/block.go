package store

import (
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	block *leveldb.DB
}

// Create a database for blocks
func NewBlockStore(path string) (*DB, error) {
	log.Infof("Creating BlockStore from '%s'", path+"/blocks")
	store, err := newLevelDbStore(path + "/blocks")
	if err != nil {
		return nil, err
	}

	log.Info("BlockStore created")
	return &DB{
		store,
	}, nil
}
