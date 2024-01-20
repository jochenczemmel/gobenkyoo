// Package books provides information about Books, Lessons and Libraqies.
// It alos provides access to the content kanji and word cards.
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
	lessonTitles  []string
	lessonsByName map[string]lesson
}

// New returns a new book with the specified information.
// func New(title, seriestitle string, volume int) Book {
func New(id ID) Book {
	return Book{
		ID:            id,
		lessonTitles:  []string{},
		lessonsByName: map[string]lesson{},
	}
}

// LessonTitles returns the odered list of titles of all lessons.
func (b Book) LessonTitles() []string {
	return b.lessonTitles
}

// AddKanjis adds the list of kanjis to the specified lesson.
// If the lesson does not exist, it is created.
// The order of the added lessons is preserved.
func (b *Book) AddKanjis(lessontitle string, cards ...kanjis.Card) {
	currentLesson := b.setupLesson(lessontitle)
	currentLesson.addKanjis(cards...)
	b.lessonsByName[lessontitle] = currentLesson
}

// AddWords adds the list of words to the specified lesson.
// If the lesson does not exist, it is created.
// The order of the added lessons is preserved.
func (b *Book) AddWords(lessontitle string, cards ...words.Card) {
	currentLesson := b.setupLesson(lessontitle)
	currentLesson.addWords(cards...)
	b.lessonsByName[lessontitle] = currentLesson
}

// setupLesson returns the lesson with the given title.
// If it does not exist, an new lesson is created.
func (b *Book) setupLesson(lessontitle string) lesson {
	currentLesson, ok := b.lessonsByName[lessontitle]
	if !ok {
		b.lessonTitles = append(b.lessonTitles, lessontitle)
		return newLesson(lessontitle)
	}
	return currentLesson
}

// KanjisFor returns the list of kanjis for the given lesson.
// If the lesson does not exist, an empty list is returned.
func (b Book) KanjisFor(lessontitle string) []kanjis.Card {
	currentLesson, ok := b.lessonsByName[lessontitle]
	if !ok {
		return []kanjis.Card{}
	}
	return currentLesson.kanjiCards
}

// WordsFor returns the list of words for the given lesson.
// If the lesson does not exist, an empty list is returned.
func (b Book) WordsFor(lessontitle string) []words.Card {
	currentLesson, ok := b.lessonsByName[lessontitle]
	if !ok {
		return []words.Card{}
	}
	return currentLesson.wordCards
}

// FindKanjiCard returns the kanji card with the id from the lesson.
// If the lesson or id is not found, an empty card is returned.
func (b Book) FindKanjiCard(lessontitle string, id int) kanjis.Card {
	return b.lessonsByName[lessontitle].findKanjiCard(id)
}

// FindWordCard returns the word card with the id from the lesson.
// If the lesson or id is not found, an empty card is returned.
func (b Book) FindWordCard(lessontitle string, id int) words.Card {
	return b.lessonsByName[lessontitle].findWordCard(id)
}
