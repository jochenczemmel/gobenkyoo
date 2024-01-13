package jsondb_test

import (
	"path/filepath"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestStoreBook(t *testing.T) {

	path := filepath.Join("testdata", "store")
	storer := jsondb.NewStorer(path)

	t.Logf("DEBUG: %v", storer)

}
