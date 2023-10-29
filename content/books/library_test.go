package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

var allBooks = []*books.Book{
	&books.Book{
		Title: "200 quick and easy phrases  for japanese conversation",
	},
	&books.Book{
		Title: "88 basic patterns for japanese conversation",
	},
	&books.Book{
		Title: "first foreign japanese",
	},
	&books.Book{
		Title: "how to speak osaka dialect",
	},
	&books.Book{
		Title:       "minna no nihongo sho 1",
		SeriesTitle: "minna no nihongo",
		Volume:      1,
	},
	&books.Book{
		Title:       "minna no nihongo sho 2",
		SeriesTitle: "minna no nihongo",
		Volume:      2,
	},
	&books.Book{
		Title:       "minna no nihongo chuu 1",
		SeriesTitle: "minna no nihongo",
		Volume:      3,
	},
	&books.Book{
		Title:       "minna no nihongo chuu 2",
		SeriesTitle: "minna no nihongo",
		Volume:      4,
	},
	&books.Book{
		Title:       "nihongo de doozo 1",
		SeriesTitle: "nihongo de doozo",
		Volume:      1,
	},
	&books.Book{
		Title:       "nihongo de doozo 2",
		SeriesTitle: "nihongo de doozo",
		Volume:      2,
	},
	&books.Book{
		Title:       "nihongo e yookoso",
		SeriesTitle: "nihongo e yookoso",
	},
}

var bookComparer = cmp.Comparer(func(x, y *books.Book) bool {
	if x.Title == y.Title &&
		x.SeriesTitle == y.SeriesTitle &&
		x.Volume == y.Volume {
		return true
	}
	return false
})

func TestLibrarySort(t *testing.T) {
	var lib books.Library

	got := lib.Content()
	if diff := cmp.Diff(got, []*books.Book{}, bookComparer); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}

	lib = books.NewLibrary(
		allBooks[2],
		allBooks[1],
		allBooks[0],
		allBooks[10],
		allBooks[5],
		allBooks[4],
		allBooks[3],
		allBooks[8],
		allBooks[7],
		allBooks[6],
		allBooks[9],
	)

	got = lib.Content()
	if diff := cmp.Diff(got, allBooks, bookComparer); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}

	wantTitles := []string{
		"200 quick and easy phrases  for japanese conversation",
		"88 basic patterns for japanese conversation",
		"first foreign japanese",
		"how to speak osaka dialect",
		"minna no nihongo sho 1",
		"minna no nihongo sho 2",
		"minna no nihongo chuu 1",
		"minna no nihongo chuu 2",
		"nihongo de doozo 1",
		"nihongo de doozo 2",
		"nihongo e yookoso",
	}

	gotTitles := lib.BookTitles()
	if diff := cmp.Diff(gotTitles, wantTitles); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}
}

func TestLibraryByTitle(t *testing.T) {
	lib := books.NewLibrary(allBooks...)

	bookTitle := "minna no nihongo sho 1"
	want := bookTitle
	if got := lib.ByTitle(bookTitle); got.Title != want {
		t.Errorf("ERROR: got %v, want %v", got.Title, want)
	}

	bookTitle = "not in lib"
	want = ""
	if got := lib.ByTitle(bookTitle); got.Title != want {
		t.Errorf("ERROR: got %v, want %v", got.Title, want)
	}
}

func TestLibraryBySeriesTitle(t *testing.T) {
	lib := books.NewLibrary(allBooks...)

	seriesTitle := "not in lib"
	want := []*books.Book{}
	got := lib.BySeriesTitle(seriesTitle)
	if diff := cmp.Diff(got, want, bookComparer); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}

	seriesTitle = "nihongo e yookoso"
	want = []*books.Book{
		&books.Book{
			Title:       "nihongo e yookoso",
			SeriesTitle: "nihongo e yookoso",
		},
	}
	got = lib.BySeriesTitle(seriesTitle)
	if diff := cmp.Diff(got, want, bookComparer); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}

	seriesTitle = "nihongo de doozo"
	want = []*books.Book{
		&books.Book{
			Title:       "nihongo de doozo 1",
			SeriesTitle: "nihongo de doozo",
			Volume:      1,
		},
		&books.Book{
			Title:       "nihongo de doozo 2",
			SeriesTitle: "nihongo de doozo",
			Volume:      2,
		},
	}
	got = lib.BySeriesTitle(seriesTitle)
	if diff := cmp.Diff(got, want, bookComparer); diff != "" {
		t.Errorf("ERROR: got-, want+\n%s", diff)
	}
}
