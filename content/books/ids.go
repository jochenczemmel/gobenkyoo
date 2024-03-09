package books

import "strconv"

// ID provides identification of a book in the library.
type ID struct {
	Title       string // the title of the book
	SeriesTitle string // the title of the book series
	Volume      int    // the volume number in the series
}

// NewID returns a new book id. If the title is missing,
// the seriestitle is used as title.
func NewID(title, seriestitle string, volume int) ID {
	if title == "" {
		title = seriestitle
	}
	return ID{
		Title:       title,
		SeriesTitle: seriestitle,
		Volume:      volume,
	}
}

// String returns a printable version of the id.
func (i ID) String() string {
	result := i.Title
	if i.SeriesTitle != "" {
		result += " (" + i.SeriesTitle
		if i.Volume > 0 {
			result += " - " + strconv.Itoa(i.Volume)
		}
		result += ")"
	}
	return result
}

// LessonID provides identification of a lesson in the library.
type LessonID struct {
	Name string // the name of the lesson
	ID          // the book id
}

// func NewLessonID(name, booktitle, seriestitle string, volume int) LessonID {
// 	return LessonID{
// 		Name: name,
// 		ID: ID{
// 			Title:       booktitle,
// 			SeriesTitle: seriestitle,
// 			Volume:      volume,
// 		},
// 	}
// }
