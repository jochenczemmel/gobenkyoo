package store

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// LibraryImporter provides importing of library data.
type LibraryImporter struct {
	Library       books.Library
	kanjiImporter KanjiImporter
	wordImporter  WordImporter
}

// NewKanjiLibraryImporter returns a library importer
// that can import a kanji file.
func NewKanjiLibraryImporter(library books.Library, importer KanjiImporter) LibraryImporter {
	return LibraryImporter{
		Library:       library,
		kanjiImporter: importer,
	}
}

// NewWordLibraryImporter returns a library importer
// that can import a word file.
func NewWordLibraryImporter(library books.Library, importer WordImporter) LibraryImporter {
	return LibraryImporter{
		Library:      library,
		wordImporter: importer,
	}
}

// Lesson imports a single lesson from a single file.
func (l *LibraryImporter) Lesson(filename string, lessonid books.LessonID) error {

	book := l.Library.Book(lessonid.ID)
	lesson := book.Lesson(lessonid.Name)

	switch {
	case l.kanjiImporter != nil:
		cards, err := l.kanjiImporter.ImportKanji(filename)
		if err != nil {
			return err
		}
		lesson.AddKanjis(cards...)
	case l.wordImporter != nil:
		cards, err := l.wordImporter.ImportWord(filename)
		if err != nil {
			return err
		}
		lesson.AddWords(cards...)
	default:
		return ConfigurationError("no importer defined")
	}

	book.SetLessons(lesson)
	l.Library.SetBooks(book)

	return nil
}
