package kanjis_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

var kanjiComparer = cmp.Comparer(func(x, y kanjis.Card) bool {
	return x.Kanji == y.Kanji
})

var (
	card1 = kanjis.NewBuilder('山').Build()
	card2 = kanjis.NewBuilder('島').Build()
	card3 = kanjis.NewBuilder('元').Build()
)

func TestCollectionCardsWithRadicals(t *testing.T) {
	testCases := []struct {
		name        string
		collection  kanjis.Collection
		radicalList string
		want        []kanjis.Card
	}{{
		name:        "uninitialized collection",
		radicalList: "山",
		want:        []kanjis.Card{},
	}, {
		name:        "empty collection",
		collection:  kanjis.NewCollection("", ""),
		radicalList: "山",
		want:        []kanjis.Card{},
	}, {
		name:        "get two",
		collection:  kanjis.NewCollection("", "", card1, card2, card3),
		radicalList: "山",
		want:        []kanjis.Card{card1, card2},
	}, {
		name:        "get one",
		collection:  kanjis.NewCollection("", "", card1, card2, card3),
		radicalList: "山鳥",
		want:        []kanjis.Card{card2},
	}, {
		name:        "duplicate in search",
		collection:  kanjis.NewCollection("", "", card1, card2, card3),
		radicalList: "山鳥山",
		want:        []kanjis.Card{card2},
	}, {
		name:        "no match",
		collection:  kanjis.NewCollection("", "", card1, card2, card3),
		radicalList: "右",
		want:        []kanjis.Card{},
	}, {
		name:        "empty search list",
		collection:  kanjis.NewCollection("", "", card1, card2, card3),
		radicalList: "",
		want:        []kanjis.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.collection.CardsWithRadicals(c.radicalList)
			if diff := cmp.Diff(got, c.want, kanjiComparer); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

func TestCollectionCardsWithIDs(t *testing.T) {

	testCases := []struct {
		name       string
		collection kanjis.Collection
		kanjiList  []string
		want       []kanjis.Card
	}{{
		name:      "uninitialized collection",
		kanjiList: []string{"山"},
		want:      []kanjis.Card{},
	}, {
		name:       "empty collection",
		collection: kanjis.NewCollection("", ""),
		kanjiList:  []string{"山"},
		want:       []kanjis.Card{},
	}, {
		name:       "get one",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"山"},
		want:       []kanjis.Card{card1},
	}, {
		name:       "get three",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"山", "島", "元"},
		want:       []kanjis.Card{card1, card2, card3},
	}, {
		name:       "three in different order",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"島", "元", "山"},
		want:       []kanjis.Card{card2, card3, card1},
	}, {
		name:       "duplicates in search",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"山", "元", "山", "元", "山"},
		want:       []kanjis.Card{card1, card3},
	}, {
		name:       "one not found",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"山", "方", "元"},
		want:       []kanjis.Card{card1, card3},
	}, {
		name:       "none found",
		collection: kanjis.NewCollection("", "", card1, card2, card3),
		kanjiList:  []string{"庭", "方"},
		want:       []kanjis.Card{},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.collection.CardsWithIDs(c.kanjiList)
			if diff := cmp.Diff(got, c.want, kanjiComparer); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
