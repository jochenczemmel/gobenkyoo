// Package books provides informaion about Books, Lessons and Content.
package books

type Series struct {
	SeriesTitle string // the title of the book collection
	Books       []*Book
}

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book struct {
	Title         string             // the title of the book
	SeriesTitle   string             // the title of the book collection
	Volume        int                // the volume number in the collection
	lessons       []string           // the ordered lessons titles
	lessonByTitle map[string]*Lesson // the lessons by title
}

// New returns a new book with the given title.
func New(title, seriestitle string, lessons ...*Lesson) Book {
	book := Book{
		Title:         title,
		SeriesTitle:   title,
		lessons:       []string{},
		lessonByTitle: map[string]*Lesson{},
	}
	if seriestitle != "" {
		book.SeriesTitle = seriestitle
	}
	book.Add(lessons...)

	return book
}

// Lessons returns all lessons in the book in the
// provided order.
func (b Book) Lessons() []*Lesson {
	if b.lessons == nil || b.lessonByTitle == nil {
		return []*Lesson{}
	}
	result := make([]*Lesson, 0, len(b.lessons))
	for _, title := range b.lessons {
		result = append(result, b.lessonByTitle[title])
	}

	return result
}

// Add adds lessons to the book. The order of the lessons
// is preserved. If the lesson already exists, it is replaced.
func (b *Book) Add(lessons ...*Lesson) {
	if b.lessonByTitle == nil {
		b.lessonByTitle = map[string]*Lesson{}
	}
	for _, lesson := range lessons {
		_, ok := b.lessonByTitle[lesson.Title]
		if ok {
			b.lessonByTitle[lesson.Title] = lesson

			continue
		}
		b.lessonByTitle[lesson.Title] = lesson
		b.lessons = append(b.lessons, lesson.Title)
	}
}
