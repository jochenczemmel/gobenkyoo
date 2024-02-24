package books

import (
	"slices"
	"sort"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
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

// AddBooks adds books to the library.
// The order is preserved.
// Duplicate additions are ignored.
func (l *Library) AddBooks(books ...Book) {
	for _, book := range books {
		_, ok := l.booksByID[book.ID]
		if !ok {
			l.booksByID[book.ID] = book
			l.Books = append(l.Books, book)
		}
	}
}

// SortedBooks returns a list of books sorted according to
// series title, volume and book title.
func (l Library) SortedBooks() []Book {
	result := slices.Clone(l.Books)
	sort.Sort(bySeriesVolumeTitle(result))

	return result
}

// getWordCard returns a the word card from the specified book and lesson
// with the specified id. If it is not found, an empty card is returned.
func (l Library) WordCard(lessonid LessonID, cardid string) words.Card {
	return l.booksByID[lessonid.ID].getWordCard(lessonid.Name, cardid)
}

// getKanjiCard returns a the kanji card from the specified book and lesson
// with the specified id. If it is not found, an empty card is returned.
func (l Library) KanjiCard(lessonid LessonID, cardid string) kanjis.Card {
	return l.booksByID[lessonid.ID].getKanjiCard(lessonid.Name, cardid)
}
