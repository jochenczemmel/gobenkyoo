package books_test

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestLibraryFindCard(t *testing.T) {

	bookID := books.ID{
		Title:       "minna no nihongo sho 1",
		SeriesTitle: "minna no nihongo",
		Volume:      1,
	}
	book := books.New(bookID)

	lessonName := "lesson 1"
	lessonID := books.NewLessonID(lessonName, bookID.Title,
		bookID.SeriesTitle, bookID.Volume)
	lesson := books.NewLesson(lessonName)

	lesson.AddWords(wordCards...)
	lesson.AddKanjis(kanjiCards...)
	book.SetLessons(lesson)
	library := books.NewLibrary("")
	library.AddBooks(book)

	testCases := []struct {
		name          string
		lessonID      books.LessonID
		id            string
		wantKanjiCard kanjis.Card
		wantWordCard  words.Card
	}{{
		name:          "book in library, word and kanji",
		id:            "1",
		lessonID:      lessonID,
		wantWordCard:  wordCards[0],
		wantKanjiCard: kanjiCards[0],
	}, {
		name:          "book in library, word and kanji",
		id:            "1",
		lessonID:      lessonID,
		wantWordCard:  wordCards[0],
		wantKanjiCard: kanjiCards[0],
	}, {
		name:         "book in library, only word",
		id:           "5",
		lessonID:     lessonID,
		wantWordCard: wordCards[4],
	}, {
		name:          "book in library, only kanji",
		id:            "8",
		lessonID:      lessonID,
		wantKanjiCard: kanjiCards[4],
	}, {
		name:     "book in library, no match",
		id:       "42",
		lessonID: lessonID,
	}, {
		name: "book in library, wrong lesson",
		id:   "1",
		lessonID: books.LessonID{
			Name: "wrong lesson",
			ID:   bookID,
		},
	}, {
		name: "book not in library",
		id:   "1",
		lessonID: books.LessonID{
			Name: lessonName,
			ID:   books.ID{Title: "not in library"},
		},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("word", func(t *testing.T) {
				got := library.GetWordCard(c.lessonID, c.id)
				if diff := cmp.Diff(got, c.wantWordCard); diff != "" {
					t.Errorf("FindWordCard(%v): -got, +want\n%s", c.id, diff)
				}
			})
			t.Run("kanji", func(t *testing.T) {
				got := library.GetKanjiCard(c.lessonID, c.id)
				if diff := cmp.Diff(got, c.wantKanjiCard); diff != "" {
					t.Errorf("FindKanjiCard(%v): -got, +want\n%s", c.id, diff)
				}
			})
		})
	}
}

func TestLibrarySort(t *testing.T) {

	// prepare test: sorted list of books
	sortedBooks := []books.Book{
		// book title starts numeric
		books.New(books.ID{Title: "200 quick and easy phrases  for japanese conversation"}),
		books.New(books.ID{Title: "88 basic patterns for japanese conversation"}),
		// book title starts with character
		books.New(books.ID{Title: "first foreign japanese"}),
		books.New(books.ID{Title: "how to speak osaka dialect"}),
		// volume only in book title
		books.New(books.NewID("Grundstudium Japanisch 1", "Grundstudium Japanisch", 0)),
		books.New(books.NewID("Grundstudium Japanisch 2", "Grundstudium Japanisch", 0)),
		// volume in book title and volume
		// order is determined by volume, not title!
		books.New(books.NewID("minna no nihongo sho 1", "minna no nihongo", 1)),
		books.New(books.NewID("minna no nihongo sho 2", "minna no nihongo", 2)),
		books.New(books.NewID("minna no nihongo chuu 1", "minna no nihongo", 3)),
		books.New(books.NewID("minna no nihongo chuu 2", "minna no nihongo", 4)),
		// volume not in book title
		books.New(books.NewID("nihongo de doozo", "nihongo de doozo", 1)),
		books.New(books.NewID("nihongo de doozo", "nihongo de doozo", 2)),
		// title and series identical, no volume
		books.New(books.NewID("nihongo e yookoso", "nihongo e yookoso", 0)),
	}

	// test multiple times, shuffle may accidentially return the same order
	for i := 0; i < 3; i++ {
		t.Run("shuffle "+strconv.Itoa(i+1), func(t *testing.T) {

			shuffledBooks := slices.Clone(sortedBooks)
			rand.Shuffle(len(shuffledBooks), func(i, j int) {
				shuffledBooks[i], shuffledBooks[j] = shuffledBooks[j], shuffledBooks[i]
			})
			t.Logf("DEBUG: shuffled: first book: %v", shuffledBooks[0])

			library := books.NewLibrary("")
			library.AddBooks(shuffledBooks...)
			got := library.SortedBooks()
			t.Logf("DEBUG: sorted: first book: %v", got[0])

			if diff := cmp.Diff(got, sortedBooks, cmpopts.IgnoreUnexported(books.Book{})); diff != "" {
				t.Errorf("ERROR: Sort: -got +want:\n%s", diff)
			}
		})
	}
}
