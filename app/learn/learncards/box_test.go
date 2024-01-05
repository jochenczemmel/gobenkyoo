package learncards_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

var cards = []learncards.Card{
	{ID: "card1"},
	{ID: "card2"},
	{ID: "card3"},
}

func TestBoxNew(t *testing.T) {

	title := "lesson 1"
	bookTitle := "minna 1"
	box := learncards.NewBox(title, bookTitle)

	t.Run("titles", func(t *testing.T) {
		assertEquals(t, box.Title, title)
		assertEquals(t, box.BookTitle, bookTitle)
	})

	mode := "mode 1"
	box.Set(mode, cards...)

	t.Run("modes", func(t *testing.T) {
		got := box.Modes()
		want := []string{mode}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: -got +want\n%s", diff)
		}
	})

	t.Run("number of cards", func(t *testing.T) {
		got := box.AllCards(mode)
		want := 3
		assertEquals(t, len(got), want)
	})

	t.Run("number of cards invalid mode", func(t *testing.T) {
		got := box.AllCards("invalid")
		want := 0
		assertEquals(t, len(got), want)
	})
}

func assertEquals[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}
