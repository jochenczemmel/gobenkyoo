package learn

type Shelf struct {
	// Boxes []cards.box
}

func NewShelf() *Shelf {
	return &Shelf{}
}

func (s *Shelf) StartWordExam(mode string, level int, titles ...string) *Exam {
	return &Exam{}
}
