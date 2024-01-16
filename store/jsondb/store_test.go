package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestStoreBook(t *testing.T) {

	book1.AddKanjis(lesson1, kanjiCardsLesson1...)
	book1.AddWords(lesson1, wordCardsLesson1...)
	book1.AddKanjis(lesson2, kanjiCardsLesson2...)
	book1.AddWords(lesson2, wordCardsLesson2...)

	library := books.NewLibrary(libName)
	library.AddBooks(book1, book2)

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

	jsonFile := filepath.Join(path, "library", libName+".json")
	got, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("ReadFile(%v): got error: %v", jsonFile, err)
	}

	wantFile := filepath.Join("testdata", "want_japanology.json")
	want, err := os.ReadFile(wantFile)
	if err != nil {
		t.Fatalf("ReadFile(%v): got error: %v", wantFile, err)
	}

	if diff := cmp.Diff(string(got), string(want)); diff != "" {
		t.Errorf("stored file: -got +want\n%s", diff)
	}
}

var (
	libName = "japanology"
	book1   = books.New("minna no nihongo sho 1", "minna no nihongo", 1)
	book2   = books.New("minna no nihongo sho 2", "minna no nihongo", 2)
	lesson1 = "lesson1"
	lesson2 = "lesson2"
)

var kanjiCardsLesson1 = []kanjis.Card{
	kanjis.NewBuilder('方').
		AddDetailsWithKana("HOO", "ホー", "Richtung", "Art und Weise, etwas zu tun").
		AddDetailsWithKana("kata", "かた", "Person", "Art und Weise, etwas zu tun").
		Build(),
	kanjis.NewBuilder('曜').AddDetails("yoo", "weekday").Build(),
}

var wordCardsLesson1 = []words.Card{{
	Nihongo: "日曜日",
	Kana:    "にちようび",
	Romaji:  "nichyoobi",
	Meaning: "sunday",
}, {
	Nihongo: "月曜日",
	Kana:    "げつようび",
	Romaji:  "getsuyoobi",
	Meaning: "monday",
}, {
	Nihongo: "火曜日",
	Kana:    "かようび",
	Romaji:  "kayoobi",
	Meaning: "tuesday",
}}

var kanjiCardsLesson2 = []kanjis.Card{
	kanjis.NewBuilder('習').
		AddDetailsWithKana("narai(u)", "なら（う）", "learn").Build(),
	kanjis.NewBuilder('世').AddDetails("se", "world", "era", "generation").Build(),
	kanjis.NewBuilder('界').AddDetails("kai", "all").Build(),
}

var wordCardsLesson2 = []words.Card{{
	Nihongo:     "習います",
	Kana:        "ならいます",
	Romaji:      "naraimasu",
	Meaning:     "to learn",
	DictForm:    "習う",
	TeForm:      "習って",
	NaiForm:     "習わない",
	Hint:        "from somebody",
	Explanation: "to study is benkyoo (勉強)",
}, {
	Nihongo: "世界",
	Kana:    "せかい",
	Romaji:  "sekai",
	Meaning: "world",
}}
