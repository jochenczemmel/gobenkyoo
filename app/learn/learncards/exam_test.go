package learncards_test

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

var cards1 = []learncards.Card{
	{ID: "card1"},
	{ID: "card2"},
	{ID: "card3"},
}

var cards2 = []learncards.Card{
	{ID: "card3"},
	{ID: "card4"},
}

var nAllCards = len(cards1) + len(cards2)

func TestExam(t *testing.T) {

	mode1 := "mode 1"
	mode2 := "mode 2"

	box1 := learncards.NewBox("lesson 1", "book 1")
	box1.Set(mode1, cards1...)
	box1.Set(mode2, cards1...)

	box2 := learncards.NewBox("lesson 2", "book 1")
	box2.Set(mode1, cards2...)
	box2.Set(mode2, cards1...)

	t.Run("cards", func(t *testing.T) {
		exam := learncards.NewExam(mode1, learncards.AllLevel, box1, box2)

		t.Run("number of cards", func(t *testing.T) {
			got := exam.NCards()
			want := nAllCards
			assertEquals(t, got, want)
		})

		t.Run("get cards", func(t *testing.T) {
			got := exam.Cards()
			want := append(cards1, cards2...)
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})

		t.Run("shuffle cards", func(t *testing.T) {
			want := append(cards1, cards2...)
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
	})

	t.Run("advance card", func(t *testing.T) {
		exam := learncards.NewExam(mode1, learncards.AllLevel, box1, box2)
		exam.Advance(cards1[1])
		testCases := []struct {
			level, want int
		}{
			{learncards.AllLevel, nAllCards},
			{learncards.MinLevel, nAllCards - 1},
			{learncards.MinLevel + 1, 1},
			{learncards.MinLevel + 2, 0},
		}

		for _, c := range testCases {
			t.Run("level "+strconv.Itoa(c.level), func(t *testing.T) {
				exam := learncards.NewExam(mode1, c.level, box1, box2)
				assertEquals(t, exam.NCards(), c.want)
			})
		}

		t.Run("box 1 ", func(t *testing.T) {
			testCases := []struct {
				name, mode  string
				level, want int
			}{
				{"min level", mode1, learncards.MinLevel, len(cards1) - 1},
				{"first level", mode1, learncards.MinLevel + 1, 1},
				{"second level", mode1, learncards.MinLevel + 2, 0},
				{"all level", mode1, learncards.AllLevel, len(cards1)},
				{"different mode min level", mode2, learncards.MinLevel, len(cards1)},
			}
			for _, c := range testCases {
				t.Run(c.name, func(t *testing.T) {
					got := box1.NCards(c.mode, c.level)
					assertEquals(t, got, c.want)
				})
			}
		})

		t.Run("box 2 unchanged", func(t *testing.T) {
			assertEquals(t, box2.NCards(mode1, learncards.MinLevel), len(cards2))
		})
	})
}
