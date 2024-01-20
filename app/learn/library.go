// Package learn provides learning vocabulary and kanjis with
// the 'Leitner System' method of learning.
// Word and kanji cards can be used to fill boxes that can be
// used for executing exams. Different learning modes are
// available for words and kanjis, the progress is tracked
// separately. The mode and box level can be selected when
// an exam is started.
package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Library provides handling a set of learning boxes.
type Library struct {
	wordBoxes  map[BoxID]Box
	kanjiBoxes map[BoxID]Box
}

// NewLibrary creates a new learn shelf.
func NewLibrary() Library {
	return Library{
		wordBoxes:  make(map[BoxID]Box),
		kanjiBoxes: make(map[BoxID]Box),
	}
}

// NewWordBox adds a list of word cards to a newly created box.
func (l *Library) NewWordBox(boxid BoxID, cards ...words.Card) {
	box := NewBox(boxid)
	for _, mode := range GetWordModes() {
		box.Set(mode, box.makeWordCards(mode, cards...)...)
	}
	l.wordBoxes[boxid] = box
}

// StartWordExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (l *Library) StartWordExam(opt Options, boxids ...BoxID) Exam {
	boxes := []Box{}
	for _, boxName := range boxids {
		boxes = append(boxes, l.wordBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}

// NewKanjiBox adds a list of kanji cards to a newly created box.
func (l *Library) NewKanjiBox(boxid BoxID, cards ...kanjis.Card) {
	box := NewBox(boxid)
	for _, mode := range GetKanjiModes() {
		box.Set(mode, box.makeKanjiCards(mode, cards...)...)
	}
	l.kanjiBoxes[boxid] = box
}

// StartKanjiExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (l *Library) StartKanjiExam(opt Options, boxids ...BoxID) Exam {
	boxes := []Box{}
	for _, boxName := range boxids {
		boxes = append(boxes, l.kanjiBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}
