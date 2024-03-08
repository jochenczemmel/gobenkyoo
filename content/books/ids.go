package books

import "strconv"

// LessonID provides identification of a lesson in the library.
type LessonID struct {
	Name string // the name of the lesson
	ID          // the book id
}

func NewLessonID(name, booktitle, seriestitle string, volume int) LessonID {
	return LessonID{
		Name: name,
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
	if title == "" {
		title = seriestitle
	}
	return ID{
		Title:       title,
		SeriesTitle: seriestitle,
		Volume:      volume,
	}
}

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
