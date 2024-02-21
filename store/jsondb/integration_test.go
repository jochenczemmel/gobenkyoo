//go:build inttest

package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestIntegrationLibraryStoreLoad(t *testing.T) {
	// store, load, compare

	baseDir := filepath.Join(testDataDir, "store")
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Fatalf("ERROR: remove store dir failed: %v", err)
	}

	bookLib := makeBooksLibrary()
	lib := jsondb.NewLibrary(baseDir)

	err = lib.Store(bookLib)
	if err != nil {
		t.Errorf("ERROR: got error %v", err)
	}

	got, err := lib.Load(testLibraryName)
	if err != nil {
		t.Errorf("ERROR: got error %v", err)
	}
	if diff := cmp.Diff(got, bookLib, cmp.AllowUnexported(
		books.Library{},
		books.Book{},
		books.Lesson{},
	)); diff != "" {
		t.Fatalf("ERROR: got- want+\n%s", diff)
	}
}
