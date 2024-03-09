package app

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

type LibraryLoadStorer interface {
	LibraryLoader
	LibraryStorer
}

type LibraryLoader interface {
	LoadLibrary(string) (books.Library, error)
}

type LibraryStorer interface {
	StoreLibrary(books.Library) error
}

type KanjiImporter interface {
	ImportKanji(string) ([]kanjis.Card, error)
}

type WordImporter interface {
	ImportWord(string) ([]words.Card, error)
}

/*
type DataLoadStorer interface {
	DataLoader
	DataStorer
}

type DataLoader interface {
	LoadClassroom(string) (learn.Classroom, error)
	LoadLibrary(string) (books.Library, error)
}

type DataStorer interface {
	StoreClassroom(learn.Classroom) error
	StoreLibrary(books.Library) error
}
*/
