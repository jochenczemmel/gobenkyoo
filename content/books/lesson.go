package books

import "github.com/jochenczemmel/gobenkyoo/content"

// Lesson reprents a single lesson within a book.
// It contains a list of cards.
type Lesson[T content.Card] struct {
	title       string
	bookTitle   string
	content     []T
	uniqContent map[string]bool
}

// NewLesson returns a new Lesson object with the
// given titles and optionally with unique content.
func NewLesson[T content.Card](title, booktitle string, content ...T) Lesson[T] {
	lesson := Lesson[T]{
		title:       title,
		bookTitle:   booktitle,
		uniqContent: map[string]bool{},
	}
	for _, c := range content {
		lesson.Add(c)
	}

	return lesson
}

// Title returns the title of the lesson.
func (l Lesson[T]) Title() string {
	return l.title
}

// BookTitle returns the title of the book that contains the lesson.
func (l Lesson[T]) BookTitle() string {
	return l.bookTitle
}

// Content returns the list of cards.
func (l Lesson[T]) Content() []T {
	return l.content
}

// Add adds Cards to the lesson, duplicates are ignored.
func (l *Lesson[T]) Add(cards ...T) {
	for _, card := range cards {
		if _, ok := l.uniqContent[card.ID()]; ok {
			continue
		}
		l.content = append(l.content, card)
		l.uniqContent[card.ID()] = true
	}
}

// Contains returns true if the given Card is in the lesson.
func (l Lesson[T]) Contains(card T) bool {
	return l.ContainsID(card.ID())
}

// ContainsID returns true if the given id is in the lesson.
func (l Lesson[T]) ContainsID(id string) bool {
	_, ok := l.uniqContent[id]
	return ok
}
