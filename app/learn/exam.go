package learn

type Exam struct {
}

func (e Exam) NextCard() (Card, bool) {
	return Card{}, true
}

func (e Exam) Advance() {}

func (e Exam) Reset() {}

func (e Exam) Save() error {
	return nil
}
