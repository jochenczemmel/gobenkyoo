// Package app provides common application logic.
package app

import (
	"errors"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// LibraryImporter provides importing of library data.
type LibraryImporter struct {
	loadStorer    LibraryLoadStorer
	kanjiImporter KanjiImporter
	wordImporter  WordImporter
	Library       books.Library
}

// NewLibraryImporter returns an importer that uses the given loadstorer.
func NewLibraryImporter(loadstorer LibraryLoadStorer) LibraryImporter {
	return LibraryImporter{
		loadStorer: loadstorer,
	}
}

// SetKanjiImporter sets the importer to use for kanji imports.
func (li *LibraryImporter) SetKanjiImporter(importer KanjiImporter) {
	li.kanjiImporter = importer
}

// SetWordImporter sets the importer to use for word imports.
func (li *LibraryImporter) SetWordImporter(importer WordImporter) {
	li.wordImporter = importer
}

// LoadLibrary loads the library with the given name.
// It returns true if it is found, false if it is not found.
// In case of another error, the error is returned.
func (li *LibraryImporter) LoadLibrary(name string) (found bool, err error) {

	if li.loadStorer == nil {
		return false, ConfigurationError("no LibraryLoadStorer defined")
	}

	li.Library, err = li.loadStorer.LoadLibrary(name)
	if err == nil {
		return true, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		return false, nil
	}

	return false, err
}

// StoreLibrary stores the loaded library.
func (li LibraryImporter) StoreLibrary() error {
	if li.loadStorer == nil {
		return ConfigurationError("no LibraryLoadStorer defined")
	}
	return li.loadStorer.StoreLibrary(li.Library)
}

// WordLesson imports the word data from the file into the given lesson.
func (li *LibraryImporter) WordLesson(filename string, lessonid books.LessonID) error {
	return li.doImport(learn.WordType, filename, lessonid)
}

// KanjiLesson imports the kanji data from the file into the given lesson.
func (li *LibraryImporter) KanjiLesson(filename string, lessonid books.LessonID) error {
	return li.doImport(learn.KanjiType, filename, lessonid)
}

// doImport does the kanji and word import.
func (li *LibraryImporter) doImport(typ, filename string, lessonid books.LessonID) (err error) {

	book := li.Library.Book(lessonid.ID)
	lesson, ok := book.Lesson(lessonid.Name)
	if !ok {
		lesson.Name = lessonid.Name
	}

	switch typ {
	case learn.KanjiType:
		err = li.importKanji(filename, &lesson)
	default:
		err = li.importWord(filename, &lesson)
	}
	if err != nil {
		return err
	}

	book.SetLessons(lesson)
	li.Library.SetBooks(book)

	return nil
}

// importKanji imports kanjis from the file and adds
// the cards to the lesson.
func (li *LibraryImporter) importKanji(filename string, lesson *books.Lesson) error {
	if li.kanjiImporter == nil {
		return ConfigurationError("no KanjiImporter defined")
	}
	cards, err := li.kanjiImporter.ImportKanji(filename)
	if err != nil {
		return err
	}
	lesson.AddKanjis(cards...)

	return nil
}

// importWord imports words from the file and adds
// the cards to the lesson.
func (li *LibraryImporter) importWord(filename string, lesson *books.Lesson) error {

	if li.wordImporter == nil {
		return ConfigurationError("no WordImporter defined")
	}

	cards, err := li.wordImporter.ImportWord(filename)
	if err != nil {
		return err
	}
	lesson.AddWords(cards...)

	return nil
}
