package books

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Lesson represents a single Lesson within a book.
type Lesson struct {
	Name       string        // lesson name
	wordCards  []words.Card  // word cards
	kanjiCards []kanjis.Card // kanji cards
}

// NewLesson returns an initialized Lesson with the given name.
func NewLesson(name string) Lesson {
	return Lesson{
		Name:       name,
		wordCards:  []words.Card{},
		kanjiCards: []kanjis.Card{},
	}
}

// AddKanjis adds a list of kanji cards to the lesson.
func (l *Lesson) AddKanjis(cards ...kanjis.Card) {
	if l.kanjiCards == nil {
		l.kanjiCards = make([]kanjis.Card, 0, len(cards))
	}
	l.kanjiCards = append(l.kanjiCards, cards...)
}

// AddWords adds a list of word cards to the lesson.
func (l *Lesson) AddWords(cards ...words.Card) {
	if l.wordCards == nil {
		l.wordCards = make([]words.Card, 0, len(cards))
	}
	l.wordCards = append(l.wordCards, cards...)
}

// KanjiCard returns the kanji card with the given id.
// If it is not found, an empty card is returned.
func (l Lesson) KanjiCard(id string) kanjis.Card {
	for _, card := range l.kanjiCards {
		if card.ID == id {
			return card
		}
	}

	return kanjis.Card{}
}

// WordCard returns the word card with the given id.
// If it is not found, an empty card is returned.
func (l Lesson) WordCard(id string) words.Card {
	for _, card := range l.wordCards {
		if card.ID == id {
			return card
		}
	}

	return words.Card{}
}

// KanjiCards returns all the kanji cards in the lesson.
func (l Lesson) KanjiCards() []kanjis.Card {
	if l.kanjiCards == nil {
		return []kanjis.Card{}
	}
	return l.kanjiCards
}

// WordCards returns all the word cards in the lesson.
func (l Lesson) WordCards() []words.Card {
	if l.wordCards == nil {
		return []words.Card{}
	}
	return l.wordCards
}
