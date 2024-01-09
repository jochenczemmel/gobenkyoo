package learn_test

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// Input kanji cards and resulting learn cards for box test.

var kanjiCards = []kanjis.Card{
	kanjis.NewBuilder('方').
		AddDetailsWithKana("HOO", "ホー", "Richtung", "Art und Weise, etwas zu tun").
		AddDetailsWithKana("kata", "かた", "Person", "Art und Weise, etwas zu tun").
		Build(),
	kanjis.NewBuilder('曜').AddDetails("yoo", "weekday").Build(),
}

var wantKanji2Native = []learn.Card{{
	ID:          "方",
	Question:    "方",
	Answer:      "Richtung, Art und Weise, etwas zu tun, Person",
	MoreAnswers: []string{"HOO, kata", "ホー, かた"},
}, {
	ID:          "曜",
	Question:    "曜",
	Answer:      "weekday",
	MoreAnswers: []string{"yoo"},
}}

var wantNative2Kanji = []learn.Card{{
	ID:          "方",
	Question:    "Richtung, Art und Weise, etwas zu tun, Person",
	Answer:      "方",
	MoreAnswers: []string{"HOO, kata", "ホー, かた"},
}, {
	ID:          "曜",
	Question:    "weekday",
	Answer:      "曜",
	MoreAnswers: []string{"yoo"},
}}

var wantKana2Kanji = []learn.Card{{
	ID:       "方",
	Question: "ホー, かた",
	Answer:   "方",
	MoreAnswers: []string{
		"Richtung, Art und Weise, etwas zu tun, Person",
		"HOO, kata",
	},
}, {
	ID:          "曜",
	Question:    "yoo",
	Answer:      "曜",
	MoreAnswers: []string{"weekday"},
}}
