package store

import (
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

func newLevelDbStore(path string) (*leveldb.DB, error) {
	log.Infof("Creating LevelDB from '%s'", path)
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	log.Info("LevelDB created")
	return db, nil
}
