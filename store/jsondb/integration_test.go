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
	lib := jsondb.New(baseDir)

	err = lib.StoreLibrary(bookLib)
	if err != nil {
		t.Errorf("ERROR: got error %v", err)
	}

	got, found, err := lib.LoadLibrary(testLibraryName)
	if !found {
		t.Errorf("ERROR: library not found")
	}

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

func TestIntegrationClassroomStoreLoad(t *testing.T) {
	// store, load, compare

	baseDir := filepath.Join(testDataDir, "store")
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Fatalf("ERROR: remove store dir failed: %v", err)
	}

	room := makeLearnClassroom()
	lib := jsondb.New(baseDir)

	err = lib.StoreClassroom(room)
	if err != nil {
		t.Errorf("ERROR: got error %v", err)
	}

	got, found, err := lib.LoadClassroom(testClassroomName)
	if !found {
		t.Errorf("ERROR: not found")
	}
	if err != nil {
		t.Errorf("ERROR: got error %v", err)
	}

	compareClassrom(t, got, room)
}
