// Package books provides informaion about Books, Lessons and Content.
package books

import "github.com/jochenczemmel/gobenkyoo/content"

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book[T content.Card] struct {
	title         string               // the title of the book
	lessons       []Lesson[T]          // the lessons in the book
	lessonByTitle map[string]Lesson[T] // the lessons in the book
}

// New returns a new book with the given title.
func New[T content.Card](title string, lessons ...Lesson[T]) Book[T] {
	result := Book[T]{
		title:         title,
		lessonByTitle: map[string]Lesson[T]{},
	}
	for _, l := range lessons {
		result.AddLesson(l)
	}
	return result
}

// Title returns the book title.
func (b Book[T]) Title() string {
	return b.title
}

// Lessons returns all lessons in the book in the provided order.
func (b Book[T]) Lessons() []Lesson[T] {
	// return a copy of the slice
	result := make([]Lesson[T], len(b.lessons))
	copy(result, b.lessons)
	return result
}

// AddLesson adds lessons to the book. The order of the lessons
// is preserved.
func (b *Book[T]) AddLesson(lessons ...Lesson[T]) {
	for _, lesson := range lessons {
		foundLesson, ok := b.lessonByTitle[lesson.Title()]
		if ok {
			foundLesson.Add(lesson.Content()...)
			b.lessonByTitle[lesson.Title()] = foundLesson
			return
		}
		b.lessonByTitle[lesson.Title()] = lesson
		b.lessons = append(b.lessons, lesson)
	}
}

/*
// Lesson returns the lesson with the given title.
// It returns false if it is not found.
func (b Book) Lesson(title string) (Lesson, bool) {
	lesson, ok := b.lessonByTitle[title]
	return lesson, ok
}


*/
