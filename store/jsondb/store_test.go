package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestStoreLibrary(t *testing.T) {

	testBook1.AddKanjis(testLessonTitle1, kanjiCardsLesson1...)
	testBook1.AddWords(testLessonTitle1, wordCardsLesson1...)
	testBook1.AddKanjis(testLessonTitle2, kanjiCardsLesson2...)
	testBook1.AddWords(testLessonTitle2, wordCardsLesson2...)

	library := books.NewLibrary(testLibraryName)
	library.AddBooks(testBook1, testBook2)

	path := filepath.Join("testdata", "store")
	err := os.RemoveAll(path)
	if err != nil {
		t.Errorf("RemoveAll(%v): got error: %v", path, err)
	}

	storer := jsondb.NewStorer(path)
	err = storer.StoreLibrary(library)
	if err != nil {
		t.Errorf("Store(): got error: %v", err)
	}

	jsonFile := filepath.Join(path, jsondb.LibraryPath,
		testLibraryName+jsondb.JsonExtension)
	got, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("ReadFile(%v): got error: %v", jsonFile, err)
	}

	wantFile := filepath.Join("testdata", jsondb.LibraryPath, "want_japanology.json")
	want, err := os.ReadFile(wantFile)
	if err != nil {
		t.Fatalf("ReadFile(%v): got error: %v", wantFile, err)
	}

	if diff := cmp.Diff(string(got), string(want)); diff != "" {
		t.Errorf("stored file: -got +want\n%s", diff)
	}
}
