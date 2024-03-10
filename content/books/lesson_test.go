package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestLessonUninitialized(t *testing.T) {

	t.Run("kanjis", func(t *testing.T) {
		lesson := books.Lesson{Name: "lesson 1"}
		want := []kanjis.Card{}
		got := lesson.KanjiCards()
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}

		lesson.AddWords(wordCards...)
		lesson.AddKanjis(kanjiCards...)
		got = lesson.KanjiCards()
		want = kanjiCards
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}

		// add again
		doubleKanjiCards := []kanjis.Card{
			{ID: "1", Kanji: '方'},
			{ID: "2", Kanji: '世'},
			{ID: "3", Kanji: '界'},
			{ID: "4", Kanji: '日'},
			{ID: "5", Kanji: '本'},
			{ID: "6", Kanji: '方'},
			{ID: "7", Kanji: '世'},
			{ID: "8", Kanji: '界'},
			{ID: "9", Kanji: '日'},
			{ID: "10", Kanji: '本'},
		}

		lesson.AddKanjis(kanjiCards...)
		got = lesson.KanjiCards()
		want = doubleKanjiCards
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})

	t.Run("words", func(t *testing.T) {

		lesson := books.Lesson{Name: "lesson 1"}
		want := []words.Card{}
		got := lesson.WordCards()
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}

		lesson.AddWords(wordCards...)
		lesson.AddKanjis(kanjiCards...)
		got = lesson.WordCards()
		want = wordCards
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}

		// add again
		doubleWordCards := []words.Card{
			{ID: "1", Nihongo: "世界"},
			{ID: "2", Nihongo: "日本"},
			{ID: "3", Nihongo: "白鳳"},
			{ID: "4", Nihongo: "大相撲"},
			{ID: "5", Nihongo: "福岡"},
			{ID: "6", Nihongo: "世界"},
			{ID: "7", Nihongo: "日本"},
			{ID: "8", Nihongo: "白鳳"},
			{ID: "9", Nihongo: "大相撲"},
			{ID: "10", Nihongo: "福岡"},
		}

		lesson.AddWords(wordCards...)
		lesson.AddKanjis(kanjiCards...)
		got = lesson.WordCards()
		want = doubleWordCards
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})
}
