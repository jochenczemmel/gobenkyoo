package learncards

import "math/rand"

// ExamOptions define the behaviour of the Exam.
type ExamOptions struct {
	LearnMode string // required, must be valid
	Level     int    // required, must be valid
	NoShuffle bool   // do not shuffle the cards
	KeepLevel bool   // do not advance/reset cards on Pass() or Fail()
	// Repeat    bool   // repeat failed cards until known?
}

// Exam provides a single learn test execution.
type Exam struct {
	ExamOptions
	containers   []container
	cards        []Card
	currentCard  Card
	currentIndex int
}

// NewExam creates a new exam using the given mode, levels
// and uses the cards from the provided boxes.
// func NewExam(mode string, level int, boxes ...Box) Exam {
func NewExam(opt ExamOptions, boxes ...Box) Exam {
	cards := []Card{}
	containers := []container{}
	for _, box := range boxes {
		cards = append(cards, box.Cards(opt.LearnMode, opt.Level)...)
		containers = append(containers, box.containers[opt.LearnMode])
	}
	result := Exam{
		ExamOptions: opt,
		containers:  containers,
		cards:       cards,
	}
	if !opt.NoShuffle {
		result.shuffle()
	}

	return result
}

// shuffle shuffles the remaining cards.
func (e *Exam) shuffle() {
	rand.Shuffle(len(e.cards), func(i, j int) {
		e.cards[i], e.cards[j] = e.cards[j], e.cards[i]
	})
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

// Advance shifts the card on the next level.
func (e Exam) Advance(card Card) {
	for _, c := range e.containers {
		c.advance(card)
	}
}

// Pass shifts the current card on the next level
// unless option KeepLevel is true.
func (e Exam) Pass() {
	if e.KeepLevel {
		return
	}
	e.Advance(e.currentCard)
}

// Reset puts the card on the minimum (starting) level.
func (e Exam) Reset(card Card) {
	for _, c := range e.containers {
		c.setLevel(card, MinLevel)
	}
}

// Fail puts the current card on the minimum (starting) level
// unless option KeepLevel is true.
func (e Exam) Fail() {
	if e.KeepLevel {
		return
	}
	e.Reset(e.currentCard)
}

// NextCard returns the next card of the exam and true.
// If there are no more cards, it returns the current card and false.
func (e *Exam) NextCard() (Card, bool) {
	if e.currentIndex >= len(e.cards) {
		return emptyCard, false
	}

	e.currentCard = e.cards[e.currentIndex]
	e.currentIndex++

	return e.currentCard, true
}
