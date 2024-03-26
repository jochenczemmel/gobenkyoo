package jsondb_test

import (
	"fmt"
	"path/filepath"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestClassroomLoad(t *testing.T) {
	baseDir := filepath.Join(testDataDir, "load")
	testCases := []struct {
		name               string
		dir                string
		roomName           string
		wantErr, wantFound bool
		want               learn.Classroom
	}{{
		name:      "ok",
		dir:       baseDir,
		roomName:  testClassroomName,
		wantFound: true,
		want:      makeLearnClassroom(),
	}, {
		name:     "classroom not found",
		dir:      baseDir,
		roomName: "does not exist",
	}, {
		name:     "dir not found",
		dir:      "does/not/exist",
		roomName: testClassroomName,
	}, {
		name:      "no permission",
		dir:       "/etc/sudoers.d",
		roomName:  testClassroomName,
		wantFound: true,
		wantErr:   true,
	}, {
		name:      "invalid json",
		dir:       baseDir,
		roomName:  "invalid",
		wantFound: true,
		wantErr:   true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			lib := jsondb.New(c.dir)
			got, found, err := lib.LoadClassroom(c.roomName)

			if found != c.wantFound {
				t.Errorf("ERROR: got %v, want %v", found, c.wantFound)
			}

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

			if !c.wantFound {
				return
			}

			compareClassrom(t, got, c.want)
		})
	}
}

func compareClassrom(t *testing.T, got, want learn.Classroom) {

	t.Run("name", func(t *testing.T) {
		if got.Name != want.Name {
			t.Errorf("ERROR: Name: got %q, want %q", got.Name, want.Name)
		}
	})

	testCases := []struct {
		name                string
		gotBoxes, wantBoxes []learn.Box
	}{
		{
			name:      "kanji",
			gotBoxes:  got.KanjiBoxes(),
			wantBoxes: want.KanjiBoxes(),
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			if len(c.gotBoxes) != len(c.wantBoxes) {
				t.Fatalf("ERROR: length KanjiBoxes: got %v, want %v",
					len(c.gotBoxes), len(c.wantBoxes))
			}

			// ensure same box order
			sort.Slice(c.gotBoxes, func(a, b int) bool {
				return c.gotBoxes[a].Name < c.gotBoxes[b].Name
			})
			sort.Slice(c.wantBoxes, func(a, b int) bool {
				return c.wantBoxes[a].Name < c.wantBoxes[b].Name
			})

			for i, gotBox := range c.gotBoxes {
				for _, mode := range gotBox.Modes() {
					for _, level := range learn.Levels() {

						t.Run(fmt.Sprintf("%d %s %s %d", i, gotBox.Name, mode, level),
							func(t *testing.T) {
								if diff := cmp.Diff(
									gotBox.Cards(mode, level),
									c.wantBoxes[i].Cards(mode, level),
								); diff != "" {
									t.Fatalf("ERROR: (%s/%v) got- want+\n%s", mode, level, diff)
								}
							})
					}
				}
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
	wb1.SetCardLevel(learn.Native2Japanese, 1, wb1.Cards(learn.Native2Japanese, 0)[1])

	boxID.Name = "box 2"
	boxID.LessonID.Name = "lesson 2"

	kb2 := learn.NewKanjiBox(boxID, kanjiCards2...)
	kb2.SetCardLevel(learn.Kanji2Native, 2, kb2.Cards(learn.Kanji2Native, 0)[1])

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
