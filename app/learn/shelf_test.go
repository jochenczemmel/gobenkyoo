package learn_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordBox(t *testing.T) {
	shelf := learn.NewShelf()
	title := learn.BoxName{
		BoxTitle:  "lesson 1",
		BookTitle: books.TitleInfo{Title: "book 1"},
	}

	cards := []words.Card{
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

	shelf.AddWordBox(title, cards...)
}
