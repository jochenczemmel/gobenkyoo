package words_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestCardString(t *testing.T) {

	testCases := []struct {
		name string
		card words.Card
		want string
	}{{
		name: "empty",
		want: "",
	}, {
		name: "id",
		card: words.Card{Identifier: "card1"},
		want: `card1: ""`,
	}, {
		name: "with value",
		card: words.Card{Identifier: "card1", Nihongo: "世界"},
		want: `card1: "世界"`,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.card.String()
			if got != c.want {
				t.Errorf("ERROR: got %q, want %q", got, c.want)
			}
		})
	}
}
