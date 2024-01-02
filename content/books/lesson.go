package books

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Lesson represents a single lesson within a book.
type Lesson struct {
	Book       TitleInfos     // book and series title, volume number
	Title      string         // Lesson title
	WordCards  []*words.Card  // word cards
	KanjiCards []*kanjis.Card // kanji cards
}

// Contains returns true if the given word card is in the lesson.
func (l Lesson) Contains(card any) bool {
	switch card.(type) {
	case *kanjis.Card:
		for _, c := range l.KanjiCards {
			if c == card {
				return true
			}
		}
	case *words.Card:
		for _, c := range l.WordCards {
			if c == card {
				return true
			}
		}
	}

	return false
}
