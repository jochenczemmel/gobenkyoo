package books_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

func TestID(t *testing.T) {
	testCases := []struct {
		name string
		id   books.ID
		want string
	}{{
		name: "title only",
		id:   books.NewID("nihongo", "", 0),
		want: "nihongo",
	}, {
		name: "title and series",
		id:   books.NewID("nihongo", "minna", 0),
		want: "nihongo (minna)",
	}, {
		name: "title and series",
		id:   books.NewID("nihongo", "minna", 1),
		want: "nihongo (minna - 1)",
	}, {
		name: "series title only",
		id:   books.NewID("", "nihongo", 0),
		want: "nihongo (nihongo)",
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.id.String()
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}
