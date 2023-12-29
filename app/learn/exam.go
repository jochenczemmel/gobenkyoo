package learn

type Exam struct {
	boxes       []*Box
	currentCard *Card
	cards       []*Card
	cardsIndex  int
	ExamOptions
}

type ExamOptions struct {
	LearnMode string
	Level     int
	MoveCards bool
	Repeat    bool
	Shuffle   bool
}

// NewExam starts a new exam.
func NewExam(opt ExamOptions, boxes ...*Box) *Exam {
	result := &Exam{
		ExamOptions: opt,
		boxes:       boxes,
		currentCard: &Card{},
	}
	for _, box := range boxes {
		result.cards = append(result.cards, box.Cards(opt.LearnMode, opt.Level)...)
	}
	return result
}

func (e *Exam) Shuffle() {}

// NextCard returns the next card of the exam and true.
// If there are no more cards, it returns an empty card and false.
func (e *Exam) NextCard() (*Card, bool) {
	if e.cardsIndex >= len(e.cards) ||
		e.cardsIndex < 0 {
		e.currentCard = &Card{}
		return e.currentCard, false
	}

	e.currentCard = e.cards[e.cardsIndex]
	e.cardsIndex++

	return e.currentCard, true
}

// Advance puts the current card on the next level.
// If it is already on the maximum level, it stays there.
func (e Exam) Advance() {
	for _, box := range e.boxes {
		box.Advance(e.LearnMode, e.currentCard)
	}
}

// Reset puts the current card on the minimum level.
func (e Exam) Reset() {
	for _, box := range e.boxes {
		box.Reset(e.LearnMode, e.currentCard)
	}
}

// TODO: is this a method for exam?
// func (e Exam) Save() error {
// 	return nil
// }

// func (e *Exam) PreviousCard() (*Card, bool) {
// 	if e.cardsIndex > len(e.cards) ||
// 		e.cardsIndex <= 0 {
//       e.currentCard = &Card{}
//          return e.currentCard, false
// 	}
//
// 	e.cardsIndex--
//  	e.currentCard = e.cards[e.cardsIndex]
//
// 	return e.currentCard, true
// }
