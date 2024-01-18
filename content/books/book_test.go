package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestBookLessons(t *testing.T) {
	book := books.New("", "", 0)
	for _, lesson := range []string{"l1", "l2"} {
		book.AddKanjis(lesson)
	}
	for _, lesson := range []string{"l2", "l3"} {
		book.AddWords(lesson)
	}
	want := []string{"l1", "l2", "l3"}
	got := book.Lessons()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Lessons(): -got +want\n%s", diff)
	}
}

func TestBookAdd(t *testing.T) {
	title := "minna1"
	seriesTitle := "minna"
	volume := 1
	book := books.New(title, seriesTitle, volume)

	lesson := "lesson 1"
	missingLesson := "not existing"

	t.Run("add kanjis", func(t *testing.T) {
		cards := []kanjis.Card{
			kanjis.Card{Kanji: '方'},
			kanjis.Card{Kanji: '世'},
			kanjis.Card{Kanji: '界'},
			kanjis.Card{Kanji: '日'},
			kanjis.Card{Kanji: '本'},
		}

		book.AddKanjis(lesson, cards[:3]...)
		got := book.KanjisFor(lesson)
		if diff := cmp.Diff(got, cards[:3]); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}

		book.AddKanjis(lesson, cards[3:]...)
		got = book.KanjisFor(lesson)
		if diff := cmp.Diff(got, cards); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}

		got = book.KanjisFor(missingLesson)
		if diff := cmp.Diff(got, cards[:0]); diff != "" {
			t.Errorf("KanjisFor(%v): -got +want\n%s", lesson, diff)
		}
	})

	t.Run("add words", func(t *testing.T) {
		cards := []words.Card{
			{Nihongo: "世界"},
			{Nihongo: "日本"},
			{Nihongo: "白鳳"},
			{Nihongo: "大相撲"},
			{Nihongo: "福岡"},
		}

		book.AddWords(lesson, cards[:3]...)
		got := book.WordsFor(lesson)
		if diff := cmp.Diff(got, cards[:3]); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}

		book.AddWords(lesson, cards[3:]...)
		got = book.WordsFor(lesson)
		if diff := cmp.Diff(got, cards); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}

		got = book.WordsFor(missingLesson)
		if diff := cmp.Diff(got, cards[:0]); diff != "" {
			t.Errorf("WordsFor(%v): -got +want\n%s", lesson, diff)
		}
	})
}
