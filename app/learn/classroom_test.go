package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func TestConvertContentCards(t *testing.T) {
	room := learn.NewClassroom("")
	boxID := learn.BoxID{Name: "1"}
	room.SetKanjiBoxes(learn.NewKanjiBox(boxID, kanjiCards...))
	room.SetWordBoxes(learn.NewWordBox(boxID, wordCards...))

	testCases := []struct {
		mode   string
		method func(learn.Options, ...learn.BoxID) learn.Exam
		want   []learn.Card
	}{{
		method: room.StartWordExam,
		mode:   learn.Native2Japanese,
		want:   wantNative2Japanese,
	}, {
		method: room.StartWordExam,
		mode:   learn.Japanese2Native,
		want:   wantJapanese2Native,
	}, {
		method: room.StartWordExam,
		mode:   learn.Native2Kana,
		want:   wantNative2Kana,
	}, {
		method: room.StartWordExam,
		mode:   learn.Kana2Native,
		want:   wantKana2Native,
	}, {
		method: room.StartKanjiExam,
		mode:   learn.Kanji2Native,
		want:   wantKanji2Native,
	}, {
		method: room.StartKanjiExam,
		mode:   learn.Native2Kanji,
		want:   wantNative2Kanji,
	}, {
		method: room.StartKanjiExam,
		mode:   learn.Kana2Kanji,
		want:   wantKana2Kanji,
	}}
	opt := learn.Options{
		Level:     learn.MinLevel,
		NoShuffle: true,
	}

	for _, c := range testCases {
		t.Run(c.mode, func(t *testing.T) {
			opt.LearnMode = c.mode
			got := c.method(opt, boxID).Cards()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

func TestClassroomGetBoxes(t *testing.T) {
	room := learn.NewClassroom("")
	kanjiBoxIds := []learn.BoxID{
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
	}
	wordBoxIds := []learn.BoxID{
		{Name: "4"},
		{Name: "2"},
	}
	kanjiBoxes := []learn.Box{}
	for _, id := range kanjiBoxIds {
		kanjiBoxes = append(kanjiBoxes, learn.NewKanjiBox(id))
	}
	wordBoxes := []learn.Box{}
	for _, id := range wordBoxIds {
		wordBoxes = append(wordBoxes, learn.NewWordBox(id))
	}

	room.SetKanjiBoxes(kanjiBoxes...)
	room.SetWordBoxes(wordBoxes...)
	room.SetKanjiBoxes(wordBoxes...)
	room.SetWordBoxes(kanjiBoxes...)

	t.Run("get kanjis", func(t *testing.T) {
		got := room.KanjiBoxes()
		if diff := cmp.Diff(got, kanjiBoxes,
			cmpopts.IgnoreUnexported(learn.Box{})); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})

	t.Run("get words", func(t *testing.T) {
		got := room.WordBoxes()
		if diff := cmp.Diff(got, wordBoxes,
			cmpopts.IgnoreUnexported(learn.Box{})); diff != "" {
			t.Errorf("ERROR: got- want+\n%s", diff)
		}
	})

	t.Run("get single box", func(t *testing.T) {
		testCases := []struct {
			name   string
			input  learn.BoxID
			method func(learn.BoxID) learn.Box
			want   learn.Box
		}{{
			name:   "kanji ok",
			input:  kanjiBoxIds[0],
			method: room.KanjiBox,
			want:   kanjiBoxes[0],
		}, {
			name:   "kanji not found",
			input:  wordBoxIds[0],
			method: room.KanjiBox,
			want:   learn.NewKanjiBox(learn.BoxID{}),
		}, {
			name:   "word ok",
			input:  wordBoxIds[0],
			method: room.WordBox,
			want:   wordBoxes[0],
		}, {
			name:   "word not found",
			input:  kanjiBoxIds[0],
			method: room.WordBox,
			want:   learn.NewWordBox(learn.BoxID{}),
		}}

		for _, c := range testCases {
			t.Run(c.name, func(t *testing.T) {
				got := c.method(c.input)
				if diff := cmp.Diff(got, c.want,
					cmpopts.IgnoreUnexported(learn.Box{})); diff != "" {
					t.Errorf("ERROR: got- want+\n%s", diff)
				}
			})
		}
	})
}
