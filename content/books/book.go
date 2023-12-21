// Package books provides information about Books, Lessons and Content.
package books

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book struct {
	Title       string    // the title of the book
	SeriesTitle string    // the title of the book collection
	Volume      int       // the volume number in the collection
	Lessons     []*Lesson // the ordered lessons
}
