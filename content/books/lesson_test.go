package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

var card1, card2, card3 = "card1", "card2", "card3"
var book1 = "book1"
var lesson1 = "lesson1"

func TestLessonContent(t *testing.T) {
	testCases := []struct {
		name        string
		lesson      books.Lesson
		wantContent []string
		wantCard1   bool
	}{{
		name:        "uninitialized",
		wantContent: []string{},
	}, {
		name:        "empty",
		lesson:      books.NewLesson(lesson1, book1),
		wantContent: []string{},
	}, {
		name:        "one card",
		lesson:      books.NewLesson(lesson1, book1, card1),
		wantContent: []string{card1},
		wantCard1:   true,
	}, {
		name:        "one nonmatching card",
		lesson:      books.NewLesson(lesson1, book1, card2),
		wantContent: []string{card2},
		wantCard1:   false,
	}, {
		name:        "three cards",
		lesson:      books.NewLesson(lesson1, book1, card1, card2, card3),
		wantContent: []string{card1, card2, card3},
		wantCard1:   true,
	}, {
		name: "duplicate cards",
		lesson: books.NewLesson(lesson1, book1,
			card1, card2, card3, card1, card2, card3),
		wantContent: []string{card1, card2, card3},
		wantCard1:   true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.lesson.Content()
			if diff := cmp.Diff(got, c.wantContent); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
			if got := c.lesson.Contains(card1); got != c.wantCard1 {
				t.Errorf("ERROR: got %v, want %v", got, c.wantCard1)
			}
		})
	}
}

func TestLessonAdd(t *testing.T) {

	testCases := []struct {
		name     string
		in, want []string
	}{
		{
			name: "add nil",
			want: []string{},
		},
		{
			name: "add nil",
			in:   []string{},
			want: []string{},
		},
		{
			name: "add one",
			in:   []string{card1},
			want: []string{card1},
		},
		{
			name: "add three",
			in:   []string{card1, card2, card3},
			want: []string{card1, card2, card3},
		},
		{
			name: "add duplicates",
			in:   []string{card1, card2, card3, card2},
			want: []string{card1, card2, card3},
		},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			var lesson books.Lesson
			lesson.Add(c.in...)
			got := lesson.Content()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
		})
	}

	t.Run("repeated add content", func(t *testing.T) {
		var lesson books.Lesson

		cards := []string{card1, card2, card3}
		for i, card := range cards {
			lesson.Add(card)
			got := lesson.Content()
			want := cards[:i+1]
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
		}
	})
}
