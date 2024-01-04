package books

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Lesson represents a single lesson within a book.
type Lesson struct {
	Book       TitleInfos    // book and series title, volume number
	Title      string        // Lesson title
	WordCards  []words.Card  // word cards
	KanjiCards []kanjis.Card // kanji cards
}
