package learn

// LessonID provides identification of a lesson in the books library.
type LessonID struct {
	Title       string // the title of the lesson
	BookTitle   string // the title of the book
	SeriesTitle string // the title of the book series
	Volume      int    // the volume number in the series
}

// BoxID provides identification of a box in the learn library.
// It consists of a name for the box and a lesson specification.
// Usually the BoxID is the name of the lesson.
// The lesson specification might be empty for mixed or indiviudally
// created boxes.
type BoxID struct {
	Title string
	LessonID
}

// CardID provides unique Identifier for a card.
// It contains the info about the book and the lesson where it is stored
// in the books.library.
type CardID struct {
	ContentID int // id from the content card
	LessonID
}
