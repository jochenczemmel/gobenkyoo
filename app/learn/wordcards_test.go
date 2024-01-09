package learn_test

// Input word cards and resulting learn cards for word box test.

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var wordCards = []words.Card{{
	Nihongo:     "習います",
	Kana:        "ならいます",
	Romaji:      "naraimasu",
	Meaning:     "to learn",
	DictForm:    "習う",
	TeForm:      "習って",
	NaiForm:     "習わない",
	Hint:        "from somebody",
	Explanation: "to study is benkyoo (勉強)",
}, {
	Nihongo: "世界",
	Kana:    "せかい",
	Romaji:  "sekai",
	Meaning: "world",
}}

var wantNative2Japanese = []learn.Card{{
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

var wantJapanese2Native = []learn.Card{{
	Question: "習います",
	Hint:     "from somebody",
	Answer:   "to learn",
	MoreAnswers: []string{
		"ならいます",
		"naraimasu",
		"習う",
		"習って",
		"習わない",
	},
	Explanation: "to study is benkyoo (勉強)",
}, {
	Question: "世界",
	Answer:   "world",
	MoreAnswers: []string{
		"せかい",
		"sekai",
	},
}}

var wantNative2Kana = []learn.Card{{
	Question: "to learn",
	Hint:     "from somebody",
	Answer:   "ならいます",
	MoreAnswers: []string{
		"naraimasu",
		"習います",
		"習う",
		"習って",
		"習わない",
	},
	Explanation: "to study is benkyoo (勉強)",
}, {
	Question: "world",
	Answer:   "せかい",
	MoreAnswers: []string{
		"sekai",
		"世界",
	},
}}

var wantKana2Native = []learn.Card{{
	Question: "ならいます",
	Hint:     "from somebody",
	Answer:   "to learn",
	MoreAnswers: []string{
		"naraimasu",
		"習います",
		"習う",
		"習って",
		"習わない",
	},
	Explanation: "to study is benkyoo (勉強)",
}, {
	Question: "せかい",
	Answer:   "world",
	MoreAnswers: []string{
		"sekai",
		"世界",
	},
}}
