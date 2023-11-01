package words

import "github.com/jochenczemmel/gobenkyoo/content"

// Collection represents a list of word cards with a title and book title.
// It uses a generic content.Collection.
type Collection struct {
	content.Collection[Card]
}

// NewCollection returns a word card Collection with the given titles and
// optionally filled with the given cards.
func NewCollection(title, booktitle string, cards ...Card) Collection {
	result := Collection{
		Collection: content.NewCollection[Card](title, booktitle, cards...),
	}
	return result
}
