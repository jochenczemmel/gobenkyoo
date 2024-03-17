package books

import (
	"slices"
	"sort"
)

// Library provides access to a list of books.
type Library struct {
	Name      string
	Books     []Book
	booksByID map[ID]Book
}

// NewLibrary returns a library object with the given title.
func NewLibrary(name string) Library {
	return Library{
		Name:      name,
		Books:     []Book{},
		booksByID: map[ID]Book{},
	}
}

// SetBooks adds or replaces books to the library.
// The order is preserved.
func (l *Library) SetBooks(books ...Book) {

LOOP:
	for _, book := range books {
		_, ok := l.booksByID[book.ID]
		if !ok {
			l.booksByID[book.ID] = book
			l.Books = append(l.Books, book)
			continue LOOP
		}
		for i, b := range l.Books {
			if b.ID == book.ID {
				l.Books[i] = book
				continue LOOP
			}
		}
	}
}

// Book returns the book with the given id.
// If it is not found, a new book with the id is returned.
func (l Library) Book(id ID) Book {
	book, ok := l.booksByID[id]
	if !ok {
		return New(id)
	}
	return book
}

// SortedBooks returns a list of books sorted according to
// series title, volume and book title.
func (l Library) SortedBooks() []Book {
	result := slices.Clone(l.Books)
	sort.Sort(bySeriesVolumeTitle(result))

	return result
}

// SortedBookIDs returns a list of book ids sorted according to
// series title, volume and book title.
func (l Library) SortedBookIDs() []ID {
	result := make([]ID, 0, len(l.Books))
	for _, book := range l.SortedBooks() {
		result = append(result, book.ID)
	}

	return result
}
