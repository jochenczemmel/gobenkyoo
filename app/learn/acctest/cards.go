package acctest

import "github.com/jochenczemmel/gobenkyoo/content/words"

var WordCards = []*words.Card{{
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
