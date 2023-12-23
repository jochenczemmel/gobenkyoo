// Package books provides information about Books, Lessons and Content.
package books

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book struct {
	Info
	Lessons []*Lesson // the ordered lessons
}

// New returns a new book with the specified infos.
func New(title, seriestitle string, volume int, lessons ...*Lesson) *Book {
	return &Book{
		Info: Info{
			Title:       title,
			SeriesTitle: seriestitle,
			Volume:      volume,
		},
		Lessons: lessons,
	}
}
