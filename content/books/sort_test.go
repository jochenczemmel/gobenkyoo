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
	sortedBooks := []*books.Book{{
		// no volume, title starting with number
		Title: "200 quick and easy phrases  for japanese conversation",
	}, {
		Title: "88 basic patterns for japanese conversation",
	}, {
		// no volume, title starting with character
		Title: "first foreign japanese",
	}, {
		Title: "how to speak osaka dialect",
	}, {
		// volume only in book title
		Title:       "Grundstudium Japanisch 1",
		SeriesTitle: "Grundstudium Japanisch",
	}, {
		Title:       "Grundstudium Japanisch 2",
		SeriesTitle: "Grundstudium Japanisch",
	}, {
		// volume in book title and volume
		Title:       "minna no nihongo sho 1",
		SeriesTitle: "minna no nihongo",
		Volume:      1,
	}, {
		Title:       "minna no nihongo sho 2",
		SeriesTitle: "minna no nihongo",
		Volume:      2,
	}, {
		Title:       "minna no nihongo chuu 1",
		SeriesTitle: "minna no nihongo",
		Volume:      3,
	}, {
		Title:       "minna no nihongo chuu 2",
		SeriesTitle: "minna no nihongo",
		Volume:      4,
	}, {
		// volume not in book title
		Title:       "nihongo de doozo",
		SeriesTitle: "nihongo de doozo",
		Volume:      1,
	}, {
		Title:       "nihongo de doozo",
		SeriesTitle: "nihongo de doozo",
		Volume:      2,
	}, {
		Title:       "nihongo e yookoso",
		SeriesTitle: "nihongo e yookoso",
	}}

	// prepare test: copy book list, shuffle for test
	shuffledBooks := make([]*books.Book, len(sortedBooks))
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
