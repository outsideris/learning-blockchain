package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type DBStore struct {
	db *leveldb.DB
}

func NewDBStore(path string) (*DBStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &DBStore{db}, nil
}

func (store *DBStore) Put(key []byte, value []byte) error {
	return store.db.Put(key, value, nil)
}

func (store *DBStore) Get(key []byte) ([]byte, error) {
	data, err := store.db.Get(key, nil)
	return data, err
}

func (store *DBStore) Has(key []byte) (bool, error) {
	return store.db.Has(key, nil)
}

func (store *DBStore) Delete(key []byte) error {
	return store.db.Delete(key, nil)
}

func (store *DBStore) Close() error {
	return store.db.Close()
}
