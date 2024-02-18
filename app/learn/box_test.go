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

/*
func TestBoxUninitialized(t *testing.T) {

	t.Run("Cards", func(t *testing.T) {
		var box learn.Box
		got := box.Cards("", 1)
		want := []learn.Card{}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: -got +want\n%s", diff)
		}
	})

	t.Run("NCards", func(t *testing.T) {
		var box learn.Box
		got := box.NCards("", 1)
		want := 0
		if got != want {
			t.Errorf("ERROR: got %v, want %v", got, want)
		}
	})

		t.Run("SetCardLevel", func(t *testing.T) {
			var box learn.Box
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("ERROR: got error: %v", err)
				}
			}()
			box.SetCardLevel("", 1, learn.Card{})
		})
}
*/

/*
func TestBoxSetLevelLimits(t *testing.T) {
	// SetCardLevel() is also tested when testing exam.Advance()
	// so here only the limits are checked
	box := learn.NewKanjiBox(learn.BoxID{})
	boxMode := "mode 1"

	testCases := []struct {
		name              string
		box               learn.Box
		mode              string
		level, checkLevel int
		want              []learn.Card
	}{
		{
			name:       "stay at min level",
			box:        box,
			mode:       boxMode,
			level:      learn.MinLevel - 1,
			checkLevel: learn.MinLevel,
			want:       cards1,
		},
		{
			name:       "set to max level",
			box:        box,
			mode:       boxMode,
			level:      learn.MaxLevel + 1,
			checkLevel: learn.MaxLevel,
			want:       cards1[1:2],
		},
		{
			name:       "unknown mode, cards stay at min",
			box:        box,
			mode:       "unknown",
			level:      learn.MaxLevel + 1,
			checkLevel: learn.MinLevel,
			want:       cards1,
		},
		{
			name:       "zero value box",
			box:        learn.Box{},
			mode:       boxMode,
			level:      learn.MaxLevel + 1,
			checkLevel: learn.MaxLevel,
			want:       cards1[1:2],
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			c.box.SetCards(boxMode, learn.MinLevel, cards1...)

			c.box.SetCards(c.mode, c.level, cards1[1])
			got := c.box.Cards(boxMode, c.checkLevel)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})

	}
}
*/
