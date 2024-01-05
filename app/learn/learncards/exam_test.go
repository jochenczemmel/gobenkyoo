package learncards_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

var cards2 = []learncards.Card{
	{ID: "card3"},
	{ID: "card4"},
}

func TestExamCards(t *testing.T) {
	mode := "mode 1"
	level := learncards.AllLevel
	box1 := learncards.NewBox("lesson 1", "book 1")
	box1.Set(mode, cards...)
	box2 := learncards.NewBox("lesson 2", "book 1")
	box2.Set(mode, cards2...)
	exam := learncards.NewExam(mode, level, box1, box2)
	t.Run("number of cards", func(t *testing.T) {
		got := exam.NCards()
		want := 5
		assertEquals(t, got, want)
	})
	t.Run("get cards", func(t *testing.T) {
		got := exam.Cards()
		want := append(cards, cards2...)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: -got +want\n%s", diff)
		}
	})
	t.Run("shuffle cards", func(t *testing.T) {
		want := append(cards, cards2...)
		nTries := 10
		for i := 0; i < nTries; i++ {
			exam.Shuffle()
			got := exam.Cards()
			if diff := cmp.Diff(got, want); diff != "" {
				t.Logf("DEBUG: %d", i)
				// t.SkipNow()
				return
			}
		}
		t.Errorf("ERROR: shuffle returned %d times the ordered cards", nTries)
	})
}
