// Package books provides information about Books, Lessons and Libraqies.
// It alos provides access to the content kanji and word cards.
package books

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
// The lesson order is preserved.
type Book struct {
	TitleInfo
	Lessons []Lesson
}

// New returns a new book with the specified infos.
func New(title, seriestitle string, volume int, lessons ...Lesson) Book {
	return Book{
		TitleInfo: TitleInfo{
			Title:       title,
			SeriesTitle: seriestitle,
			Volume:      volume,
		},
		Lessons: lessons,
	}
}
