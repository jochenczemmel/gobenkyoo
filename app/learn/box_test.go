package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func assertEquals[T comparable](tb testing.TB, got, want T) {
	tb.Helper()
	if got != want {
		tb.Errorf("ERROR: got %v, want %v", got, want)
	}
}

func TestBoxModes(t *testing.T) {
	testCases := []struct {
		name string
		box  learn.Box
		want []string
	}{
		{
			name: "kanji",
			box:  learn.NewKanjiBox(learn.BoxID{}),
			want: learn.KanjiModes(),
		},
		{
			name: "words",
			box:  learn.NewWordBox(learn.BoxID{}),
			want: learn.WordModes(),
		},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.box.Modes()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

func TestBoxSetCardLevel(t *testing.T) {
	box := learn.NewKanjiBox(learn.BoxID{})
	boxMode := learn.Kanji2Native
	initialLevel := learn.MinLevel + 1
	newLevel := initialLevel + 1

	box.AddCards(boxMode, initialLevel, cards1...)

	// 3 cards in initial, 0 cards in new level
	assertIntEqual(t, box.NCards(boxMode, initialLevel), len(cards1))
	assertIntEqual(t, box.NCards(boxMode, newLevel), 0)

	// set one card to the next level
	box.SetCardLevel(boxMode, newLevel, cards1[0])

	// 2 cards in initial, 1 card in new level
	assertIntEqual(t, box.NCards(boxMode, initialLevel), len(cards1)-1)
	assertIntEqual(t, box.NCards(boxMode, newLevel), 1)

	// invalid mode, nothing happens:
	box.SetCardLevel("invalid", newLevel, cards1[1])

	// 2 cards in initial, 1 card in new level
	assertIntEqual(t, box.NCards(boxMode, initialLevel), len(cards1)-1)
	assertIntEqual(t, box.NCards(boxMode, newLevel), 1)
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}

func TestBoxCards(t *testing.T) {

	box := learn.NewKanjiBox(learn.BoxID{})
	boxMode := learn.Kanji2Native
	box.AddCards(boxMode, learn.MinLevel, cards1...)
	box.AddCards("invalid mode", learn.MinLevel, cards2...)

	testCases := []struct {
		name        string
		mode        string
		level, want int
		wantCards   []learn.Card
	}{{
		name:      "all cards",
		mode:      boxMode,
		level:     learn.AllLevel,
		want:      len(cards1),
		wantCards: cards1,
	}, {
		name:      "first level",
		mode:      boxMode,
		level:     learn.MinLevel,
		want:      len(cards1),
		wantCards: cards1,
	}, {
		name:      "next level",
		mode:      boxMode,
		level:     learn.MinLevel + 1,
		want:      0,
		wantCards: []learn.Card{},
	}, {
		name:      "unknown mode",
		mode:      "unknown",
		level:     learn.AllLevel,
		want:      0,
		wantCards: []learn.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := box.NCards(c.mode, c.level)
			assertEquals(t, got, c.want)
			gotCards := box.Cards(c.mode, c.level)
			if diff := cmp.Diff(gotCards, c.wantCards); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

func TestBoxSetLevelLimits(t *testing.T) {
	// SetCardLevel() is also tested when testing exam.Advance()
	// so here only the limits are checked

	boxMode := learn.Kanji2Native

	testCases := []struct {
		name              string
		mode              string
		level, checkLevel int
		want              []learn.Card
	}{
		{
			name:       "stay at min level",
			mode:       boxMode,
			level:      learn.MinLevel - 1,
			checkLevel: learn.MinLevel,
			want:       cards1,
		},
		{
			name:       "set to max level",
			mode:       boxMode,
			level:      learn.MaxLevel + 1,
			checkLevel: learn.MaxLevel,
			want:       cards1[1:2],
		},
		{
			name:       "unknown mode, cards stay at min",
			mode:       "unknown",
			level:      learn.MaxLevel + 1,
			checkLevel: learn.MinLevel,
			want:       cards1,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			box := learn.NewKanjiBox(learn.BoxID{})
			box.AddCards(boxMode, learn.MinLevel, cards1...)

			box.AddCards(c.mode, c.level, cards1[1])
			got := box.Cards(boxMode, c.checkLevel)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
