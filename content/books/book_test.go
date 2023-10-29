package books_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

var series1 = "series1"

func TestBookGetters(t *testing.T) {

	lesson1 := books.NewLesson("lesson1", book1)
	lesson2 := books.NewLesson("lesson2", book1)
	lesson3 := books.NewLesson("lesson3", book1)

	var emptyBook books.Book
	emptyBook.Add(&lesson1, &lesson2, &lesson3)

	testCases := []struct {
		name    string
		book    books.Book
		wantLen int
	}{
		{
			name:    "uninitialized",
			wantLen: 0,
		},
		{
			name:    "empty",
			book:    books.New(book1, series1),
			wantLen: 0,
		},
		{
			name:    "one lesson",
			book:    books.New(book1, series1, &lesson1),
			wantLen: 1,
		},
		{
			name:    "three lessons",
			book:    books.New(book1, series1, &lesson1, &lesson2, &lesson3),
			wantLen: 3,
		},
		{
			name:    "empty book three lessons",
			book:    emptyBook,
			wantLen: 3,
		},
		{
			name: "duplicate lessons",
			book: books.New(book1, series1,
				&lesson1, &lesson2, &lesson3,
				&lesson1, &lesson2, &lesson3,
			),
			wantLen: 3,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			if got := len(c.book.Lessons()); got != c.wantLen {
				t.Errorf("ERROR: got %v, want %v", got, c.wantLen)
			}
		})
	}
}
