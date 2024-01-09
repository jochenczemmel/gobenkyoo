package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func TestConvertContentCards(t *testing.T) {
	shelf := learn.NewLibrary()
	boxTitle := learn.BoxName{BoxTitle: "lesson 1"}
	shelf.AddWordBox(boxTitle, wordCards...)
	shelf.AddKanjiBox(boxTitle, kanjiCards...)

	testCases := []struct {
		mode   string
		method func(learn.Options, ...learn.BoxName) learn.Exam
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
			got := c.method(opt, boxTitle).Cards()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: -got +want\n%s", diff)
			}
		})
	}
}
