// Package learn provides learning vocabulary and kanjis with
// the 'Leitner System' method of learning.
// Word and kanji cards can be used to fill boxes that can be
// used for executing exams. Different learning modes are
// available for words and kanjis, the progress is tracked
// separately. The mode and box level can be selected when
// an exam is started.
package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// BoxInfo consists of a name for the box and a book specification.
// Usually the BoxInfo is the name of the lesson.
type BoxInfo struct {
	Title     string
	BookTitle books.TitleInfo
}

// Library provides handling a set of learning boxes.
type Library struct {
	wordBoxes  map[BoxInfo]Box
	kanjiBoxes map[BoxInfo]Box
}

// NewLibrary creates a new learn shelf.
func NewLibrary() Library {
	return Library{
		wordBoxes:  make(map[BoxInfo]Box),
		kanjiBoxes: make(map[BoxInfo]Box),
	}
}

// AddWordBox adds a list of word cards to a box.
func (l *Library) AddWordBox(name BoxInfo, cards ...words.Card) {
	box := NewBox(name)
	for _, mode := range GetWordModes() {
		box.Set(mode, makeWordCards(mode, cards...)...)
	}
	l.wordBoxes[name] = box
}

// StartWordExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (l *Library) StartWordExam(opt Options, boxnames ...BoxInfo) Exam {
	boxes := []Box{}
	for _, boxName := range boxnames {
		boxes = append(boxes, l.wordBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}

// AddKanjiBox adds a list of kanji cards to a box.
func (l *Library) AddKanjiBox(name BoxInfo, cards ...kanjis.Card) {
	box := NewBox(name)
	for _, mode := range GetKanjiModes() {
		box.Set(mode, makeKanjiCards(mode, cards...)...)
	}
	l.kanjiBoxes[name] = box
}

// StartKanjiExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (l *Library) StartKanjiExam(opt Options, boxnames ...BoxInfo) Exam {
	boxes := []Box{}
	for _, boxName := range boxnames {
		boxes = append(boxes, l.kanjiBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}
