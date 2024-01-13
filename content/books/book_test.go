package books_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

func TestBookAddKanjis(t *testing.T) {
	title := "minna1"
	seriesTitle := "minna"
	volume := 1
	book := books.New(title, seriesTitle, volume)
	lesson := "lesson 1"
	kanjiCards := []kanjis.Card{
		kanjis.NewBuilder('方').Build(),
		kanjis.NewBuilder('世').Build(),
		kanjis.NewBuilder('界').Build(),
	}
	book.AddKanjis(lesson, kanjiCards...)
	got := book.KanjiFor(lesson)
	if diff := cmp.Diff(got, kanjiCards,
		cmp.Comparer(kanjiEqual)); diff != "" {
		t.Errorf("KanjiFor(%v): -got +want\n%s", lesson, diff)
	}
}

func kanjiEqual(got, want kanjis.Card) bool {
	return got.Kanji() == want.Kanji()
}
