package kanjis

import "github.com/jochenczemmel/gobenkyoo/content"

// Collection represents a list of kanji cards with a title and book title.
// It uses a generic content.Collection.
type Collection struct {
	content.Collection[Card]
}

// NewCollection returns a kanji card Collection with the given titles and
// optionally filled with the given cards.
func NewCollection(title, booktitle string, cards ...Card) Collection {
	result := Collection{
		Collection: content.NewCollection[Card](title, booktitle, cards...),
	}

	return result
}

// CardsWithIDs returns a list of kanji cards from the kanjilist.
// The order in kanjilist is preserved.
// If a kanji is not in the collection, it is skipped.
// Duplicate kanjis are only returned on the first occurrence.
func (c Collection) CardsWithIDs(kanjilist []string) []Card {
	result := []Card{}
	found := map[string]bool{}

	for _, id := range kanjilist {
		card := c.Get(id)
		if card.ID() == "" {
			continue
		}
		if found[card.ID()] {
			continue
		}
		found[card.ID()] = true
		result = append(result, card)
	}

	return result
}

// CardsWithRadicals returns all cards with Kanjis that
// have the requested radicals.
func (c Collection) CardsWithRadicals(radicallist string) []Card {
	result := []Card{}

	if radicallist == "" {
		return result
	}

LOOP:
	for _, kanji := range c.Content() {
		for _, rad := range radicallist {
			if !kanji.HasRadical(rad) {
				continue LOOP
			}
		}
		result = append(result, kanji)
	}

	return result
}
