package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

var minnaLessons = []*books.Lesson{
	{Title: "minnalesson1", BookTitle: "minna no nihongo sho 1"},
	{Title: "minnalesson2", BookTitle: "minna no nihongo sho 1"},
	{Title: "minnalesson3", BookTitle: "minna no nihongo sho 1"},
	{Title: "minnalesson4", BookTitle: "minna no nihongo sho 2"},
	{Title: "minnalesson5", BookTitle: "minna no nihongo sho 2"},
	{Title: "minnalesson6", BookTitle: "minna no nihongo sho 2"},
	{Title: "minnalesson7", BookTitle: "minna no nihongo sho 2"},
	{Title: "minnalesson8", BookTitle: "minna no nihongo chuu 1"},
	{Title: "minnalesson9", BookTitle: "minna no nihongo chuu 1"},
}

func makeLibrary() books.Library {
	yookosoLessons := []*books.Lesson{
		{Title: "yookosolesson1", BookTitle: "nihongo e yookoso"},
		{Title: "yookosolesson2", BookTitle: "nihongo e yookoso"},
		{Title: "yookosolesson3", BookTitle: "nihongo e yookoso"},
		{Title: "yookosolesson4", BookTitle: "nihongo e yookoso"},
		{Title: "yookosolesson5", BookTitle: "nihongo e yookoso"},
	}
	osakaLessons := []*books.Lesson{
		{Title: "firstlesson1", BookTitle: "first foreign japanese"},
		{Title: "firstlesson2", BookTitle: "first foreign japanese"},
		{Title: "firstlesson3", BookTitle: "first foreign japanese"},
		{Title: "firstlesson4", BookTitle: "first foreign japanese"},
	}

	book1 := books.New("minna no nihongo sho 1", "minna no nihongo", 1,
		minnaLessons[:3]...)
	book2 := books.New("minna no nihongo sho 2", "minna no nihongo", 2,
		minnaLessons[3:7]...)
	book3 := books.New("minna no nihongo chuu 1", "minna no nihongo", 3,
		minnaLessons[8:]...)
	book4 := books.New("nihongo e yookoso", "nihongo e yookoso", 0,
		yookosoLessons...)
	book5 := books.New("first foreign japanese", "", 0, osakaLessons...)

	return books.NewLibrary(&book1, &book2, &book3, &book4, &book5)
}

func TestAccLibraryLessonsUntil(t *testing.T) {
	wantEmpty := []*books.Lesson{}

	t.Run("uninitialized library", func(t *testing.T) {
		var lib books.Library
		got := lib.LessonsUntil("minna no nihongo sho 1", "minnalesson2")
		if diff := cmp.Diff(got, wantEmpty, lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})

	t.Run("empty library", func(t *testing.T) {
		lib := books.NewLibrary()
		got := lib.LessonsUntil("minna no nihongo sho 1", "minnalesson2")
		if diff := cmp.Diff(got, wantEmpty, lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})

	lib := makeLibrary()

	t.Run("two lessons", func(t *testing.T) {
		got := lib.LessonsUntil("minna no nihongo sho 1", "minnalesson2")
		if diff := cmp.Diff(got, minnaLessons[:2], lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})

	t.Run("lessons in two books", func(t *testing.T) {
		got := lib.LessonsUntil("minna no nihongo sho 2", "minnalesson6")
		if diff := cmp.Diff(got, minnaLessons[:6], lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})

	t.Run("lessons not in books", func(t *testing.T) {
		got := lib.LessonsUntil("minna no nihongo sho 1", "not found")
		if diff := cmp.Diff(got, wantEmpty, lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})

	t.Run("book not in library", func(t *testing.T) {
		got := lib.LessonsUntil("not found", "minnalesson2")
		if diff := cmp.Diff(got, wantEmpty, lessonComparer); diff != "" {
			t.Errorf("ERROR: got-, want+\n%s", diff)
		}
	})
}

var lessonComparer = cmp.Comparer(func(x, y *books.Lesson) bool {
	if x.Title == y.Title &&
		x.BookTitle == y.BookTitle &&
		len(x.Content()) == len(y.Content()) {
		return true
	}

	return false
})
