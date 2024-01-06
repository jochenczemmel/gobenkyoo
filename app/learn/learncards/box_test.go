package learncards_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
)

func TestBoxNew(t *testing.T) {

	title := "lesson 1"
	bookTitle := "minna 1"
	box := learncards.NewBox(title, bookTitle)

	t.Run("titles", func(t *testing.T) {
		assertEquals(t, box.Title, title)
		assertEquals(t, box.BookTitle, bookTitle)
	})
}

func TestBoxModes(t *testing.T) {

	box := learncards.NewBox("", "")

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

	box := learncards.NewBox("", "")
	boxMode := "mode 1"
	box.Set(boxMode, cards1...)

	testCases := []struct {
		name        string
		mode        string
		level, want int
		wantCards   []learncards.Card
	}{{
		name:      "all cards",
		mode:      boxMode,
		level:     learncards.AllLevel,
		want:      len(cards1),
		wantCards: cards1,
	}, {
		name:      "first level",
		mode:      boxMode,
		level:     learncards.MinLevel,
		want:      len(cards1),
		wantCards: cards1,
	}, {
		name:      "next level",
		mode:      boxMode,
		level:     learncards.MinLevel + 1,
		want:      0,
		wantCards: []learncards.Card{},
	}, {
		name:      "unknown mode",
		mode:      "unknown",
		level:     learncards.AllLevel,
		want:      0,
		wantCards: []learncards.Card{},
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
	box := learncards.NewBox("", "")
	boxMode := "mode 1"

	testCases := []struct {
		name, mode        string
		level, checkLevel int
		want              []learncards.Card
	}{
		{
			name:       "stay at min level",
			mode:       boxMode,
			level:      learncards.MinLevel - 1,
			checkLevel: learncards.MinLevel,
			want:       cards1,
		},
		{
			name:       "set to max level",
			mode:       boxMode,
			level:      learncards.MaxLevel + 1,
			checkLevel: learncards.MaxLevel,
			want:       cards1[1:2],
		},
		{
			name:       "unknown mode, cards stay at min",
			mode:       "unknown",
			level:      learncards.MaxLevel + 1,
			checkLevel: learncards.MinLevel,
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

func assertEquals[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}
