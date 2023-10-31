package words

import "github.com/jochenczemmel/gobenkyoo/content"

type Collection struct {
	content.Collection[Card]
	BookTitle string
}

func NewCollection(title, booktitle string, cards ...Card) Collection {
	result := Collection{
		Collection: content.NewCollection[Card](title, cards...),
		BookTitle:  booktitle,
	}
	return result
}
