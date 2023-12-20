package books

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Lesson represents a single lesson within a book.
type Lesson struct {
	Title      string
	BookTitle  string
	WordCards  []*words.Card
	KanjiCards []*kanjis.Card
}

// String displays the lesson metadata.
// Mainly used for debugging.
func (l *Lesson) String() string {
	return fmt.Sprintf("%s (%s)", l.Title, l.BookTitle)
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
