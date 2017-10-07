package boltutil

import (
	"github.com/boltdb/bolt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSetDB(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "boltutil_test")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(tmpdir)

	dbPath := filepath.Join(tmpdir, "database.bolt")

	t.Logf("database file: %s", dbPath)

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		return Set(tx, []string{"foo", "bar"}, "key", "value")
	})
	if err != nil {
		t.Error(err)
	}

	var ret string
	err = db.View(func(tx *bolt.Tx) error {
		return Get(tx, []string{"foo", "bar"}, "key", &ret)
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(ret)

	if ret != "value" {
		t.Errorf("expected 'value' but got: %s", ret)
	}

}
