package books

import "fmt"

// Lesson represents a single lesson within a book.
type Lesson struct {
	Title       string
	BookTitle   string
	content     []string
	uniqContent map[string]bool
}

// NewLesson returns a new Lesson object with the
// given title.
func NewLesson(title, booktitle string, ids ...string) Lesson {
	lesson := Lesson{
		Title:       title,
		BookTitle:   booktitle,
		content:     []string{},
		uniqContent: map[string]bool{},
	}
	lesson.Add(ids...)

	return lesson
}

// String displays the lesson metadata.
func (l *Lesson) String() string {
	return fmt.Sprintf("%s (%s)", l.Title, l.BookTitle)
}

// Add adds ids to the lesson. Duplicates are ignored.
func (l *Lesson) Add(ids ...string) {
	if l.uniqContent == nil {
		l.uniqContent = map[string]bool{}
	}
	for _, id := range ids {
		if l.uniqContent[id] {
			continue
		}
		l.content = append(l.content, id)
		l.uniqContent[id] = true
	}
}

// Content returns the Content (= id-list) of the lesson.
func (l Lesson) Content() []string {
	if l.content == nil {
		return []string{}
	}
	result := make([]string, len(l.content))
	copy(result, l.content)

	return result
}

// Contains returns true if the given id is in the lesson.
func (l Lesson) Contains(id string) bool {
	_, ok := l.uniqContent[id]

	return ok
}
