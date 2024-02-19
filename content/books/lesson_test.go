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
	})
}
