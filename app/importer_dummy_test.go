package app_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app"
)

// TestImporterLoadLibrary uses a test dummy to check error handling.
func TestImporterLoadLibrary(t *testing.T) {
	testCases := []struct {
		name            string
		loadStorer      app.LibraryLoadStorer
		wantOK          bool
		wantLoadErrMsg  string
		wantStoreErrMsg string
	}{{
		name:       "ok",
		loadStorer: dummy{},
		wantOK:     true,
	}, {
		name:            "LibraryLoadStorer is nil",
		wantLoadErrMsg:  "no LibraryLoadStorer defined",
		wantStoreErrMsg: "no LibraryLoadStorer defined",
	}, {
		name:           "load returns error",
		loadStorer:     dummy{loadError: "load failed"},
		wantLoadErrMsg: "load failed",
	}, {
		name:       "load returns path error",
		loadStorer: dummy{pathError: "file does not exist"},
	}, {
		name:            "store error",
		loadStorer:      dummy{storeError: "store failed"},
		wantOK:          true,
		wantStoreErrMsg: "store failed",
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			importer := app.NewLibraryImporter(c.loadStorer)

			t.Run("load", func(t *testing.T) {
				ok, err := importer.LoadLibrary("")
				checkErrorMessage(t, err, c.wantLoadErrMsg)
				if ok != c.wantOK {
					t.Errorf("ERROR: got %v, want %v", ok, c.wantOK)
				}
			})

			t.Run("store", func(t *testing.T) {
				err := importer.StoreLibrary()
				checkErrorMessage(t, err, c.wantStoreErrMsg)
			})
		})
	}
}

func checkErrorMessage(t *testing.T, err error, want string) {
	t.Helper()

	if want != "" {
		if err == nil {
			t.Fatalf("ERROR: wanted error not detected")
		}
		if err.Error() != want {
			t.Fatalf("ERROR: got %q, want %q", err.Error(), want)
		}
		t.Logf("INFO: error message: %v", err)
		return
	}

	if err != nil {
		t.Errorf("ERROR: got error: %v", err)
	}
}
