// Package books provides information about Books, Lessons and Libraqies.
// It alos provides access to the content kanji and word cards.
package books

import "github.com/jochenczemmel/gobenkyoo/content/kanjis"

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
// The lesson order is preserved.
type Book struct {
	TitleInfo
	Lessons       []Lesson
	lessonsByName map[string]Lesson
}

// New returns a new book with the specified infos.
func New(title, seriestitle string, volume int, lessons ...Lesson) Book {
	return Book{
		TitleInfo: TitleInfo{
			Title:       title,
			SeriesTitle: seriestitle,
			Volume:      volume,
		},
		Lessons:       lessons,
		lessonsByName: make(map[string]Lesson),
	}
}

func (b *Book) AddKanjis(lessontitle string, cards ...kanjis.Card) {
	lesson, ok := b.lessonsByName[lessontitle]
	if !ok {
		lesson = Lesson{
			Title: lessontitle,
			Book:  b.TitleInfo,
		}
	}
	lesson.KanjiCards = append(lesson.KanjiCards, cards...)
	b.lessonsByName[lessontitle] = lesson
}

func (b Book) KanjiFor(lessontitle string) []kanjis.Card {
	return b.lessonsByName[lessontitle].KanjiCards
}
