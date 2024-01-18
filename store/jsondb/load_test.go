package jsondb_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestLoadLibrary(t *testing.T) {
	dirName := filepath.Join("testdata", jsondb.LibraryPath)
	loader := jsondb.NewLoader(dirName)

	libraries, err := loader.LoadLibraries()
	if err != nil {
		t.Errorf("LoadLibraries(): got error: %v", err)
	}

	wantLen := 1
	if len(libraries) != wantLen {
		t.Fatalf("number of libraries: got %v, want %v", len(libraries), wantLen)
	}

	lib := libraries[0]
	if lib.Title != testLibraryName {
		t.Errorf("Title: got %v, want %v", lib.Title, testLibraryName)
	}

	books := lib.Books()
	wantLen = 2
	if len(books) != wantLen {
		t.Fatalf("number of books: got %v, want %v", len(books), wantLen)
	}

	book := books[0]
	if book.ID.Title != testBookTitle1 ||
		book.ID.SeriesTitle != testBookSeriesTitle1 ||
		book.ID.Volume != testBookVolume1 {
		t.Errorf("TitleInfo: got %v, %v, %v, want %v, %v, %v",
			book.ID.Title,
			book.ID.SeriesTitle,
			book.ID.Volume,
			testBookTitle1,
			testBookSeriesTitle1,
			testBookVolume1,
		)
	}

	lessonTitles := book.Lessons()
	wantLessons := []string{testLessonTitle1, testLessonTitle2}
	if diff := cmp.Diff(lessonTitles, wantLessons); diff != "" {
		t.Errorf("Lessons(): -got +want\n%s", diff)
	}

	gotKanjis := book.KanjisFor(testLessonTitle1)
	if diff := cmp.Diff(gotKanjis, kanjiCardsLesson1); diff != "" {
		t.Errorf("KanjisFor(): -got +want\n%s", diff)
	}

	gotWords := book.WordsFor(testLessonTitle1)
	if diff := cmp.Diff(gotWords, wordCardsLesson1); diff != "" {
		t.Errorf("WordsFor(): -got +want\n%s", diff)
	}
}
