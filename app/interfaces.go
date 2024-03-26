// Package app provides common application logic.
package app

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// KanjiImporter defines the interface to import kanji cards.
type KanjiImporter interface {
	ImportKanji(string) ([]kanjis.Card, error)
}

// WordImporter defines the interface to import word cards.
type WordImporter interface {
	ImportWord(string) ([]words.Card, error)
}

// LibraryLoadStorer defines the interface to load and store
// kanji and word content.
type LibraryLoadStorer interface {
	LibraryLoader
	LibraryStorer
}

// ClassroomLoadStorer defines the interface to update
// learn boxes from kanji and word content.
type ClassroomLoadStorer interface {
	LibraryLoader
	ClassroomLoader
	ClassroomStorer
}

// LoadStorer defines the interface to load and store
// word and kanji content and learn boxes.
type LoadStorer interface {
	LibraryLoader
	LibraryStorer
	ClassroomLoader
	ClassroomStorer
}

// LibraryLoader defines the interface to load
// kanji and word content.
type LibraryLoader interface {
	LoadLibrary(string) (books.Library, bool, error)
}

// LibraryStorer defines the interface to store
// kanji and word content.
type LibraryStorer interface {
	StoreLibrary(books.Library) error
}

// ClassroomLoader defines the interface to load learn boxes.
type ClassroomLoader interface {
	LoadClassroom(string) (learn.Classroom, bool, error)
}

// ClassroomStorer defines the interface to store learn boxes.
type ClassroomStorer interface {
	StoreClassroom(learn.Classroom) error
}
