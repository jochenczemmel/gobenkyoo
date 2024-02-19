package learn

import "github.com/jochenczemmel/gobenkyoo/content/books"

// BoxID provides identification of a box in the learn library.
// It consists of a name for the box and a lesson specification.
// Usually the BoxID is the name of the lesson.
// The lesson specification might be empty for mixed or individually
// created boxes.
type BoxID struct {
	Name           string // Name of the box
	books.LessonID        // reference to the lesson
}

// CardID provides unique Identifier for a card.
// It contains the info about the book and the lesson where it is stored
// in the books.library.
type CardID struct {
	ContentID      string // id from the content card
	books.LessonID        // reference to the lesson
}
