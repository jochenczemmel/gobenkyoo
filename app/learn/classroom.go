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
	uniqIDs    map[BoxID]bool
	boxIDs     []BoxID
}

// NewClassroom creates a new learn shelf.
func NewClassroom() Classroom {
	return Classroom{
		wordBoxes:  make(map[BoxID]Box),
		kanjiBoxes: make(map[BoxID]Box),
		uniqIDs:    make(map[BoxID]bool),
	}
}

// Boxes returns the boxes in the order they
// have been added to the classroom.
func (c Classroom) Boxes() []Box {
	result := make([]Box, 0, len(c.boxIDs))
	for _, id := range c.boxIDs {
		if box, ok := c.wordBoxes[id]; ok {
			result = append(result, box)
		}
		if box, ok := c.kanjiBoxes[id]; ok {
			result = append(result, box)
		}
	}
	return result
}

// NewWordBox adds a list of word cards to a newly created box.
func (c *Classroom) NewWordBox(boxid BoxID, cards ...words.Card) {
	_, ok := c.wordBoxes[boxid]
	if ok {
		// append?
		return
	}
	box := NewBox(boxid, WordType)
	for _, mode := range GetWordModes() {
		box.Set(mode, box.makeWordCards(mode, cards...)...)
	}
	c.wordBoxes[boxid] = box

	if !c.uniqIDs[boxid] {
		c.boxIDs = append(c.boxIDs, boxid)
		c.uniqIDs[boxid] = true
	}
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
	_, ok := c.kanjiBoxes[boxid]
	if ok {
		// append?
		return
	}
	box := NewBox(boxid, KanjiType)
	for _, mode := range GetKanjiModes() {
		box.Set(mode, box.makeKanjiCards(mode, cards...)...)
	}
	c.kanjiBoxes[boxid] = box
	if !c.uniqIDs[boxid] {
		c.boxIDs = append(c.boxIDs, boxid)
		c.uniqIDs[boxid] = true
	}
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
