package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/app/learn/learncards"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordBox(t *testing.T) {
	shelf := learn.NewShelf()
	boxTitle := learn.BoxName{
		BoxTitle:  "lesson 1",
		BookTitle: books.TitleInfo{Title: "book 1"},
	}

	shelf.AddWordBox(boxTitle, wordCards...)

	opt := learncards.Options{
		LearnMode: learn.Native2Japanese,
		Level:     learncards.MinLevel,
		NoShuffle: true,
	}

	got := shelf.StartWordExam(opt, boxTitle).Cards()
	if diff := cmp.Diff(got, wantNative2Japanese); diff != "" {
		t.Errorf("ERROR: -got +want\n%s", diff)
	}
}

var wordCards = []words.Card{
	{
		Nihongo:     "習います",
		Kana:        "ならいます",
		Romaji:      "naraimasu",
		Meaning:     "to learn",
		DictForm:    "習う",
		TeForm:      "習って",
		NaiForm:     "習わない",
		Hint:        "from somebody",
		Explanation: "to study is benkyoo (勉強)",
	},
	{
		Nihongo: "世界",
		Kana:    "せかい",
		Romaji:  "sekai",
		Meaning: "world",
	},
}
var wantNative2Japanese = []learncards.Card{{
	Question: "to learn",
	Hint:     "from somebody",
	Answer:   "習います",
	MoreAnswers: []string{
		"ならいます",
		"naraimasu",
		"習う",
		"習って",
		"習わない",
	},
	Explanation: "to study is benkyoo (勉強)",
}, {
	Question: "world",
	Answer:   "世界",
	MoreAnswers: []string{
		"せかい",
		"sekai",
	},
}}
