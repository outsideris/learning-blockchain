package store

import (
	"io/ioutil"
	"testing"
)

func TestNewLevelDbStore(t *testing.T) {
	tempPath, err := ioutil.TempDir("", "bitcoin")
	if err != nil {
		t.Error("Fail to create temprary DB path")
	}

	_, err = newLevelDbStore(tempPath)
	if err != nil {
		t.Error(err)
	}
}
