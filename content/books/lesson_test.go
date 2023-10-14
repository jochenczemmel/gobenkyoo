package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

type DummyCard struct{ Identifier string }

func (c DummyCard) ID() string { return c.Identifier }

func TestLesson(t *testing.T) {
	title := "lesson1"
	bookTitle := "book1"

	content := []DummyCard{
		{"card1"},
		{"card2"},
		{"card3"},
	}

	t.Run("empty lesson", func(t *testing.T) {
		lesson := books.NewLesson[DummyCard](title, bookTitle)
		got := lesson.Title()
		if got != title {
			t.Errorf("ERROR: got %s, want %s", got, title)
		}
		got = lesson.BookTitle()
		if got != bookTitle {
			t.Errorf("ERROR: got %s, want %s", got, bookTitle)
		}
	})

	t.Run("initial filled", func(t *testing.T) {
		lesson := books.NewLesson(title, bookTitle, content...)
		got := lesson.Content()
		want := content
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got-, want+%s\n", diff)
		}
	})

	t.Run("add content", func(t *testing.T) {
		lesson := books.NewLesson[DummyCard](title, bookTitle)
		lesson.Add(content...)
		got := lesson.Content()
		want := content
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got-, want+%s\n", diff)
		}
	})

	t.Run("add duplicates", func(t *testing.T) {
		lesson := books.NewLesson(title, bookTitle, content...)
		lesson.Add(content...)
		got := lesson.Content()
		want := content
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("ERROR: got-, want+%s\n", diff)
		}
	})

	t.Run("contains", func(t *testing.T) {
		lesson := books.NewLesson(title, bookTitle, content...)

		cases := []struct {
			name string
			id   string
			want bool
		}{
			{name: "found", id: "card1", want: true},
			{name: "not found", id: "notfound", want: false},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := lesson.Contains(DummyCard{c.id})
				if got != c.want {
					t.Errorf("ERROR: got %v, want %v", got, c.want)
				}
			})
			t.Run("id_"+c.name, func(t *testing.T) {
				got := lesson.ContainsID(c.id)
				if got != c.want {
					t.Errorf("ERROR: got %v, want %v", got, c.want)
				}
			})
		}
	})
}
