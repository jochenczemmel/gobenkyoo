package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func TestBoxModes(t *testing.T) {

	box := learn.NewBox(learn.BoxID{}, "")

	testCases := []struct {
		name string
		mode string
		want []string
	}{{
		name: "first mode",
		mode: "mode 1",
		want: []string{"mode 1"},
	}, {
		name: "second mode",
		mode: "mode 2",
		want: []string{"mode 1", "mode 2"},
	}, {
		name: "duplicate mode",
		mode: "mode 2",
		want: []string{"mode 1", "mode 2"},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			box.Set(c.mode, cards1...)
			got := box.Modes()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})
	}
}

func TestBoxCards(t *testing.T) {

	box := learn.NewBox(learn.BoxID{}, "")
	boxMode := "mode 1"
	box.Set(boxMode, cards1...)

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
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})
	}
}

func TestBoxSetLevelLimits(t *testing.T) {
	// SetCardLevel() is also tested when testing exam.Advance()
	// so here only the limits are checked
	box := learn.NewBox(learn.BoxID{}, "")
	boxMode := "mode 1"

	testCases := []struct {
		name, mode        string
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
			box.Set(boxMode, cards1...)

			box.SetCardLevel(c.mode, cards1[1], c.level)
			got := box.Cards(boxMode, c.checkLevel)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})

	}
}

func assertEquals[T comparable](tb testing.TB, got, want T) {
	tb.Helper()
	if got != want {
		tb.Errorf("ERROR: got %v, want %v", got, want)
	}
}
