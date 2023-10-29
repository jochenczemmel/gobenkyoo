// Package books provides informaion about Books, Lessons and Content.
package books

type series struct {
	seriesTitle string // the title of the book collection
	Books       []*Book
}

// Book represents a book with lessons. It is optionally
// a volume of a series/collection of books.
type Book struct {
	title         string             // the title of the book
	seriesTitle   string             // the title of the book collection
	volume        int                // the volume number in the collection
	lessons       []string           // the ordered lessons titles
	lessonByTitle map[string]*Lesson // the lessons by title
}

// New returns a new book with the given title.
func New(title, seriestitle string) Book {
	book := Book{
		title:         title,
		seriesTitle:   title,
		lessons:       []string{},
		lessonByTitle: map[string]*Lesson{},
	}
	if seriestitle != "" {
		book.seriesTitle = seriestitle
	}

	return book
}

// Title returns the book title.
func (b Book) Title() string {
	return b.title
}

// SeriesTitle returns the title of the series
// in which the book is contained.
func (b Book) SeriesTitle() string {
	return b.seriesTitle
}

// VolumeNumber returns the number of the book in the series.
func (b Book) VolumeNumber() int {
	return b.volume
}

// Lessons returns all lessons in the book in the
// provided order.
func (b Book) Lessons() []*Lesson {
	if b.lessons == nil || b.lessonByTitle == nil {
		return []*Lesson{}
	}
	result := make([]*Lesson, len(b.lessons))
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
		title := lesson.Title()
		_, ok := b.lessonByTitle[title]
		if ok {
			b.lessonByTitle[title] = lesson
			continue
		}
		b.lessonByTitle[title] = lesson
		b.lessons = append(b.lessons, title)
	}
}

// SetTitle sets the book title.
func (b *Book) SetTitle(title string) {
	b.title = title
}

// SetSeriesTitle sets the book series title.
func (b *Book) SetSeriesTitle(title string) {
	b.seriesTitle = title
}

// SetVolume sets the book series volume.
func (b *Book) SetVolume(volume int) {
	b.volume = volume
}
