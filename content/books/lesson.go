package books

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// lesson represents a single lesson within a book.
type lesson struct {
	title      string        // Lesson title
	wordCards  []words.Card  // word cards
	kanjiCards []kanjis.Card // kanji cards
}

func newLesson(title string) lesson {
	return lesson{
		title:      title,
		wordCards:  []words.Card{},
		kanjiCards: []kanjis.Card{},
	}
}

func (l *lesson) addKanjis(cards ...kanjis.Card) {
	l.kanjiCards = append(l.kanjiCards, cards...)
}

func (l *lesson) addWords(cards ...words.Card) {
	l.wordCards = append(l.wordCards, cards...)
}
