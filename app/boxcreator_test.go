package app_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func TestCreateBoxFromList(t *testing.T) {

	creator := app.NewBoxCreator(
		jsondb.New(filepath.Join(testDataDir, jsondb.BaseDir)),
	)
	_, err := creator.Load(cfg.DefaultLibrary, cfg.DefaultClassroom)
	if err != nil {
		t.Fatalf("ERROR: test preparation failed: %v", err)
	}

	boxID := learn.BoxID{
		Name: "box 1",
		LessonID: books.LessonID{
			Name: "box 1",
			ID:   books.NewID("kanji learning", "", 0),
		},
	}
	fromID := books.NewID("minna kyu 1", "minna", 1)
	kanjiList := "日四五六七"
	creator.KanjiBoxFromList(kanjiList, fromID, boxID)

	wantCards := wantKanji2Cards
	gotCards := creator.Classroom.KanjiBox(boxID).
		Cards(learn.DefaultKanjiMode, learn.AllLevel)

	if diff := cmp.Diff(gotCards, wantCards); diff != "" {
		t.Errorf("ERROR: got- want+\n%s", diff)
	}
	// t.Logf("DEBUG: \n%#v\n", gotCards)
}

func TestCreateBoxFromLesson(t *testing.T) {

	bookID := books.NewID("minna kyu 1", "minna", 1)
	lessonID := books.LessonID{
		Name: "lesson 1",
		ID:   bookID,
	}
	boxID := learn.BoxID{
		Name:     "lesson 1",
		LessonID: lessonID,
	}

	testCases := []struct {
		name      string
		boxID     learn.BoxID
		testFunc  func(*app.BoxCreator, learn.BoxID) error
		getFunc   func(learn.Classroom, learn.BoxID) learn.Box
		mode      string
		wantErr   bool
		wantCards []learn.Card
	}{{
		name:      "kanji ok",
		boxID:     boxID,
		testFunc:  (*app.BoxCreator).KanjiBox,
		getFunc:   learn.Classroom.KanjiBox,
		mode:      learn.DefaultKanjiMode,
		wantCards: wantKanji1Cards,
	}, {
		name:      "kanji box not found",
		boxID:     learn.BoxID{},
		testFunc:  (*app.BoxCreator).KanjiBox,
		getFunc:   learn.Classroom.KanjiBox,
		mode:      learn.DefaultKanjiMode,
		wantCards: []learn.Card{},
		wantErr:   true,
	}, {
		name:      "word ok",
		boxID:     boxID,
		testFunc:  (*app.BoxCreator).WordBox,
		getFunc:   learn.Classroom.WordBox,
		mode:      learn.DefaultWordMode,
		wantCards: wantWord1Cards,
	}, {
		name:      "word box not found",
		boxID:     learn.BoxID{},
		testFunc:  (*app.BoxCreator).WordBox,
		getFunc:   learn.Classroom.WordBox,
		mode:      learn.DefaultWordMode,
		wantCards: []learn.Card{},
		wantErr:   true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			creator := app.NewBoxCreator(
				jsondb.New(filepath.Join(testDataDir, jsondb.BaseDir)),
			)
			_, _ = creator.Load(cfg.DefaultLibrary, cfg.DefaultClassroom)
			err := c.testFunc(&creator, c.boxID)
			checkError(t, err, c.wantErr)

			box := c.getFunc(creator.Classroom, c.boxID)
			gotCards := box.Cards(c.mode, learn.AllLevel)
			if diff := cmp.Diff(gotCards, c.wantCards); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

var lesson1ID = books.LessonID{
	Name: "lesson 1", ID: books.ID{
		Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
	},
}

var wantKanji1Cards = []learn.Card{{
	ID: learn.CardID{
		ContentID: "1",
		LessonID:  lesson1ID,
	},
	Question:    "人",
	Answer:      "Mensch",
	MoreAnswers: []string{"hito, NIN, JIN"},
}, {
	ID: learn.CardID{
		ContentID: "2",
		LessonID:  lesson1ID,
	},
	Question:    "一",
	Answer:      "eins",
	MoreAnswers: []string{"hito.tsu, ICHI"},
}, {
	ID: learn.CardID{
		ContentID: "3",
		LessonID:  lesson1ID,
	},
	Question:    "二",
	Answer:      "zwei",
	MoreAnswers: []string{"futa, futa.tsu, NI"},
}, {
	ID: learn.CardID{
		ContentID: "4",
		LessonID:  lesson1ID,
	},
	Question:    "三",
	Answer:      "drei",
	MoreAnswers: []string{"mi, mi.ttsu, SAN"},
}, {
	ID: learn.CardID{
		ContentID: "5",
		LessonID:  lesson1ID,
	},
	Question:    "日",
	Answer:      "Tag, Sonne",
	MoreAnswers: []string{"hi, ka, JITSU, NICHI"},
}, {
	ID: learn.CardID{
		ContentID: "6",
		LessonID:  lesson1ID,
	},
	Question:    "四",
	Answer:      "vier",
	MoreAnswers: []string{"yon, yo.ttsu, SHI"},
}}

var wantWord1Cards = []learn.Card{{
	ID: learn.CardID{
		ContentID: "1",
		LessonID:  lesson1ID,
	},
	Question:    "Lehrer",
	Hint:        "andere Personen",
	Answer:      "先生",
	MoreAnswers: []string{"せんせい", "sensei"},
	Explanation: "für sich selbst anderer Ausdruck",
}, {
	ID: learn.CardID{
		ContentID: "2",
		LessonID:  lesson1ID,
	},
	Question:    "Arzt, Ärztin",
	Answer:      "医者",
	MoreAnswers: []string{"いしゃ", "isha"},
}, {
	ID: learn.CardID{
		ContentID: "3",
		LessonID:  lesson1ID,
	},
	Question: "Wie heißen Sie bitte?",
	Answer:   "お名前\u3000は\u3000「何\u3000です\u3000か」。",
	MoreAnswers: []string{
		"お\u3000なまえ\u3000は\u3000「なん\u3000です\u3000か」。",
		"onamae wa (nan desu ka).",
	},
}, {
	ID: learn.CardID{
		ContentID: "4",
		LessonID:  lesson1ID,
	},
	Question: "aufstehen",
	Answer:   "起きます",
	MoreAnswers: []string{
		"おきます",
		"okimasu",
		"おきる",
		"おきて",
		"おきない",
	},
}}

var wantKanji2Cards = []learn.Card{
	{
		ID: learn.CardID{
			ContentID: "5",
			LessonID: books.LessonID{
				Name: "lesson 1",
				ID: books.ID{
					Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
				},
			},
		},
		Question:    "日",
		Answer:      "Tag, Sonne",
		MoreAnswers: []string{"hi, ka, JITSU, NICHI"},
	},
	{
		ID: learn.CardID{
			ContentID: "6",
			LessonID: books.LessonID{
				Name: "lesson 1",
				ID: books.ID{
					Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
				},
			},
		},
		Question:    "四",
		Answer:      "vier",
		MoreAnswers: []string{"yon, yo.ttsu, SHI"},
	},
	{
		ID: learn.CardID{
			ContentID: "1",
			LessonID: books.LessonID{
				Name: "lesson 2",
				ID: books.ID{
					Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
				},
			},
		},
		Question:    "五",
		Answer:      "fünf, 5",
		MoreAnswers: []string{"itsu.tsu, GO"},
	},
	{
		ID: learn.CardID{
			ContentID: "2",
			LessonID: books.LessonID{
				Name: "lesson 2",
				ID: books.ID{
					Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
				},
			},
		},
		Question:    "六",
		Answer:      "sechs, 6",
		MoreAnswers: []string{"mu.ttsu, ROKU"},
	},
	{
		ID: learn.CardID{
			ContentID: "3",
			LessonID: books.LessonID{
				Name: "lesson 2",
				ID: books.ID{
					Title: "minna kyu 1", SeriesTitle: "minna", Volume: 1,
				},
			},
		},
		Question:    "七",
		Answer:      "sieben, 7",
		MoreAnswers: []string{"nana, nana.tsu, SHICHI"},
	},
}
