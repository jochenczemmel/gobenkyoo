package learncards

import "math/rand"

// Exam provides a single learn test execution.
type Exam struct {
	mode  string
	level int
	boxes []Box
	cards []Card
}

// NewExam creates a new exam using the given mode, levels
// and uses the cards from the provided boxes.
func NewExam(mode string, level int, boxes ...Box) Exam {
	cards := []Card{}
	for _, box := range boxes {
		cards = append(cards, box.Cards(mode, level)...)
	}
	return Exam{
		mode:  mode,
		level: level,
		boxes: boxes,
		cards: cards,
	}
}

// NCards returns the number of cards in the exam.
func (e Exam) NCards() int {
	return len(e.cards)
}

// Cards returns the cards in the exam.
// If the cards have not been shuffled, the cards are in the
// same order as returned from the boxes.
func (e Exam) Cards() []Card {
	result := make([]Card, len(e.cards))
	copy(result, e.cards)
	return result
}

// Shuffle shuffles the remaining cards.
func (e *Exam) Shuffle() {
	rand.Shuffle(len(e.cards), func(i, j int) {
		e.cards[i], e.cards[j] = e.cards[j], e.cards[i]
	})
}
