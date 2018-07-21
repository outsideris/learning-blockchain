package store

import (
	"io/ioutil"
	"testing"
)

func TestNewBlockStore(t *testing.T) {
	tempPath, err := ioutil.TempDir("", "bitcoin")
	if err != nil {
		t.Error("Fail to create temprary DB path")
	}

	_, err = NewBlockStore(tempPath)
	if err != nil {
		t.Error(err)
	}
}
