// Package store handles importing, reading and storing of all
// kinds of application content.
package store

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

const DBTypeJSON = "json"

var DbTypes = []string{dbTypeJson}

// KanjiImporter defines the interface to import kanji cards.
type KanjiImporter interface {
	ImportKanji(string) ([]kanjis.Card, error)
}

// WordImporter defines the interface to import word cards.
type WordImporter interface {
	ImportWord(string) ([]words.Card, error)
}

// LibraryLoader defines the interface to load
// kanji and word content.
type LibraryLoader interface {
	LoadLibrary(string) (books.Library, error)
}

// LibraryStorer defines the interface to store
// kanji and word content.
type LibraryStorer interface {
	StoreLibrary(books.Library) error
}

// ClassroomLoader defines the interface to load learn boxes.
type ClassroomLoader interface {
	LoadClassroom(string) (learn.Classroom, error)
}

// ClassroomStorer defines the interface to store learn boxes.
type ClassroomStorer interface {
	StoreClassroom(learn.Classroom) error
}

// LibraryLoadStorer defines the interface to load and store
// kanji and word content.
type LibraryLoadStorer interface {
	LibraryLoader
	LibraryStorer
}

// LoadStorer defines the interface to load and store
// word and kanji content and learn boxes.
type LoadStorer interface {
	LibraryLoader
	LibraryStorer
	ClassroomLoader
	ClassroomStorer
}

// Loader can load learn boxes and kanji and word content.
type Loader interface {
	LibraryLoader
	ClassroomLoader
}
