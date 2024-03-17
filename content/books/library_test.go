package books_test

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

func TestLibrarySort(t *testing.T) {

	// prepare test: sorted list of books

	// test multiple times, shuffle may accidentially return the same order
	for i := 0; i < 3; i++ {
		t.Run("shuffle "+strconv.Itoa(i+1), func(t *testing.T) {

			shuffledBooks := slices.Clone(sortedBooks)
			rand.Shuffle(len(shuffledBooks), func(i, j int) {
				shuffledBooks[i], shuffledBooks[j] = shuffledBooks[j], shuffledBooks[i]
			})
			t.Logf("DEBUG: shuffled: first book: %v", shuffledBooks[0])

			library := books.NewLibrary("")
			library.SetBooks(shuffledBooks...)

			t.Run("books", func(t *testing.T) {
				got := library.SortedBooks()
				t.Logf("DEBUG: sorted: first book: %v", got[0])

				if diff := cmp.Diff(got, sortedBooks, cmpopts.IgnoreUnexported(books.Book{})); diff != "" {
					t.Errorf("ERROR: got- want+:\n%s", diff)
				}
			})
			t.Run("book ids", func(t *testing.T) {
				got := library.SortedBookIDs()
				if diff := cmp.Diff(got, sortedBookIDs); diff != "" {
					t.Errorf("ERROR: got- want+:\n%s", diff)
				}
			})
		})
	}
}

func TestLibraryBooks(t *testing.T) {

	library := books.NewLibrary("")
	library.SetBooks(sortedBooks...)
	library.SetBooks(sortedBooks...)

	t.Run("number of books", func(t *testing.T) {
		got := len(library.Books)
		want := len(sortedBooks)
		if got != want {
			t.Errorf("ERROR: got %v, want %v", got, want)
		}
	})

	t.Run("get existing book", func(t *testing.T) {
		wantID := sortedBooks[2].ID
		got := library.Book(wantID)
		if diff := cmp.Diff(got, sortedBooks[2],
			cmpopts.IgnoreUnexported(books.Book{})); diff != "" {
			t.Errorf("ERROR: got- want+:\n%s", diff)
		}
	})

	t.Run("get new book", func(t *testing.T) {
		wantID := books.ID{Title: "new book"}
		wantBook := books.New(wantID)
		got := library.Book(wantID)
		if diff := cmp.Diff(got, wantBook,
			cmpopts.IgnoreUnexported(books.Book{})); diff != "" {
			t.Errorf("ERROR: got- want+:\n%s", diff)
		}
	})
}

func init() {
	for _, id := range sortedBookIDs {
		sortedBooks = append(sortedBooks, books.New(id))
	}
}

var sortedBooks []books.Book

var sortedBookIDs = []books.ID{
	// book title starts numeric
	{Title: "200 quick and easy phrases  for japanese conversation"},
	{Title: "88 basic patterns for japanese conversation"},
	// book title starts with character
	{Title: "first foreign japanese"},
	{Title: "how to speak osaka dialect"},
	// volume only in book title
	books.NewID("Grundstudium Japanisch 1", "Grundstudium Japanisch", 0),
	books.NewID("Grundstudium Japanisch 2", "Grundstudium Japanisch", 0),
	// volume in book title and volume
	// order is determined by volume, not title!
	books.NewID("minna no nihongo sho 1", "minna no nihongo", 1),
	books.NewID("minna no nihongo sho 2", "minna no nihongo", 2),
	books.NewID("minna no nihongo chuu 1", "minna no nihongo", 3),
	books.NewID("minna no nihongo chuu 2", "minna no nihongo", 4),
	// volume not in book title
	books.NewID("nihongo de doozo", "nihongo de doozo", 1),
	books.NewID("nihongo de doozo", "nihongo de doozo", 2),
	// title and series identical, no volume
	books.NewID("nihongo e yookoso", "nihongo e yookoso", 0),
}
