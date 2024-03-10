package books

import (
	"strconv"

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
	idValue := len(l.kanjiCards)
	for i := range cards {
		idValue++
		cards[i].ID = strconv.Itoa(idValue)
	}
	l.kanjiCards = append(l.kanjiCards, cards...)
}

// AddWords adds a list of word cards to the lesson.
func (l *Lesson) AddWords(cards ...words.Card) {
	if l.wordCards == nil {
		l.wordCards = make([]words.Card, 0, len(cards))
	}
	idValue := len(l.wordCards)
	for i := range cards {
		idValue++
		cards[i].ID = strconv.Itoa(idValue)
	}
	l.wordCards = append(l.wordCards, cards...)
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
