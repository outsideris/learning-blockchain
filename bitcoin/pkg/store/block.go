package store

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// Create a database for blocks
func CreateBlock(path string) (*leveldb.DB, error) {
	var block *leveldb.DB
	block, err := leveldb.OpenFile(path, nil)

	return block, err
}
