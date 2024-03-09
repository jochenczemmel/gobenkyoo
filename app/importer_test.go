package app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestImporter(t *testing.T) {
	testCases := []struct {
		name            string
		loadStorer      app.LibraryLoadStorer
		kanjiImporter   app.KanjiImporter
		wordImporter    app.WordImporter
		wantOK          bool
		wantStoreErrMsg string
	}{{
		name:       "ok",
		loadStorer: dummy{},
		wantOK:     true,
	}, {
		name:            "LibraryLoadStorer is nil",
		wantStoreErrMsg: "no LibraryLoadStorer defined",
	}, {
		name:            "load returns error",
		loadStorer:      dummy{loadError: "load failed"},
		wantStoreErrMsg: "load failed",
	}, {
		name:       "load returns path error",
		loadStorer: dummy{pathError: "file does not exist"},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			importer := app.NewLibraryImporter(c.loadStorer)
			ok, err := importer.LoadLibrary("")
			if c.wantStoreErrMsg != "" {
				if err == nil {
					t.Fatalf("ERROR: wanted error not detected")
				}
				if err.Error() != c.wantStoreErrMsg {
					t.Fatalf("ERROR: got %q, want %q",
						err.Error(), c.wantStoreErrMsg)
				}
				t.Logf("INFO: error message: %v", err)
				return
			}

			if err != nil {
				t.Errorf("ERROR: got error: %v", err)
			}
			if ok != c.wantOK {
				t.Errorf("ERROR: got %v, want %v", ok, c.wantOK)
			}
		})
	}
}

type dummy struct {
	loadError, pathError, storeError string
	kanjiError, wordError            string
}

func (d dummy) LoadLibrary(string) (books.Library, error) {
	var result books.Library
	if d.loadError != "" {
		return result, fmt.Errorf("%s", d.loadError)
	}
	if d.pathError != "" {
		return result, &os.PathError{
			Op:   "open",
			Path: ".",
			Err:  os.ErrNotExist,
		}
	}
	return result, nil
}

func (d dummy) StoreLibrary(books.Library) error {
	if d.storeError != "" {
		return fmt.Errorf("%s", d.storeError)
	}
	return nil
}

func (d dummy) ImportKanji(string) ([]kanjis.Card, error) {
	var result []kanjis.Card
	if d.kanjiError != "" {
		return result, fmt.Errorf("%s", d.kanjiError)
	}
	return result, nil
}

func (d dummy) ImportWord(string) ([]words.Card, error) {
	var result []words.Card
	if d.wordError != "" {
		return result, fmt.Errorf("%s", d.wordError)
	}
	return result, nil
}
