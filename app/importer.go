// Package app provides common application logic.
package app

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store"
)

// LibraryImporter provides importing of library data.
type LibraryImporter struct {
	libName    string
	loadStorer LibraryLoadStorer
}

// NewLibraryImporter returns an importer that uses the given loadstorer.
func NewLibraryImporter(libname string, ls store.LibraryLoadStorer) LibraryImporter {
	return LibraryImporter{
		libName:    libname,
		loadStorer: ls,
	}
}

func (l LibraryImporter) Kanji(importer store.KanjiImporter, filename string, lessonid books.LessonID) error {

	library, found, err := l.loadStorer.LoadLibrary(l.libName)
	if err != nil && found {
		return err
	}

	libImporter := store.NewKanjiLibraryImporter(library, importer)
	err = libImporter.Lesson(filename, lessonid)
	if err != nil {
		return err
	}

	return l.loadStorer.StoreLibrary(library)
}

func (l LibraryImporter) Word(importer store.WordImporter, filename string, lessonid books.LessonID) error {

	library, found, err := l.loadStorer.LoadLibrary(l.libName)
	if err != nil && found {
		return err
	}

	libImporter := store.NewWordLibraryImporter(library, importer)
	err = libImporter.Lesson(filename, lessonid)
	if err != nil {
		return err
	}

	return l.loadStorer.StoreLibrary(library)
}
