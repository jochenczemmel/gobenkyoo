// Package books provides information about Books, Lessons and Libraries.
// It also provides access to the content kanji and word cards.
package books

// Book represents a book with lessons that can contain
// words and kanjis. It is optionally
// a volume of a series/collection of books.
// The lesson order is preserved.
type Book struct {
	ID
	lessonNames   []string
	lessonsByName map[string]Lesson
}

// New returns a new book with the specified information.
func New(id ID) Book {
	return Book{
		ID:            id,
		lessonNames:   []string{},
		lessonsByName: map[string]Lesson{},
	}
}

// LessonNames returns the odered list of names of all lessons.
func (b Book) LessonNames() []string {
	return b.lessonNames
}

// Lessons returns the odered list of lessons.
func (b Book) Lessons() []Lesson {
	result := make([]Lesson, 0, len(b.lessonNames))
	for _, name := range b.lessonNames {
		result = append(result, b.Lesson(name))
	}
	return result
}

// SetLessons adds or replaces the specified lessons.
// The order of newly added lessons is preserved, the order
// of replaced lessons is not preserved.
func (b *Book) SetLessons(lessons ...Lesson) {
	for _, lesson := range lessons {
		_, ok := b.lessonsByName[lesson.Name]
		if !ok {
			b.lessonNames = append(b.lessonNames, lesson.Name)
		}
		b.lessonsByName[lesson.Name] = lesson
	}
}

// Lesson returns the lesson with the given title.
// If it is not found, a new Lesson is returned.
func (b Book) Lesson(lessonname string) Lesson {
	found, ok := b.lessonsByName[lessonname]
	if !ok {
		return NewLesson(lessonname)
	}
	return found
}
