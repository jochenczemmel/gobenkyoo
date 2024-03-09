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
		lesson, ok := b.Lesson(name)
		if ok {
			result = append(result, lesson)
		}
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

// Lesson returns the lesson with the given title and true.
// If it is not found, an empty Lesson and false is returned.
func (b Book) Lesson(lessonname string) (Lesson, bool) {
	found, ok := b.lessonsByName[lessonname]
	return found, ok
}

// // GetKanjiCard returns the kanji card with the id from the lesson.
// // If the lesson or id is not found, an empty card is returned.
// func (b Book) getKanjiCard(lessonname, id string) kanjis.Card {
// 	return b.lessonsByName[lessonname].KanjiCard(id)
// }
//
// // GetWordCard returns the word card with the id from the lesson.
// // If the lesson or id is not found, an empty card is returned.
// func (b Book) getWordCard(lessonname, id string) words.Card {
// 	return b.lessonsByName[lessonname].WordCard(id)
// }
