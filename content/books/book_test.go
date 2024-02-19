package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var allNames = []string{"lesson 1", "lesson 2", "lesson 3"}

var wordCards = []words.Card{
	{ID: 1, Nihongo: "世界"},
	{ID: 2, Nihongo: "日本"},
	{ID: 3, Nihongo: "白鳳"},
	{ID: 4, Nihongo: "大相撲"},
	{ID: 5, Nihongo: "福岡"},
}

var kanjiCards = []kanjis.Card{
	{ID: 1, Kanji: '方'},
	{ID: 2, Kanji: '世'},
	{ID: 6, Kanji: '界'},
	{ID: 7, Kanji: '日'},
	{ID: 8, Kanji: '本'},
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

	testCases := []struct {
		name        string
		input, want string
		wantOK      bool
	}{
		{
			name:   "lesson in book",
			input:  allNames[1],
			want:   allNames[1],
			wantOK: true,
		},
		{
			name:   "lesson in book",
			input:  allNames[2],
			want:   "",
			wantOK: false,
		},
	}
	for _, c := range testCases {

		t.Run(c.name, func(t *testing.T) {
			got, ok := book.Lesson(c.input)
			if got.Name != c.want || ok != c.wantOK {
				t.Errorf("ERROR: got %q, %v, want %q, %v",
					got.Name, ok, c.want, c.wantOK)
			}
		})
	}
}
