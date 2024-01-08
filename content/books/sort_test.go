package books_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// TestSort tests the sort function for a list of pointers to book.
func TestSort(t *testing.T) {

	// prepare test: sorted list of books
	sortedBooks := []books.Book{
		// book title starts numeric
		books.New("200 quick and easy phrases  for japanese conversation", "", 0),
		books.New("88 basic patterns for japanese conversation", "", 0),
		// book title starts with character
		books.New("first foreign japanese", "", 0),
		books.New("how to speak osaka dialect", "", 0),
		// volume only in book title
		books.New("Grundstudium Japanisch 1", "Grundstudium Japanisch", 0),
		books.New("Grundstudium Japanisch 2", "Grundstudium Japanisch", 0),
		// volume in book title and volume
		// order is determined by volume, not title!
		books.New("minna no nihongo sho 1", "minna no nihongo", 1),
		books.New("minna no nihongo sho 2", "minna no nihongo", 2),
		books.New("minna no nihongo chuu 1", "minna no nihongo", 3),
		books.New("minna no nihongo chuu 2", "minna no nihongo", 4),
		// volume not in book title
		books.New("nihongo de doozo", "nihongo de doozo", 1),
		books.New("nihongo de doozo", "nihongo de doozo", 2),
		// title and series identical, no volume
		books.New("nihongo e yookoso", "nihongo e yookoso", 0),
	}

	// prepare test: copy book list, shuffle for test
	shuffledBooks := make([]books.Book, len(sortedBooks))
	copy(shuffledBooks, sortedBooks)
	rand.Shuffle(len(shuffledBooks), func(i, j int) {
		shuffledBooks[i], shuffledBooks[j] = shuffledBooks[j], shuffledBooks[i]
	})
	t.Logf("DEBUG: shuffled: first book: %v", shuffledBooks[0])

	// execute test
	sort.Sort(books.BySeriesVolumeTitle(shuffledBooks))
	t.Logf("DEBUG: sorted: first book: %v", shuffledBooks[0])

	if diff := cmp.Diff(shuffledBooks, sortedBooks); diff != "" {
		t.Errorf("ERROR: Sort: -got +want:\n%s", diff)
	}
}
