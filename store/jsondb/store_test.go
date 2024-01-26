package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

var storePath = filepath.Join("testdata", "store")

func TestStoreLibrary(t *testing.T) {

	err := os.RemoveAll(storePath)
	if err != nil {
		t.Errorf("remove %q: got error: %v", storePath, err)
	}

	testBook1.AddKanjis(testLessonName1, kanjiCardsLesson1...)
	testBook1.AddWords(testLessonName1, wordCardsLesson1...)
	testBook1.AddKanjis(testLessonName2, kanjiCardsLesson2...)
	testBook1.AddWords(testLessonName2, wordCardsLesson2...)

	library := books.NewLibrary(testLibraryName)
	library.AddBooks(testBook1, testBook2)

	storer := jsondb.NewStorer(storePath)
	err = storer.StoreLibrary(library)
	if err != nil {
		t.Errorf("Store(): got error: %v", err)
	}

	jsonFile := filepath.Join(storePath, jsondb.LibraryPath,
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

func TestStoreClassroom(t *testing.T) {

	err := os.RemoveAll(storePath)
	if err != nil {
		t.Errorf("remove %q: got error: %v", storePath, err)
	}

	classroomName := "class 1"
	room := learn.NewClassroom(classroomName)
	boxID := learn.BoxID{
		Name: "box 1",
		LessonID: books.LessonID{
			Name: "lesson 1",
			ID: books.ID{
				Title:       "book 1",
				SeriesTitle: "book",
				Volume:      1,
			},
		},
	}
	room.NewWordBox(boxID, wordCards...)
	room.NewKanjiBox(boxID, kanjiCards...)
	boxID.Name = "box 2"
	room.NewWordBox(boxID)

	storer := jsondb.NewStorer(storePath)
	err = storer.StoreClassroom(room)
	if err != nil {
		t.Fatalf("StoreClassroom(): got error: %v", err)
	}
}

var wordCards = []words.Card{{
	ID:          1,
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
	ID:      2,
	Nihongo: "世界",
	Kana:    "せかい",
	Romaji:  "sekai",
	Meaning: "world",
}}

var kanjiCards = []kanjis.Card{{
	ID:    1,
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
	ID:    2,
	Kanji: '曜',
	Details: []kanjis.Detail{{
		Reading:  "yoo",
		Meanings: []string{"weekday"},
	}},
}}
