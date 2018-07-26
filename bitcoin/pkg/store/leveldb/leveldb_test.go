package leveldb

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testDBStore *DBStore

func TestMain(m *testing.M) {
	tempPath, err := ioutil.TempDir("", "leveldb")
	if err != nil {
		fmt.Printf("Fail to create temporary DB path: %s\n", err)
	}

	testDBStore, err = NewDBStore(tempPath)
	if err != nil {
		fmt.Printf("NewLevelDBStore error:%s\n", err)
	}

	code := m.Run()

	testDBStore.Close()
	os.RemoveAll(tempPath)
	os.Exit(code)
}

func TestDBStore(t *testing.T) {
	k := "foo"
	v := "bar"
	err := testDBStore.Put([]byte(k), []byte(v))
	if err != nil {
		t.Errorf("DBStore.Put:%s", err)
		return
	}

	data, err := testDBStore.Get([]byte(k))
	if err != nil {
		t.Errorf("DBStore.Get:%s", err)
		return
	}
	if string(data) != v {
		t.Errorf("%s != %s", data, v)
		return
	}

	err = testDBStore.Delete([]byte(k))
	if err != nil {
		t.Errorf("DBStore.Delete:%s", err)
		return
	}

	ok, err := testDBStore.Has([]byte(k))
	if err != nil {
		t.Errorf("DBStore.Has:%s", err)
		return
	}
	if ok {
		t.Errorf("%s should not be exist", k)
		return
	}
}
