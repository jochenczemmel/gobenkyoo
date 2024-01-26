// Package books provides information about Books, Lessons and Libraries.
// It also provides access to the content kanji and word cards.
package books

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Book represents a book with lessons that can contain
// words and kanjis. It is optionally
// a volume of a series/collection of books.
// The lesson order is preserved.
type Book struct {
	ID
	lessonNames   []string
	lessonsByName map[string]lesson
}

// New returns a new book with the specified information.
// func New(title, seriestitle string, volume int) Book {
func New(id ID) Book {
	return Book{
		ID:            id,
		lessonNames:   []string{},
		lessonsByName: map[string]lesson{},
	}
}

// LessonNames returns the odered list of names of all lessons.
func (b Book) LessonNames() []string {
	return b.lessonNames
}

// AddKanjis adds the list of kanjis to the specified lesson.
// If the lesson does not exist, it is created.
// The order of the added lessons is preserved.
func (b *Book) AddKanjis(lessonname string, cards ...kanjis.Card) {
	currentLesson := b.setupLesson(lessonname)
	currentLesson.addKanjis(cards...)
	b.lessonsByName[lessonname] = currentLesson
}

// AddWords adds the list of words to the specified lesson.
// If the lesson does not exist, it is created.
// The order of the added lessons is preserved.
func (b *Book) AddWords(lessonname string, cards ...words.Card) {
	currentLesson := b.setupLesson(lessonname)
	currentLesson.addWords(cards...)
	b.lessonsByName[lessonname] = currentLesson
}

// setupLesson returns the lesson with the given name.
// If it does not exist, an new lesson is created.
func (b *Book) setupLesson(lessonname string) lesson {
	currentLesson, ok := b.lessonsByName[lessonname]
	if !ok {
		b.lessonNames = append(b.lessonNames, lessonname)
		return newLesson(lessonname)
	}
	return currentLesson
}

// KanjisFor returns the list of kanjis for the given lesson.
// If the lesson does not exist, an empty list is returned.
func (b Book) KanjisFor(lessonname string) []kanjis.Card {
	currentLesson, ok := b.lessonsByName[lessonname]
	if !ok {
		return []kanjis.Card{}
	}
	return currentLesson.kanjiCards
}

// WordsFor returns the list of words for the given lesson.
// If the lesson does not exist, an empty list is returned.
func (b Book) WordsFor(lessonname string) []words.Card {
	currentLesson, ok := b.lessonsByName[lessonname]
	if !ok {
		return []words.Card{}
	}
	return currentLesson.wordCards
}

// GetKanjiCard returns the kanji card with the id from the lesson.
// If the lesson or id is not found, an empty card is returned.
func (b Book) getKanjiCard(lessonname string, id int) kanjis.Card {
	return b.lessonsByName[lessonname].getKanjiCard(id)
}

// GetWordCard returns the word card with the id from the lesson.
// If the lesson or id is not found, an empty card is returned.
func (b Book) getWordCard(lessonname string, id int) words.Card {
	return b.lessonsByName[lessonname].getWordCard(id)
}
