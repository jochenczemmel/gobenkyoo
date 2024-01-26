package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

func TestConvertContentCards(t *testing.T) {
	shelf := learn.NewClassroom("")
	boxID := learn.BoxID{}
	shelf.NewWordBox(boxID, wordCards...)
	shelf.NewKanjiBox(boxID, kanjiCards...)

	testCases := []struct {
		mode   string
		method func(learn.Options, ...learn.BoxID) learn.Exam
		want   []learn.Card
	}{
		{
			method: shelf.StartWordExam,
			mode:   learn.Native2Japanese,
			want:   wantNative2Japanese,
		},
		{
			method: shelf.StartWordExam,
			mode:   learn.Japanese2Native,
			want:   wantJapanese2Native,
		},
		{
			method: shelf.StartWordExam,
			mode:   learn.Native2Kana,
			want:   wantNative2Kana,
		},
		{
			method: shelf.StartWordExam,
			mode:   learn.Kana2Native,
			want:   wantKana2Native,
		},
		{
			method: shelf.StartKanjiExam,
			mode:   learn.Kanji2Native,
			want:   wantKanji2Native,
		},
		{
			method: shelf.StartKanjiExam,
			mode:   learn.Native2Kanji,
			want:   wantNative2Kanji,
		},
		{
			method: shelf.StartKanjiExam,
			mode:   learn.Kana2Kanji,
			want:   wantKana2Kanji,
		},
	}
	opt := learn.Options{
		Level:     learn.MinLevel,
		NoShuffle: true,
	}

	for _, c := range testCases {
		t.Run(c.mode, func(t *testing.T) {
			opt.LearnMode = c.mode
			got := c.method(opt, boxID).Cards()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})
	}
}

func TestConvertCardsIDs(t *testing.T) {
	shelf := learn.NewClassroom("")
	boxID := learn.BoxID{
		Name:     "box 1",
		LessonID: books.NewLessonID("lesson 1", "book 1", "book", 1),
	}
	shelf.NewWordBox(boxID, wordCards...)
	shelf.NewKanjiBox(boxID, kanjiCards...)
	opt := learn.Options{
		Level:     learn.MinLevel,
		NoShuffle: true,
	}

	t.Run("word cards", func(t *testing.T) {
		opt.LearnMode = learn.Native2Japanese
		got := shelf.StartWordExam(opt, boxID).Cards()
		if diff := cmp.Diff(got, wantNative2JapaneseWithLesson); diff != "" {
			t.Errorf("ERROR: -got +want\n%s", diff)
		}
	})

	t.Run("kanji cards", func(t *testing.T) {
		opt.LearnMode = learn.Kanji2Native
		got := shelf.StartKanjiExam(opt, boxID).Cards()
		if diff := cmp.Diff(got, wantKanji2NativeWithLesson); diff != "" {
			t.Errorf("ERROR: -got +want\n%s", diff)
		}
	})
}
