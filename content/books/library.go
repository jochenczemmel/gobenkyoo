package books

import (
	"sort"
)

// Library represents a list of books.
type Library struct {
	bookList     []*Book
	booksByTitle map[string]*Book
}

// NewLibrary returns a new Library object that uese the given books.
func NewLibrary(booklist ...*Book) Library {
	lib := Library{
		bookList:     booklist,
		booksByTitle: map[string]*Book{},
	}
	for _, book := range booklist {
		lib.booksByTitle[book.Title] = book
	}
	lib.sortBooks()

	return lib
}

// Content returns the books of the libraries.
func (l Library) Content() []*Book {
	if len(l.bookList) == 0 {
		return []*Book{}
	}
	result := make([]*Book, len(l.bookList))
	copy(result, l.bookList)

	return result
}

// ByTitle returns the book with the given title.
func (l Library) ByTitle(title string) *Book {
	book, ok := l.booksByTitle[title]
	if !ok {
		return &Book{}
	}

	return book
}

// BySeriesTitle returns the books with the given series title.
func (l Library) BySeriesTitle(title string) []*Book {
	result := []*Book{}
	for _, book := range l.bookList {
		if book.SeriesTitle == title {
			result = append(result, book)
		}
	}

	return result
}

// sortBooks sorts the books of the same series according
// to the volume number and creates sorted lists of lessons.
func (l *Library) sortBooks() {
	sort.Slice(l.bookList, func(i, j int) bool {
		if l.bookList[i].SeriesTitle < l.bookList[j].SeriesTitle {
			return true
		}
		if l.bookList[i].SeriesTitle > l.bookList[j].SeriesTitle {
			return false
		}
		if l.bookList[i].Volume < l.bookList[j].Volume {
			return true
		}
		if l.bookList[i].Volume > l.bookList[j].Volume {
			return false
		}

		return l.bookList[i].Title < l.bookList[j].Title
	})
}

// BookTitles returns the titles of all books in sorted order.
func (l *Library) BookTitles() []string {
	result := []string{}
	for _, book := range l.bookList {
		result = append(result, book.Title)
	}

	return result
}

// LessonsUntil returns all lessons from a series until the
// provided lesson (included). If the lesson is not found,
// all lessons are returned. If the book is not found, an empty
// list is returned.
func (l Library) LessonsUntil(booktitle, lessontitle string) []*Lesson {
	result := []*Lesson{}
	foundBook, ok := l.booksByTitle[booktitle]
	if !ok {
		return result
	}

	// if the book is part of a series, get all books, if not,
	// use the found book.
	bookList := []*Book{foundBook}
	if foundBook.SeriesTitle != "" {
		bookList = l.BySeriesTitle(foundBook.SeriesTitle)
	}

	found := false

	// find the lessons of all books, cumulate lessons.
LOOP:
	for _, book := range bookList {
		for _, lesson := range book.Lessons() {
			result = append(result, lesson)
			if lesson.Title == lessontitle &&
				lesson.BookTitle == booktitle {
				found = true

				break LOOP
			}
		}
	}

	if !found {
		return []*Lesson{}
	}

	return result
}
