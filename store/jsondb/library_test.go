package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

const (
	testDataDir = "testdata"
)

func TestLibraryStore(t *testing.T) {

	bookLib := makeBooksLibrary()
	storeDir := filepath.Join(testDataDir, "store")
	err := os.RemoveAll(storeDir)
	if err != nil {
		t.Fatalf("ERROR: remove store dir failed: %v", err)
	}

	testCases := []struct {
		name    string
		dir     string
		wantErr bool
	}{
		{
			name:    "ok",
			dir:     storeDir,
			wantErr: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			lib := jsondb.NewLibrary(c.dir, bookLib)
			err := lib.Store()
			if err != nil {
				t.Errorf("ERROR: got error %v", err)
			}
		})
	}
}

func makeBooksLibrary() books.Library {

	lesson1 := books.NewLesson("lesson 1")
	lesson1.AddWords(wordCardsLesson1...)
	lesson1.AddKanjis(kanjiCardsLesson1...)
	lesson2 := books.NewLesson("lesson 2")
	lesson2.AddWords(wordCardsLesson2...)
	lesson2.AddKanjis(kanjiCardsLesson2...)
	lesson3 := books.NewLesson("lesson 3")
	lesson3.AddKanjis(kanjiCardsLesson3...)

	book1 := books.New(books.ID{
		Title:       "minna no nihongo sho 1",
		SeriesTitle: "minna no nihongo",
		Volume:      1,
	})
	book1.SetLessons(lesson1, lesson2)
	book2 := books.New(books.ID{
		Title:       "minna no nihongo sho 2",
		SeriesTitle: "minna no nihongo",
		Volume:      2,
	})
	book2.SetLessons(lesson3)

	library := books.NewLibrary("meine Bücher")
	library.AddBooks(book1, book2)

	return library
}

var kanjiCardsLesson1 = []kanjis.Card{{
	Kanji: '方',
	Details: []kanjis.Detail{{
		Reading:     "HOO",
		ReadingKana: "ホー",
		Meanings:    []string{"Richtung", "Art und Weise, etwas zu tun"},
	}, {
		Reading:     "kata",
		ReadingKana: "かた",
		Meanings:    []string{"Person", "Art und Weise, etwas zu tun"},
	}},
}, {
	Kanji: '曜',
	Details: []kanjis.Detail{
		{Reading: "yoo", Meanings: []string{"Wochentag"}},
	},
}}

var wordCardsLesson1 = []words.Card{{
	Nihongo: "日曜日",
	Kana:    "にちようび",
	Romaji:  "nichyoobi",
	Meaning: "Sonntag",
}, {
	Nihongo: "月曜日",
	Kana:    "げつようび",
	Romaji:  "getsuyoobi",
	Meaning: "Montag",
}, {
	Nihongo: "火曜日",
	Kana:    "かようび",
	Romaji:  "kayoobi",
	Meaning: "Dienstag",
}}

var kanjiCardsLesson2 = []kanjis.Card{{
	Kanji: '習',
	Details: []kanjis.Detail{{
		Reading:     "narai(u)",
		ReadingKana: "なら（う）",
		Meanings:    []string{"lernen"},
	}},
}, {
	Kanji: '世',
	Details: []kanjis.Detail{{
		Reading:  "se",
		Meanings: []string{"Welt", "Ära", "Generation"},
	}},
}, {
	Kanji: '界',
	Details: []kanjis.Detail{{
		Reading:  "kai",
		Meanings: []string{"alle"},
	}},
}}

var wordCardsLesson2 = []words.Card{{
	Nihongo:     "習います",
	Kana:        "ならいます",
	Romaji:      "naraimasu",
	Meaning:     "lernen",
	DictForm:    "習う",
	TeForm:      "習って",
	NaiForm:     "習わない",
	Hint:        "von jemandem",
	Explanation: "studieren ist benkyoo (勉強)",
}, {
	Nihongo: "世界",
	Kana:    "せかい",
	Romaji:  "sekai",
	Meaning: "Welt",
}}

var kanjiCardsLesson3 = []kanjis.Card{{
	Kanji: '外',
	Details: []kanjis.Detail{{
		Reading:     "soto",
		ReadingKana: "そと",
		Meanings:    []string{"außen", "draußen"},
	}},
}}
