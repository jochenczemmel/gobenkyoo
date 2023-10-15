package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

func TestBook(t *testing.T) {

	bookTitle := "book1"
	lessons := []books.Lesson[DummyCard]{
		books.NewLesson[DummyCard]("lesson1", bookTitle),
		books.NewLesson[DummyCard]("lesson2", bookTitle),
		books.NewLesson[DummyCard]("lesson3", bookTitle),
	}
	wantLessonTitles := []string{
		"lesson1",
		"lesson2",
		"lesson3",
	}

	t.Run("empty book", func(t *testing.T) {

		book := books.New[DummyCard](bookTitle)
		got := book.Title()
		if got != bookTitle {
			t.Errorf("ERROR: got %s, want %s", got, bookTitle)
		}
		gotLessons := book.Lessons()
		want := 0
		if got := len(gotLessons); got > want {
			t.Errorf("ERROR: got %v, want %v", got, want)
		}
	})

	t.Run("initial content", func(t *testing.T) {
		book := books.New(bookTitle, lessons...)
		compareLessonTitles(t, book.Lessons(), wantLessonTitles)
	})

	t.Run("add content", func(t *testing.T) {
		book := books.New[DummyCard](bookTitle)
		book.AddLesson(lessons...)
		compareLessonTitles(t, book.Lessons(), wantLessonTitles)
	})

	t.Run("add duplicates", func(t *testing.T) {
		book := books.New(bookTitle, lessons...)
		book.AddLesson(lessons...)
		compareLessonTitles(t, book.Lessons(), wantLessonTitles)
	})

	t.Run("add more content", func(t *testing.T) {
		book := books.New(bookTitle, lessons...)
		book.AddLesson(
			books.NewLesson[DummyCard]("lesson4", bookTitle),
		)
		compareLessonTitles(t, book.Lessons(),
			append(wantLessonTitles, "lesson4"),
		)
	})

	t.Run("get lesson by title", func(t *testing.T) {
		book := books.New(bookTitle, lessons...)
		cases := []struct {
			name  string
			title string
			want  bool
		}{
			{name: "found", title: "lesson2", want: true},
			{name: "not found", title: "not found", want: false},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				_, got := book.Lesson(c.title)
				if got != c.want {
					t.Errorf("ERROR: got %v, want %v", got, c.want)
				}
			})
		}

		compareLessonTitles(t, book.Lessons(), wantLessonTitles)
	})
}

func compareLessonTitles[T content.Card](t *testing.T, got []books.Lesson[T], want []string) {
	t.Helper()
	gotTitles := []string{}
	for _, l := range got {
		gotTitles = append(gotTitles, l.Title())
	}

	if diff := cmp.Diff(gotTitles, want); diff != "" {
		t.Errorf("ERROR: got- want+%s\n", diff)
	}
}
