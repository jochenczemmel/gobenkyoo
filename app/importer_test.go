package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

const (
	testDataDir = "testdata"
	storeDir    = "store"
)

var outDir = filepath.Join(testDataDir, storeDir, jsondb.BaseDir)

func TestImportWordLesson(t *testing.T) {

	err := os.RemoveAll(outDir)
	if err != nil {
		t.Fatalf("ERROR: test setup failed: %v", err)
	}

	bookID := books.ID{Title: "minna"}
	lessonID := books.LessonID{Name: "lesson1", ID: bookID}
	importer := csvimport.NewWord(',', true,
		[]string{"KANA", "NIHONGO", "ROMAJI", "MEANING", "HINT",
			"EXPLANATION", "DICTFORM", "TEFORM", "NAIFORM"})

	testCases := []struct {
		name             string
		fileName         string
		importer         app.WordImporter
		wantErr          bool
		wantNBooks       int
		wantLessonTitles []string
		wantCards        []words.Card
	}{{
		name:             "ok",
		fileName:         filepath.Join(testDataDir, "words1.csv"),
		importer:         importer,
		wantNBooks:       1,
		wantLessonTitles: []string{lessonID.Name},
		wantCards:        words1,
	}, {
		name:             "file not found",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		importer:         importer,
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

			importer := app.NewLibraryImporter(cfg.DefaultLibrary,
				jsondb.New(outDir))

			err := importer.Word(c.importer, c.fileName, lessonID)
			checkError(t, err, c.wantErr)
		})
	}
}

func TestImportKanjiLesson(t *testing.T) {

	err := os.RemoveAll(outDir)
	if err != nil {
		t.Fatalf("ERROR: test setup failed: %v", err)
	}

	bookID := books.ID{Title: "kanjidic"}
	lessonID := books.LessonID{Name: "lesson1", ID: bookID}
	importer := csvimport.NewKanji(';', '/', true,
		[]string{"kanji", "", "", "reading", "meanings",
			"hint", "explanation"})

	testCases := []struct {
		name             string
		fileName         string
		importer         app.KanjiImporter
		wantErr          bool
		wantNBooks       int
		wantLessonTitles []string
		wantCards        []kanjis.Card
	}{{
		name:             "ok",
		fileName:         filepath.Join(testDataDir, "kanjis1.csv"),
		importer:         importer,
		wantNBooks:       1,
		wantLessonTitles: []string{lessonID.Name},
		wantCards:        kanjis1,
	}, {
		name:             "file not found",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		importer:         importer,
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []kanjis.Card{},
	}, {
		name:             "missing importer",
		fileName:         filepath.Join(testDataDir, "does not exist"),
		wantErr:          true,
		wantNBooks:       0,
		wantLessonTitles: []string{},
		wantCards:        []kanjis.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			importer := app.NewLibraryImporter(cfg.DefaultLibrary,
				jsondb.New(outDir))
			err := importer.Kanji(c.importer, c.fileName, lessonID)
			checkError(t, err, c.wantErr)

		})
	}
}

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
