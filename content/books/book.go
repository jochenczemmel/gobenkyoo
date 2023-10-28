// Package books provides informaion about Books, Lessons and Content.
package books

import "github.com/jochenczemmel/gobenkyoo/content"

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book[T content.Card] struct {
	title         string                // the title of the book
	lessons       []*Lesson[T]          // the lessons in the book
	lessonByTitle map[string]*Lesson[T] // the lessons in the book
}

// New returns a new book with the given title.
func New[T content.Card](title string, lessons ...Lesson[T]) Book[T] {
	result := Book[T]{
		title:         title,
		lessonByTitle: map[string]*Lesson[T]{},
	}
	result.AppendLesson(lessons...)

	return result
}

// Title returns the book title.
func (b Book[T]) Title() string {
	return b.title
}

// Lessons returns all lessons in the book in the provided order.
func (b Book[T]) Lessons() []Lesson[T] {
	result := make([]Lesson[T], 0, len(b.lessons))
	for _, l := range b.lessons {
		result = append(result, *l)
	}

	return result
}

// TODO: how should this work?
// AppendLesson adds lessons to the book. The order of the lessons
// is preserved.
func (b *Book[T]) AppendLesson(lessons ...Lesson[T]) {

	for _, l := range lessons {
		lesson := l
		_, ok := b.lessonByTitle[lesson.Title()]
		if ok {
			return
		}
		lesson.SetBookTitle(b.title)
		b.lessonByTitle[lesson.Title()] = &lesson
		b.lessons = append(b.lessons, &lesson)
	}
}

// Lesson returns the lesson with the given title.
// If it is not found, it returns an empty Lesson and false.
func (b Book[T]) Lesson(title string) (Lesson[T], bool) {
	lesson, ok := b.lessonByTitle[title]
	if !ok {
		return Lesson[T]{}, false
	}

	return *lesson, true
}

/*
// TODO: is it necessary to change a lesson or set to a certain position?
// SetLesson changes the value of an existing Lesson.
func (b *Book[T]) SetLesson(lessons Lesson[T], position int) {
}
*/
