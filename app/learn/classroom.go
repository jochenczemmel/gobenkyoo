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

// Classroom provides handling a set of learning boxes.
type Classroom struct {
	wordBoxes  map[BoxID]Box
	kanjiBoxes map[BoxID]Box
}

// NewClassroom creates a new learn shelf.
func NewClassroom() Classroom {
	return Classroom{
		wordBoxes:  make(map[BoxID]Box),
		kanjiBoxes: make(map[BoxID]Box),
	}
}

// NewWordBox adds a list of word cards to a newly created box.
func (c *Classroom) NewWordBox(boxid BoxID, cards ...words.Card) {
	box := NewBox(boxid)
	for _, mode := range GetWordModes() {
		box.Set(mode, box.makeWordCards(mode, cards...)...)
	}
	c.wordBoxes[boxid] = box
}

// StartWordExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (c *Classroom) StartWordExam(opt Options, boxids ...BoxID) Exam {
	boxes := []Box{}
	for _, boxName := range boxids {
		boxes = append(boxes, c.wordBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}

// NewKanjiBox adds a list of kanji cards to a newly created box.
func (c *Classroom) NewKanjiBox(boxid BoxID, cards ...kanjis.Card) {
	box := NewBox(boxid)
	for _, mode := range GetKanjiModes() {
		box.Set(mode, box.makeKanjiCards(mode, cards...)...)
	}
	c.kanjiBoxes[boxid] = box
}

// StartKanjiExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (c *Classroom) StartKanjiExam(opt Options, boxids ...BoxID) Exam {
	boxes := []Box{}
	for _, boxName := range boxids {
		boxes = append(boxes, c.kanjiBoxes[boxName])
	}
	return NewExam(opt, boxes...)
}
