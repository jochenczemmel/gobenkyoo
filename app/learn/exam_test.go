package learn_test

import (
	"strconv"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordExam(t *testing.T) {

	inputCards := []words.Card{{
		Nihongo: "世界",
		Kana:    "せかい",
		Romaji:  "sekai",
		Meaning: "world",
	}, {
		Nihongo:  "見ます",
		Kana:     "みます",
		Romaji:   "mimasu",
		Meaning:  "to see",
		DictForm: "見る",
		TeForm:   "見て",
		NaiForm:  "見ない",
	}, {
		Nihongo:     "今日",
		Kana:        "きょう",
		Romaji:      "kyoo",
		Meaning:     "today",
		Hint:        "often written in kana",
		Explanation: "might also be honjitsu(本日)",
	}}

	box := learn.NewWordBox(inputCards...)

	options := learn.ExamOptions{
		LearnMode: learn.Native2Japanese,
		Level:     learn.MinLevel,
		NoShuffle: true,
	}
	exam := learn.NewExam(options, box)

	testCases := []struct {
		wantQuestion string
		wantOk       bool
	}{
		{wantQuestion: inputCards[0].Meaning, wantOk: true},
		{wantQuestion: inputCards[1].Meaning, wantOk: true},
		{wantQuestion: inputCards[2].Meaning, wantOk: true},
		{wantQuestion: inputCards[2].Meaning, wantOk: false},
	}

	for i, c := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			challenge, ok := exam.NextCard()
			if ok != c.wantOk {
				t.Fatalf("ERROR: card index: got %v, want %v", ok, c.wantOk)
			}
			if challenge.Question != c.wantQuestion {
				t.Errorf("ERROR: got %v, want %v", challenge.Question, c.wantQuestion)
			}

		})
	}
}
