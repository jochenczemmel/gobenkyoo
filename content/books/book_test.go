package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var allNames = []string{"lesson 1", "lesson 2", "lesson 3"}

var wordCards = []words.Card{
	{ID: "1", Nihongo: "世界"},
	{ID: "2", Nihongo: "日本"},
	{ID: "3", Nihongo: "白鳳"},
	{ID: "4", Nihongo: "大相撲"},
	{ID: "5", Nihongo: "福岡"},
}

var kanjiCards = []kanjis.Card{
	{ID: "1", Kanji: '方'},
	{ID: "2", Kanji: '世'},
	{ID: "3", Kanji: '界'},
	{ID: "4", Kanji: '日'},
	{ID: "5", Kanji: '本'},
}

func TestBookSetLesson(t *testing.T) {
	wantNames := allNames[:2]
	book := books.New(books.ID{})
	lesson1 := books.NewLesson(wantNames[0])
	lesson2 := books.NewLesson(wantNames[1])
	book.SetLessons(lesson1, lesson2)

	t.Run("LessonNames", func(t *testing.T) {
		got := book.LessonNames()
		if diff := cmp.Diff(got, wantNames); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})

	t.Run("Lessons", func(t *testing.T) {
		got := book.Lessons()
		want := []books.Lesson{lesson1, lesson2}
		if diff := cmp.Diff(got, want,
			cmpopts.IgnoreUnexported(books.Lesson{})); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})

	testCases := []struct {
		name        string
		input, want string
	}{
		{
			name:  "lesson in book",
			input: allNames[1],
			want:  allNames[1],
		},
		{
			name:  "lesson not in book",
			input: allNames[2],
			want:  allNames[2],
		},
	}
	for _, c := range testCases {

		t.Run(c.name, func(t *testing.T) {
			got := book.Lesson(c.input)
			if got.Name != c.want {
				t.Errorf("ERROR: got %q, want %q", got.Name, c.want)
			}
		})
	}
}
