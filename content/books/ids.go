package books

// LessonID provides identification of a lesson in the library.
type LessonID struct {
	Title string // the title of the lesson
	ID           // the book id
}

func NewLessonID(title, booktitle, seriestitle string, volume int) LessonID {
	return LessonID{
		Title: title,
		ID: ID{
			Title:       booktitle,
			SeriesTitle: seriestitle,
			Volume:      volume,
		},
	}
}

// ID provides identification of a book in the library.
type ID struct {
	Title       string // the title of the book
	SeriesTitle string // the title of the book series
	Volume      int    // the volume number in the series
}

func NewID(title, seriestitle string, volume int) ID {
	return ID{
		Title:       title,
		SeriesTitle: seriestitle,
		Volume:      volume,
	}
}
