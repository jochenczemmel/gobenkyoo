package books

import "sort"

// Library provides access to a list of books.
type Library struct {
	Name      string
	booksByID map[ID]Book
}

// NewLibrary returns a library object with the given title.
func NewLibrary(name string) Library {
	return Library{
		Name:      name,
		booksByID: map[ID]Book{},
	}
}

// SetBooks adds or replaces books to the library.
// The order is preserved.
func (l *Library) SetBooks(books ...Book) {
	for _, book := range books {
		l.booksByID[book.ID] = book
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

// SortedBookIDs returns a list of book ids sorted according to
// series title, volume and book title.
func (l Library) SortedBookIDs() []ID {
	result := make([]ID, 0, len(l.booksByID))
	for id := range l.booksByID {
		result = append(result, id)
	}
	sort.Sort(bySeriesVolumeTitle(result))

	return result
}
