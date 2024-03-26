package store

import (
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// LibraryImporter provides importing of library data.
type LibraryImporter struct {
	Library       books.Library
	kanjiImporter KanjiImporter
	wordImporter  WordImporter
}

// NewKanjiImporter returns a library importer that can import a kanji file.
func NewKanjiImporter(importer KanjiImporter) LibraryImporter {
	return LibraryImporter{
		Library:       books.NewLibrary(cfg.DefaultLibrary),
		kanjiImporter: importer,
	}
}

// NewWordImporter returns a library importer that can import a word file.
func NewWordImporter(importer WordImporter) LibraryImporter {
	return LibraryImporter{
		Library:      books.NewLibrary(cfg.DefaultLibrary),
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
