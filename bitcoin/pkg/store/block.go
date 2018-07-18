package store

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	block *leveldb.DB
}

// Create a database for blocks
func NewBlock(path string) (*DB, error) {
	b, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		b,
	}, nil
}
