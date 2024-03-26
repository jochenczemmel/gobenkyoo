package store_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

const testDataDir = "testdata"

// TestImportWordLesson tests the import of WordImporter
// using the csvimport implementation.
func TestImportWordLesson(t *testing.T) {

	bookID := books.ID{Title: "minna"}
	lessonID := books.LessonID{Name: "lesson1", ID: bookID}
	wordImporter := csvimport.NewWord(',', true,
		[]string{"KANA", "NIHONGO", "ROMAJI", "MEANING", "HINT",
			"EXPLANATION", "DICTFORM", "TEFORM", "NAIFORM"})

	testCases := []struct {
		name             string
		fileName         string
		csvImporter      store.WordImporter
		wantErr          bool
		wantNBooks       int
		wantLessonTitles []string
		wantCards        []words.Card
	}{{
		name:             "ok",
		fileName:         filepath.Join(testDataDir, "words1.csv"),
		csvImporter:      wordImporter,
		wantNBooks:       1,
		wantLessonTitles: []string{lessonID.Name},
		wantCards:        words1,
	}, {
		name:             "file not found",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		csvImporter:      wordImporter,
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []words.Card{},
	}, {
		name:             "missing importer",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []words.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			importer := store.NewWordImporter(
				books.NewLibrary("testlib"), c.csvImporter)
			checkError(t, importer.Lesson(c.fileName, lessonID), c.wantErr)

			gotNbooks := len(importer.Library.SortedBookIDs())
			if gotNbooks != c.wantNBooks {
				t.Errorf("ERROR: got %v, want %v",
					gotNbooks, c.wantNBooks)
			}

			book := importer.Library.Book(bookID)
			if diff := cmp.Diff(book.LessonNames(), c.wantLessonTitles); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}

			lesson := book.Lesson(lessonID.Name)
			if diff := cmp.Diff(lesson.WordCards(), c.wantCards); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

// TestImportKanjiLesson tests the import of KanjiImporter
// using the csvimport implementation.
func TestImportKanjiLesson(t *testing.T) {

	bookID := books.ID{Title: "kanjidic"}
	lessonID := books.LessonID{Name: "lesson1", ID: bookID}
	kanjiImporter := csvimport.NewKanji(';', '/', true,
		[]string{"kanji", "", "", "reading", "meanings",
			"hint", "explanation"})

	testCases := []struct {
		name             string
		fileName         string
		csvImporter      store.KanjiImporter
		wantErr          bool
		wantNBooks       int
		wantLessonTitles []string
		wantCards        []kanjis.Card
	}{{
		name:             "ok",
		fileName:         filepath.Join(testDataDir, "kanjis1.csv"),
		csvImporter:      kanjiImporter,
		wantNBooks:       1,
		wantLessonTitles: []string{lessonID.Name},
		wantCards:        kanjis1,
	}, {
		name:             "file not found",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		csvImporter:      kanjiImporter,
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []kanjis.Card{},
	}, {
		name:             "missing importer",
		fileName:         filepath.Join(testDataDir, "kanjis1.csv"),
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []kanjis.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			importer := store.NewKanjiImporter(
				books.NewLibrary("testlib"), c.csvImporter)
			checkError(t, importer.Lesson(c.fileName, lessonID), c.wantErr)

			gotNbooks := len(importer.Library.SortedBookIDs())
			if gotNbooks != c.wantNBooks {
				t.Errorf("ERROR: got %v, want %v",
					gotNbooks, c.wantNBooks)
			}

			book := importer.Library.Book(bookID)
			if diff := cmp.Diff(book.LessonNames(), c.wantLessonTitles); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}

			lesson := book.Lesson(lessonID.Name)
			if diff := cmp.Diff(lesson.KanjiCards(), c.wantCards); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

// checkError checks for wanted or unwanted errors checks for wanted or unwanted errors.
func checkError(t *testing.T, err error, want bool) {
	t.Helper()

	if want {
		if err == nil {
			t.Fatalf("ERROR: wanted error not detected")
		}
		t.Logf("INFO: error message: %v", err)
		return
	}

	if err != nil {
		t.Errorf("ERROR: got error: %v", err)
	}
}

// words1 is the content of file testdata/word1.csv.
var words1 = []words.Card{{
	ID:          "1",
	Nihongo:     "先生",
	Kana:        "せんせい",
	Romaji:      "sensei",
	Meaning:     "Lehrer",
	Hint:        "andere Personen",
	Explanation: "für sich selbst anderer Ausdruck",
}, {
	ID:      "2",
	Nihongo: "医者",
	Kana:    "いしゃ",
	Romaji:  "isha",
	Meaning: "Arzt, Ärztin",
}, {
	ID:      "3",
	Nihongo: "お名前\u3000は\u3000「何\u3000です\u3000か」。",
	Kana:    "お\u3000なまえ\u3000は\u3000「なん\u3000です\u3000か」。",
	Romaji:  "onamae wa (nan desu ka).",
	Meaning: "Wie heißen Sie bitte?",
}, {
	ID:       "4",
	Nihongo:  "起きます",
	Kana:     "おきます",
	Romaji:   "okimasu",
	Meaning:  "aufstehen",
	DictForm: "おきる",
	TeForm:   "おきて",
	NaiForm:  "おきない",
}}

// kanjis1 is the content of file testdata/kanjis1.csv.
var kanjis1 = []kanjis.Card{{
	ID: "1", Kanji: '人',
	Explanation: "nicht: 入",
	Details: []kanjis.Detail{
		{Reading: "hito", Meanings: []string{"Mensch"}},
		{Reading: "NIN", Meanings: []string{"Mensch"}},
		{Reading: "JIN", Meanings: []string{"Mensch"}},
	},
}, {
	ID: "2", Kanji: '一',
	Details: []kanjis.Detail{
		{Reading: "hito.tsu", Meanings: []string{"eins"}},
		{Reading: "ICHI", Meanings: []string{"eins"}},
	},
}, {
	ID: "3", Kanji: '二',
	Hint: "auch ein kana",
	Details: []kanjis.Detail{
		{Reading: "futa", Meanings: []string{"zwei"}},
		{Reading: "futa.tsu", Meanings: []string{"zwei"}},
		{Reading: "NI", Meanings: []string{"zwei"}},
	},
}, {
	ID: "4", Kanji: '三',
	Details: []kanjis.Detail{
		{Reading: "mi", Meanings: []string{"drei"}},
		{Reading: "mi.ttsu", Meanings: []string{"drei"}},
		{Reading: "SAN", Meanings: []string{"drei"}},
	},
}, {
	ID: "5", Kanji: '日',
	Details: []kanjis.Detail{
		{Reading: "hi", Meanings: []string{"Tag", "Sonne"}},
		{Reading: "ka", Meanings: []string{"Tag", "Sonne"}},
		{Reading: "JITSU", Meanings: []string{"Tag", "Sonne"}},
		{Reading: "NICHI", Meanings: []string{"Tag", "Sonne"}},
	},
}, {
	ID: "6", Kanji: '四',
	Details: []kanjis.Detail{
		{Reading: "yon", Meanings: []string{"vier"}},
		{Reading: "yo.ttsu", Meanings: []string{"vier"}},
		{Reading: "SHI", Meanings: []string{"vier"}},
	},
}}
