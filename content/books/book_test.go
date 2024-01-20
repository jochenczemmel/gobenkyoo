package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var wordCards = []words.Card{
	{ID: 1, Nihongo: "世界"},
	{ID: 2, Nihongo: "日本"},
	{ID: 3, Nihongo: "白鳳"},
	{ID: 4, Nihongo: "大相撲"},
	{ID: 5, Nihongo: "福岡"},
}

var kanjiCards = []kanjis.Card{
	kanjis.Card{ID: 1, Kanji: '方'},
	kanjis.Card{ID: 2, Kanji: '世'},
	kanjis.Card{ID: 6, Kanji: '界'},
	kanjis.Card{ID: 7, Kanji: '日'},
	kanjis.Card{ID: 8, Kanji: '本'},
}

func TestBookGetCard(t *testing.T) {
	lesson := "lesson 1"
	book := books.New(books.ID{})
	book.AddWords(lesson, wordCards...)
	book.AddKanjis(lesson, kanjiCards...)

	testCases := []struct {
		name          string
		lesson        string
		id            int
		wantKanjiCard kanjis.Card
		wantWordCard  words.Card
	}{{
		name:          "word and kanji",
		lesson:        lesson,
		id:            1,
		wantWordCard:  wordCards[0],
		wantKanjiCard: kanjiCards[0],
	}, {
		name:         "only word",
		lesson:       lesson,
		id:           5,
		wantWordCard: wordCards[4],
	}, {
		name:          "only kanji",
		lesson:        lesson,
		id:            8,
		wantKanjiCard: kanjiCards[4],
	}, {
		name:   "no match",
		lesson: lesson,
		id:     42,
	}, {
		name:   "wrong lesson",
		lesson: "not found",
		id:     1,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("word", func(t *testing.T) {
				got := book.FindWordCard(c.lesson, c.id)
				if diff := cmp.Diff(got, c.wantWordCard); diff != "" {
					t.Errorf("FindWordCard(%v): -got, +want\n%s", c.id, diff)
				}
			})
			t.Run("kanji", func(t *testing.T) {
				got := book.FindKanjiCard(c.lesson, c.id)
				if diff := cmp.Diff(got, c.wantKanjiCard); diff != "" {
					t.Errorf("FindKanjiCard(%v): -got, +want\n%s", c.id, diff)
				}
			})
		})
	}
}

func TestBookLessons(t *testing.T) {
	book := books.New(books.ID{})
	for _, lesson := range []string{"l1", "l2"} {
		book.AddKanjis(lesson)
	}
	for _, lesson := range []string{"l2", "l3"} {
		book.AddWords(lesson)
	}
	want := []string{"l1", "l2", "l3"}
	got := book.LessonTitles()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Lessons(): -got +want\n%s", diff)
	}
}

func TestBookAdd(t *testing.T) {
	id := books.ID{
		Title:       "minna1",
		SeriesTitle: "minna",
		Volume:      1,
	}
	book := books.New(id)

	lesson := "lesson 1"
	missingLesson := "not existing"

	t.Run("add kanjis", func(t *testing.T) {

		book.AddKanjis(lesson, kanjiCards[:3]...)
		got := book.KanjisFor(lesson)
		if diff := cmp.Diff(got, kanjiCards[:3]); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}

		book.AddKanjis(lesson, kanjiCards[3:]...)
		got = book.KanjisFor(lesson)
		if diff := cmp.Diff(got, kanjiCards); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}

		got = book.KanjisFor(missingLesson)
		if diff := cmp.Diff(got, kanjiCards[:0]); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}
	})

	t.Run("add words", func(t *testing.T) {

		book.AddWords(lesson, wordCards[:3]...)
		got := book.WordsFor(lesson)
		if diff := cmp.Diff(got, wordCards[:3]); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}

		book.AddWords(lesson, wordCards[3:]...)
		got = book.WordsFor(lesson)
		if diff := cmp.Diff(got, wordCards); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}

		got = book.WordsFor(missingLesson)
		if diff := cmp.Diff(got, wordCards[:0]); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}
	})
}
