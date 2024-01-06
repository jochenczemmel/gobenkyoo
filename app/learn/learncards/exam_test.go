package learncards_test

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

var cards = []learncards.Card{
	{ID: "card1"},
	{ID: "card2"},
	{ID: "card3"},
}

var cards2 = []learncards.Card{
	{ID: "card3"},
	{ID: "card4"},
}

var nAllCards = len(cards) + len(cards2)

func TestExam(t *testing.T) {

	mode1 := "mode 1"
	mode2 := "mode 2"

	box1 := learncards.NewBox("lesson 1", "book 1")
	box1.Set(mode1, cards...)
	box1.Set(mode2, cards...)

	box2 := learncards.NewBox("lesson 2", "book 1")
	box2.Set(mode1, cards2...)
	box2.Set(mode2, cards...)

	t.Run("cards", func(t *testing.T) {
		exam := learncards.NewExam(mode1, learncards.AllLevel, box1, box2)

		t.Run("number of cards", func(t *testing.T) {
			got := exam.NCards()
			want := nAllCards
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
	})

	t.Run("advance card", func(t *testing.T) {
		exam := learncards.NewExam(mode1, learncards.AllLevel, box1, box2)
		exam.Advance(cards[1])
		testCases := []struct {
			level, want int
		}{
			{level: learncards.AllLevel, want: nAllCards},
			{level: learncards.MinLevel, want: nAllCards - 1},
			{level: learncards.MinLevel + 1, want: 1},
			{level: learncards.MinLevel + 2, want: 0},
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
				{
					name:  "min level",
					mode:  mode1,
					level: learncards.MinLevel,
					want:  2,
				},
				{
					name:  "first level",
					mode:  mode1,
					level: learncards.MinLevel + 1,
					want:  1,
				},
				{
					name:  "second level",
					mode:  mode1,
					level: learncards.MinLevel + 2,
					want:  0,
				},
				{
					name:  "all level",
					mode:  mode1,
					level: learncards.AllLevel,
					want:  3,
				},
				{
					name:  "different mode min level",
					mode:  mode2,
					level: learncards.MinLevel,
					want:  3,
				},
			}
			for _, c := range testCases {
				t.Run(c.name, func(t *testing.T) {
					got := box1.NCards(c.mode, c.level)
					assertEquals(t, got, c.want)
				})
			}
		})
	})
}
