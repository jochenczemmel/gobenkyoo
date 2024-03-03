package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

const (
	testDataDir = "testdata"
)

func TestStoreError(t *testing.T) {

	bookLib := makeBooksLibrary()
	classroom := makeLearnClassroom()
	baseDir := filepath.Join(testDataDir, "store")
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Fatalf("ERROR: remove store dir failed: %v", err)
	}

	testCases := []struct {
		name    string
		dir     string
		wantErr bool
	}{{
		name:    "ok",
		dir:     baseDir,
		wantErr: false,
	}, {
		name:    "not ok",
		dir:     "/can/not/create",
		wantErr: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			lib := jsondb.New(c.dir)

			err := lib.StoreLibrary(bookLib)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ERROR: error not detected")
				}
				t.Logf("INFO: got error: %v", err)
				return
			}
			if err != nil {
				t.Errorf("ERROR: got error %v", err)
			}

			err = lib.StoreClassroom(classroom)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ERROR: error not detected")
				}
				t.Logf("INFO: got error: %v", err)
				return
			}
			if err != nil {
				t.Errorf("ERROR: got error %v", err)
			}

		})
	}
}
