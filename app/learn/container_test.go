package learn

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContainer(t *testing.T) {

	inputCards := []Card{
		{ID: "一"},
		{ID: "二"},
		{ID: "三"},
	}

	t.Run("get cards", func(t *testing.T) {
		container := newContainer(inputCards...)

		cand := []struct {
			name  string
			level int
			want  int
		}{
			{"MinLevel", MinLevel, 3},
			{"MaxLevel", MaxLevel, 0},
		}

		for _, c := range cand {
			t.Run(c.name, func(t *testing.T) {
				got := container.cards(c.level)
				if len(got) != c.want {
					t.Errorf("ERROR: got %v, want %v", len(got), c.want)
				}
			})
		}

		t.Run("card order", func(t *testing.T) {
			got := container.cards(MinLevel)
			if diff := cmp.Diff(got, inputCards); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
		})
	})

	t.Run("move card", func(t *testing.T) {
		container := newContainer(inputCards...)

		newLevel := MinLevel + 1
		got := container.cards(newLevel)
		want := 0
		if len(got) != want {
			t.Errorf("ERROR: new level not empty: got %v, want %v", len(got), want)
		}

		container.setLevel(inputCards[1], newLevel)

		got = container.cards(newLevel)
		want = 1
		if len(got) != want {
			t.Errorf("ERROR: card not in new level: got %v, want %v", len(got), want)
		}

		got = container.cards(MinLevel)
		want = 2
		if len(got) != want {
			t.Errorf("ERROR: invalid previous level: got %v, want %v", len(got), want)
		}

		got = container.cards(AllLevel)
		want = 3
		if len(got) != want {
			t.Errorf("ERROR: get all cards: got %v, want %v", len(got), want)
		}
	})

	t.Run("move too low", func(t *testing.T) {
		container := newContainer(inputCards...)

		newLevel := MinLevel - 1
		container.setLevel(inputCards[1], newLevel)

		want := 3
		got := container.cards(MinLevel)
		if len(got) != want {
			t.Errorf("ERROR: invalid previous level: got %v, want %v", len(got), want)
		}
	})

	t.Run("move too high", func(t *testing.T) {
		container := newContainer(inputCards...)
		newLevel := MaxLevel + 1

		container.setLevel(inputCards[1], newLevel)

		got := container.cards(MinLevel)
		want := 3
		if len(got) != want {
			t.Errorf("ERROR: invalid previous level: got %v, want %v", len(got), want)
		}
	})

	t.Run("move unknown card", func(t *testing.T) {
		container := newContainer(inputCards...)

		newLevel := MinLevel + 1
		newCard := Card{ID: "unknown"}
		container.setLevel(newCard, newLevel)

		got := container.cards(MinLevel)
		want := 3
		if len(got) != want {
			t.Errorf("ERROR: invalid previous level: got %v, want %v", len(got), want)
		}
	})
}
