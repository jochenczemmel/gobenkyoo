package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
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
	wordBoxes  map[BoxName]Box
	kanjiBoxes map[BoxName]Box
}

// NewShelf creates a new learn shelf.
func NewShelf() Shelf {
	return Shelf{
		wordBoxes:  make(map[BoxName]Box),
		kanjiBoxes: make(map[BoxName]Box),
	}
}

// AddWordBox adds a list of word cards to a box.
func (s *Shelf) AddWordBox(name BoxName, cards ...words.Card) {
	box := NewBox()
	for _, mode := range GetWordModes() {
		box.Set(mode, makeWordCards(mode, cards...)...)
	}
	s.wordBoxes[name] = box
}

// StartWordExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (s *Shelf) StartWordExam(opt Options, boxnames ...BoxName) Exam {
	boxes := []Box{}
	for _, boxName := range boxnames {
		boxes = append(boxes, s.wordBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}

// AddKanjiBox adds a list of kanji cards to a box.
func (s *Shelf) AddKanjiBox(name BoxName, cards ...kanjis.Card) {
	box := NewBox()
	for _, mode := range GetKanjiModes() {
		box.Set(mode, makeKanjiCards(mode, cards...)...)
	}
	s.kanjiBoxes[name] = box
}

// StartKanjiExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (s *Shelf) StartKanjiExam(opt Options, boxnames ...BoxName) Exam {
	boxes := []Box{}
	for _, boxName := range boxnames {
		boxes = append(boxes, s.kanjiBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}
