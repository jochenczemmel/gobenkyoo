package learn

import "math/rand"

// Exam provides functions to execute an exam for a list of cards.
type Exam struct {
	// Number of cards in the exam.
	NCards      int
	boxes       []*Box
	currentCard *Card
	cards       []*Card // might grow due to appending failed cards
	cardsIndex  int
	opt         ExamOptions
}

// ExamOptions define the behaviour of the Exam.
type ExamOptions struct {
	LearnMode string // required, must be valid
	Level     int    // required, must be valid
	MoveCards bool   // advance/reset cards in box?
	Repeat    bool   // repeat failed cards until known?
	NoShuffle bool   // do not shuffle the cards
}

// NewExam starts a new exam.
func NewExam(opt ExamOptions, boxes ...*Box) *Exam {
	result := &Exam{
		opt:         opt,
		boxes:       boxes,
		currentCard: &Card{},
	}
	for _, box := range boxes {
		result.cards = append(result.cards, box.Cards(opt.LearnMode, opt.Level)...)
	}
	result.NCards = len(result.cards)
	if !opt.NoShuffle {
		result.shuffle()
	}
	return result
}

// shuffle shuffles the cards.
func (e *Exam) shuffle() {
	rand.Shuffle(len(e.cards), func(i, j int) {
		e.cards[i], e.cards[j] = e.cards[j], e.cards[i]
	})
}

// NextCard returns the next card of the exam and true.
// If there are no more cards, it returns the current card and false.
func (e *Exam) NextCard() (*Card, bool) {
	if e.cardsIndex >= len(e.cards) ||
		e.cardsIndex < 0 {
		return e.currentCard, false
	}

	e.currentCard = e.cards[e.cardsIndex]
	e.cardsIndex++

	return e.currentCard, true
}

// Pass means the current card has been answered correctly.
// If the option MoveCards is true, it puts the current card
// on the next level.
// If it is already on the maximum level, it stays there.
func (e *Exam) Pass() {
	if e.opt.MoveCards {
		for _, box := range e.boxes {
			box.Advance(e.opt.LearnMode, e.currentCard)
		}
	}
}

// Fail means the current card has not been answered correctly.
// If the option MoveCards is true, it puts the current card on
// the minimum level.
// If the option Repeat is true, it appends the card to the list
// so it is shown again.
func (e *Exam) Fail() {
	if e.opt.MoveCards {
		for _, box := range e.boxes {
			box.Reset(e.opt.LearnMode, e.currentCard)
		}
	}
	if e.opt.Repeat {
		e.cards = append(e.cards, e.currentCard)
	}
}

// PreviousCard returns the previous Card and true.
// If there are no more cards, it returns the current card and false.
func (e *Exam) PreviousCard() (*Card, bool) {
	if e.cardsIndex > len(e.cards) ||
		e.cardsIndex <= 0 {
		return e.currentCard, false
	}

	e.cardsIndex--
	e.currentCard = e.cards[e.cardsIndex]

	return e.currentCard, true
}
