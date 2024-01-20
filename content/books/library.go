package books

import (
	"slices"
	"sort"
)

// Library provides access to a list of books.
type Library struct {
	Title string
	Books []Book
}

// NewLibrary returns a library object with the given title.
func NewLibrary(title string) Library {
	return Library{
		Title: title,
		Books: []Book{},
	}
}

// AddBooks adds books to the library.
// The order is preserved.
func (l *Library) AddBooks(books ...Book) {
	l.Books = append(l.Books, books...)
}

// SortedBooks returns a list of books sorted according to
// series title, volume and book title.
func (l Library) SortedBooks() []Book {
	result := slices.Clone(l.Books)
	sort.Sort(bySeriesVolumeTitle(result))
	return result
}
