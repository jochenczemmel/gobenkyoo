package learn

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// BoxName consists of a name for the box and a book specification.
// Usually the BoxName is the name of the lesson.
type BoxName struct {
	BoxTitle  string
	BookTitle books.TitleInfo
}

// Shelf provides handling a set of learning boxes.
type Shelf struct {
	wordBoxes  map[BoxName]learncards.Box
	kanjiBoxes map[BoxName]learncards.Box
}

// NewShelf creates a new learn shelf.
func NewShelf() Shelf {
	return Shelf{
		wordBoxes:  make(map[BoxName]learncards.Box),
		kanjiBoxes: make(map[BoxName]learncards.Box),
	}
}

// AddWordBox adds a list of word cards to a learncards box.
func (s *Shelf) AddWordBox(name BoxName, cards ...words.Card) {
	box := learncards.NewBox(name.BoxTitle, name.BookTitle.Title)
	for _, mode := range GetWordModes() {
		box.Set(mode, makeWordCards(mode, cards...)...)
	}
	s.wordBoxes[name] = box
}

// func (s *Shelf) StartWordExam(title, booktitle string, cards ...words.Card) {
