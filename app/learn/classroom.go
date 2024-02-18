// Package learn provides learning vocabulary and kanjis with
// the 'Leitner System' method of learning.
// Word and kanji cards can be used to fill boxes that can be
// used for executing exams. Different learning modes are
// available for words and kanjis, the progress is tracked
// separately. The mode and box level can be selected when
// an exam is started.
package learn

// Classroom provides handling a set of learning boxes.
type Classroom struct {
	Name        string
	kanjiBoxIDs []BoxID
	wordBoxIDs  []BoxID
	kanjiBoxes  map[BoxID]Box
	wordBoxes   map[BoxID]Box
}

// NewClassroom creates a new learn shelf.
func NewClassroom(name string) Classroom {
	return Classroom{
		Name:        name,
		kanjiBoxIDs: []BoxID{},
		wordBoxIDs:  []BoxID{},
		kanjiBoxes:  make(map[BoxID]Box),
		wordBoxes:   make(map[BoxID]Box),
	}
}

// SetKanjiBoxes adds or replaces a list of kanji boxes.
func (c *Classroom) SetKanjiBoxes(boxes ...Box) {
	for _, box := range boxes {
		if box.Type != KanjiType {
			continue
		}
		if _, ok := c.kanjiBoxes[box.BoxID]; !ok {
			c.kanjiBoxIDs = append(c.kanjiBoxIDs, box.BoxID)
		}
		c.kanjiBoxes[box.BoxID] = box
	}
}

// SetWordBoxes adds or replaces a list of word boxes.
func (c *Classroom) SetWordBoxes(boxes ...Box) {
	for _, box := range boxes {
		if box.Type != WordType {
			continue
		}
		if _, ok := c.wordBoxes[box.BoxID]; !ok {
			c.wordBoxIDs = append(c.wordBoxIDs, box.BoxID)
		}
		c.wordBoxes[box.BoxID] = box
	}
}

// KanjiBox returns the kanji box with the given id.
func (c Classroom) KanjiBox(boxid BoxID) Box {
	box, ok := c.kanjiBoxes[boxid]
	if !ok {
		return NewKanjiBox(BoxID{})
	}
	return box
}

// WordBox returns the word box with the given id.
func (c Classroom) WordBox(boxid BoxID) Box {
	box, ok := c.wordBoxes[boxid]
	if !ok {
		return NewWordBox(BoxID{})
	}
	return box
}

// getBoxes returns a list of boxes that match the given ids.
func getBoxes(boxes map[BoxID]Box, boxids ...BoxID) []Box {
	result := []Box{}
	for _, boxName := range boxids {
		if box, ok := boxes[boxName]; ok {
			result = append(result, box)
		}
	}
	return result
}

// KanjiBoxes returns the kanji boxes in the order they
// have been added to the classroom.
func (c Classroom) KanjiBoxes() []Box {
	return getBoxes(c.kanjiBoxes, c.kanjiBoxIDs...)
}

// WordBoxes returns the word boxes in the order they
// have been added to the classroom.
func (c Classroom) WordBoxes() []Box {
	return getBoxes(c.wordBoxes, c.wordBoxIDs...)
}

// StartWordExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (c *Classroom) StartWordExam(opt Options, boxids ...BoxID) Exam {
	return NewExam(opt, getBoxes(c.wordBoxes, boxids...)...)
}

// StartKanjiExam starts an exam with the given options that uses
// the cards from the specified box(es).
func (c *Classroom) StartKanjiExam(opt Options, boxids ...BoxID) Exam {
	return NewExam(opt, getBoxes(c.kanjiBoxes, boxids...)...)
}
