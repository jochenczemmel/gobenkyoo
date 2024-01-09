package learn

import "math/rand"

// Options define the behaviour of the Exam.
type Options struct {
	LearnMode string // required, must be one of the defined modes
	Level     int    // required, must be within the defined boundaries
	NoShuffle bool   // do not shuffle the cards
	KeepLevel bool   // do not advance/reset cards on Pass() or Fail()
	Repeat    bool   // repeat failed cards until known
}

// Exam provides a single learn test execution.
type Exam struct {
	opt          Options     // modify exam behaviour
	containers   []container // cards for the given learn mode
	initialCards []Card      // fixed list
	cards        []Card      // list may grow when Repeat==true and Fail() is called
	currentCard  Card        // used to iterate through the cards list
	currentIndex int         // used to iterate through the cards list
}

// NewExam creates a new exam using the given mode, levels
// and uses the cards from the provided boxes.
func NewExam(opt Options, boxes ...Box) Exam {
	cards := []Card{}
	containers := []container{}
	for _, box := range boxes {
		cards = append(cards, box.Cards(opt.LearnMode, opt.Level)...)
		containers = append(containers, box.containers[opt.LearnMode])
	}
	result := Exam{
		opt:          opt,
		containers:   containers,
		initialCards: cards,
	}
	if !opt.NoShuffle {
		result.shuffle()
	}
	result.cards = result.initialCards

	return result
}

// shuffle shuffles the cards.
func (e *Exam) shuffle() {
	rand.Shuffle(len(e.initialCards), func(i, j int) {
		e.initialCards[i], e.initialCards[j] = e.initialCards[j], e.initialCards[i]
	})
}

// NCards returns the number of distinct (initial) cards in the exam.
func (e Exam) NCards() int {
	return len(e.initialCards)
}

// Cards returns the distinct (initial) cards in the exam.
// If the cards have not been shuffled, the cards are in the
// same order as returned from the boxes.
func (e Exam) Cards() []Card {
	result := make([]Card, len(e.initialCards))
	copy(result, e.initialCards)
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
	if e.opt.KeepLevel {
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
// If Repeat is true, the current card is appended to the list
// of cards to present.
func (e *Exam) Fail() {
	if e.opt.Repeat {
		e.cards = append(e.cards, e.currentCard)
	}
	if e.opt.KeepLevel {
		return
	}
	e.Reset(e.currentCard)
}

// NextCard returns the next card of the exam and true.
// If Repeat is true, failed cards are presented after the initial
// card set has been presented. The order of the repeated cards is preserved.
//
// If there are no more cards, it returns the current card and false.
func (e *Exam) NextCard() (Card, bool) {
	if e.currentIndex >= len(e.cards) {
		return emptyCard, false
	}

	e.currentCard = e.cards[e.currentIndex]
	e.currentIndex++

	return e.currentCard, true
}
