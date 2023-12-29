package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

type Card struct {
	Question    string
	Hint        string
	Answer      string
	MoreAnswers []string
	Explanation string
	WordCard    *words.Card
	KanjiCard   *kanjis.Card
}
