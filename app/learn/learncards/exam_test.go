package learncards_test

//
// TODO: add method Next()
// TODO: add method Previous() ?
// TODO: add revisit of card that was not known
// TODO: disable shuffeling after first call to Next()
//

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

// static data for test fixture setup
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

var (
	mode1 = "mode 1"
	mode2 = "mode 2"
)

func makeBoxes() (learncards.Box, learncards.Box) {
	box1 := learncards.NewBox("lesson 1", "book 1")
	box1.Set(mode1, cards1...)
	box1.Set(mode2, cards1...)

	box2 := learncards.NewBox("lesson 2", "book 1")
	box2.Set(mode1, cards2...)
	box2.Set(mode2, cards2...)
	return box1, box2
}

func TestExamCards(t *testing.T) {

	box1, box2 := makeBoxes()
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
}

func TestExamAdvance(t *testing.T) {

	box1, box2 := makeBoxes()

	exam := learncards.NewExam(mode1, learncards.AllLevel, box1, box2)
	testCases := []struct {
		level, wantFirst, wantSecond int
	}{
		{learncards.AllLevel, nAllCards, nAllCards},
		{learncards.MinLevel, nAllCards - 1, nAllCards - 1},
		{learncards.MinLevel + 1, 1, 0},
		{learncards.MinLevel + 2, 0, 1},
		{learncards.MinLevel + 3, 0, 0},
	}

	exam.Advance(cards1[1])
	t.Run("first advance", func(t *testing.T) {

		for _, c := range testCases {
			t.Run("level "+strconv.Itoa(c.level), func(t *testing.T) {
				exam := learncards.NewExam(mode1, c.level, box1, box2)
				assertEquals(t, exam.NCards(), c.wantFirst)
			})
		}
	})

	t.Run("box 1", func(t *testing.T) {
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

	exam.Advance(cards1[1])
	t.Run("second advance", func(t *testing.T) {

		for _, c := range testCases {
			t.Run("level "+strconv.Itoa(c.level), func(t *testing.T) {
				exam := learncards.NewExam(mode1, c.level, box1, box2)
				assertEquals(t, exam.NCards(), c.wantSecond)
			})
		}
	})
}

func TestExamReset(t *testing.T) {
	box1, box2 := makeBoxes()
	card := cards1[1]
	exam := learncards.NewExam(mode1, learncards.MinLevel, box1, box2)
	exam.Advance(card)
	exam.Advance(card)

	exam.Reset(card)
	exam = learncards.NewExam(mode1, learncards.MinLevel, box1, box2)
	got := exam.Cards()
	want := append(cards1, cards2...)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("ERROR: -got +want\n%s", diff)
	}
}
