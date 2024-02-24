package jsondb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestClassroomStore(t *testing.T) {

	classroom := makeLearnClassroom()
	baseDir := filepath.Join(testDataDir, "store")
	err := os.RemoveAll(baseDir)
	if err != nil {
		t.Fatalf("ERROR: remove store dir failed: %v", err)
	}

	testCases := []struct {
		name    string
		dir     string
		wantErr bool
	}{{
		name:    "ok",
		dir:     baseDir,
		wantErr: false,
	}, {
		name:    "not ok",
		dir:     "/can/not/create",
		wantErr: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			lib := jsondb.New(c.dir)
			err := lib.StoreClassroom(classroom)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ERROR: error not detected")
				}
				t.Logf("INFO: got error: %v", err)
				return
			}
			if err != nil {
				t.Errorf("ERROR: got error %v", err)
			}
		})
	}
}

const testClassroomName = "VHS japanisch"

func makeLearnClassroom() learn.Classroom {
	room := learn.NewClassroom(testClassroomName)
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

	kb1 := learn.NewKanjiBox(boxID, kanjiCards1...)
	wb1 := learn.NewWordBox(boxID, wordCards1...)

	boxID.Name = "box 2"
	boxID.LessonID.Name = "lesson 2"

	kb2 := learn.NewKanjiBox(boxID, kanjiCards2...)

	room.SetKanjiBoxes(kb1, kb2)
	room.SetWordBoxes(wb1)

	return room
}

var wordCards1 = []words.Card{{
	ID:          "1",
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
	ID:      "2",
	Nihongo: "世界",
	Kana:    "せかい",
	Romaji:  "sekai",
	Meaning: "world",
}}

var kanjiCards1 = []kanjis.Card{{
	ID:    "1",
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
	ID:    "2",
	Kanji: '曜',
	Details: []kanjis.Detail{{
		Reading:  "yoo",
		Meanings: []string{"weekday"},
	}},
}}

var kanjiCards2 = []kanjis.Card{{
	ID:    "3",
	Kanji: '本',
	Details: []kanjis.Detail{{
		Reading:     "moto",
		ReadingKana: "もと",
		Meanings:    []string{"Wurzel", "Ursprung"},
	}},
}, {
	ID:    "4",
	Kanji: '木',
	Details: []kanjis.Detail{{
		Reading:  "ki",
		Meanings: []string{"Baum", "Holz"},
	}},
}, {
	ID:    "5",
	Kanji: '水',
	Details: []kanjis.Detail{{
		Reading:  "mizu",
		Meanings: []string{"Wasser"},
	}},
}}
