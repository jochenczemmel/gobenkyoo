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

		book := books.New[content.Card](bookTitle)
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
		got := book.Lessons()
		gotTitles := []string{}
		for _, l := range got {
			gotTitles = append(gotTitles, l.Title())
		}

		if diff := cmp.Diff(gotTitles, wantLessonTitles); diff != "" {
			t.Errorf("ERROR: got- want+%s\n", diff)
		}
	})

	t.Run("add content", func(t *testing.T) {

		book := books.New[content.Card](bookTitle)
		book.AddLesson(lessons...)
		got := book.Lessons()
		gotTitles := []string{}
		for _, l := range got {
			gotTitles = append(gotTitles, l.Title())
		}

		if diff := cmp.Diff(gotTitles, wantLessonTitles); diff != "" {
			t.Errorf("ERROR: got- want+%s\n", diff)
		}
	})
}
