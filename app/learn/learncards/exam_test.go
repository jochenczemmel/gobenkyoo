package learncards_test

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

func makeExam(level int, boxes ...learncards.Box) learncards.Exam {
	opt := learncards.ExamOptions{
		LearnMode: mode1,
		Level:     level,
		NoShuffle: true,
	}
	return learncards.NewExam(opt, boxes...)
}

func TestExamCards(t *testing.T) {
	box1, box2 := makeBoxes()
	exam := makeExam(learncards.AllLevel, box1, box2)

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
}

func TestExamShuffled(t *testing.T) {
	box1, box2 := makeBoxes()

	want := append(cards1, cards2...)
	nTries := 10
	for i := 0; i < nTries; i++ {
		exam := learncards.NewExam(
			learncards.ExamOptions{
				LearnMode: mode1,
				Level:     learncards.AllLevel,
			},
			box1, box2)
		got := exam.Cards()
		if diff := cmp.Diff(got, want); diff != "" {
			t.Logf("DEBUG: %d", i)
			return
		}
	}
	t.Errorf("ERROR: shuffle returned %d times the ordered cards", nTries)
}

func TestExamAdvance(t *testing.T) {
	box1, box2 := makeBoxes()
	exam := makeExam(learncards.AllLevel, box1, box2)

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
	t.Run("first pass", func(t *testing.T) {

		for _, c := range testCases {
			t.Run("level "+strconv.Itoa(c.level), func(t *testing.T) {
				exam := makeExam(c.level, box1, box2)
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
	t.Run("second pass", func(t *testing.T) {

		for _, c := range testCases {
			t.Run("level "+strconv.Itoa(c.level), func(t *testing.T) {
				exam := makeExam(c.level, box1, box2)
				assertEquals(t, exam.NCards(), c.wantSecond)
			})
		}
	})
}

func TestExamReset(t *testing.T) {
	box1, box2 := makeBoxes()
	card := cards1[1]
	exam := makeExam(learncards.MinLevel, box1, box2)
	exam.Advance(card)
	exam.Advance(card)

	exam.Reset(card)
	exam = makeExam(learncards.MinLevel, box1, box2)

	got := exam.Cards()
	want := append(cards1, cards2...)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("ERROR: -got +want\n%s", diff)
	}
}

func TestExamNextCard(t *testing.T) {
	box1, box2 := makeBoxes()
	exam := makeExam(learncards.MinLevel, box1, box2)
	want := append(cards1, cards2...)

	for i := 0; i < len(want); i++ {
		got, ok := exam.NextCard()
		t.Run("card "+strconv.Itoa(i), func(t *testing.T) {
			assertEquals(t, ok, true)
			if diff := cmp.Diff(got, want[i]); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})
	}

	got, ok := exam.NextCard()
	assertEquals(t, got.ID, "")
	assertEquals(t, ok, false)
}

func TestExamPassFail(t *testing.T) {
	box1, _ := makeBoxes()
	level := learncards.MinLevel + 1
	for _, card := range cards1 {
		box1.SetCardLevel(mode1, card, level)
	}
	exam := learncards.NewExam(
		learncards.ExamOptions{
			LearnMode: mode1,
			Level:     level,
			NoShuffle: true,
		},
		box1)

	exam.NextCard()
	exam.Pass()

	t.Run("one card in next level", func(t *testing.T) {
		assertEquals(t, box1.NCards(mode1, learncards.MinLevel), 0)
		assertEquals(t, box1.NCards(mode1, level), 2)
		assertEquals(t, box1.NCards(mode1, level+1), 1)
	})

	exam.NextCard()
	exam.Fail()

	t.Run("one card in min level", func(t *testing.T) {
		assertEquals(t, box1.NCards(mode1, learncards.MinLevel), 1)
		assertEquals(t, box1.NCards(mode1, level), 1)
		assertEquals(t, box1.NCards(mode1, level+1), 1)
	})
}

func TestExamKeepLevel(t *testing.T) {
	box1, _ := makeBoxes()
	level := learncards.MinLevel + 1
	for _, card := range cards1 {
		box1.SetCardLevel(mode1, card, level)
	}
	exam := learncards.NewExam(
		learncards.ExamOptions{
			LearnMode: mode1,
			Level:     level,
			NoShuffle: true,
			KeepLevel: true,
		},
		box1)

	exam.NextCard()
	exam.Pass()

	t.Run("one card not in next level", func(t *testing.T) {
		assertEquals(t, box1.NCards(mode1, learncards.MinLevel), 0)
		assertEquals(t, box1.NCards(mode1, level), 3)
		assertEquals(t, box1.NCards(mode1, level+1), 0)
	})

	exam.NextCard()
	exam.Fail()

	t.Run("one card not in min level", func(t *testing.T) {
		assertEquals(t, box1.NCards(mode1, learncards.MinLevel), 0)
		assertEquals(t, box1.NCards(mode1, level), 3)
		assertEquals(t, box1.NCards(mode1, level+1), 0)
	})
}

func TestExamRepeat(t *testing.T) {

	t.Run("no more cards", func(t *testing.T) {
		box1, _ := makeBoxes()
		exam := learncards.NewExam(
			learncards.ExamOptions{
				LearnMode: mode1,
				Level:     learncards.MinLevel,
				NoShuffle: true,
			},
			box1)

		for range cards1 {
			exam.NextCard()
			exam.Fail()
		}

		_, ok := exam.NextCard()
		assertEquals(t, ok, false)
	})

	t.Run("repeat cards", func(t *testing.T) {
		box1, _ := makeBoxes()
		exam := learncards.NewExam(
			learncards.ExamOptions{
				LearnMode: mode1,
				Level:     learncards.MinLevel,
				NoShuffle: true,
				Repeat:    true,
			},
			box1)

		for range cards1 {
			exam.NextCard()
			exam.Fail()
		}

		got, ok := exam.NextCard()
		assertEquals(t, ok, true)
		assertEquals(t, got.ID, cards1[0].ID)

		// NCards returns the number of initial cards
		assertEquals(t, exam.NCards(), len(cards1))
	})
}
