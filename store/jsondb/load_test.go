package jsondb_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
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
	if book.TitleInfo.Title != testBookTitle1 ||
		book.TitleInfo.SeriesTitle != testBookSeriesTitle1 ||
		book.TitleInfo.Volume != testBookVolume1 {
		t.Errorf("TitleInfo: got %v, %v, %v, want %v, %v, %v",
			book.TitleInfo.Title,
			book.TitleInfo.SeriesTitle,
			book.TitleInfo.Volume,
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
	if diff := cmp.Diff(gotKanjis, kanjiCardsLesson1,
		cmp.Comparer(kanjiEqual)); diff != "" {
		t.Errorf("KanjisFor(): -got +want\n%s", diff)
	}

	gotWords := book.WordsFor(testLessonTitle1)
	if diff := cmp.Diff(gotWords, wordCardsLesson1); diff != "" {
		t.Errorf("WordsFor(): -got +want\n%s", diff)
	}
}

func kanjiEqual(got, want kanjis.Card) bool {
	if got.Kanji() != want.Kanji() {
		return false
	}
	gotDetails := got.Details()
	wantDetails := want.Details()
	if len(gotDetails) != len(wantDetails) {
		return false
	}

	for i := range gotDetails {
		if gotDetails[i].Reading != wantDetails[i].Reading ||
			gotDetails[i].ReadingKana != wantDetails[i].ReadingKana {
			return false
		}
		if len(gotDetails[i].Meanings) != len(wantDetails[i].Meanings) {
			return false
		}
		for j := range gotDetails[i].Meanings {
			if gotDetails[i].Meanings[j] != wantDetails[i].Meanings[j] {
				return false
			}
		}
	}
	return true
}
