package words_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestCollectionGet(t *testing.T) {
	var emptyCard words.Card

	card1 := words.Card{Identifier: "card1"}
	card2 := words.Card{Identifier: "card2"}
	card3 := words.Card{Identifier: "card3"}

	testCases := []struct {
		name       string
		collection words.Collection
		id         string
		wantCard   words.Card
		wantAll    []words.Card
	}{{
		name:     "uninitialized collection",
		id:       "card1",
		wantCard: emptyCard,
	}, {
		name:       "empty collection",
		collection: words.NewCollection("", ""),
		id:         "card1",
		wantCard:   emptyCard,
	}, {
		name:       "card found",
		collection: words.NewCollection("", "", card1, card2, card3),
		id:         "card1",
		wantCard:   card1,
		wantAll:    []words.Card{card1, card2, card3},
	}, {
		name: "duplicates",
		collection: words.NewCollection("", "",
			card1, card2, card3, card3, card2, card1),
		id:       "card1",
		wantCard: card1,
		wantAll:  []words.Card{card1, card2, card3},
	}, {
		name:       "card not found",
		collection: words.NewCollection("", "", card1, card2, card3),
		id:         "notfound",
		wantCard:   emptyCard,
		wantAll:    []words.Card{card1, card2, card3},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.collection.Get(c.id); got != c.wantCard {
				t.Errorf("ERROR: got %v, want %v", got, c.wantCard)
			}
			got := c.collection.Content()
			if diff := cmp.Diff(got, c.wantAll); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
		})
	}
}
