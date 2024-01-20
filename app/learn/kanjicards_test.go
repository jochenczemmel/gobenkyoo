package learn_test

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// Input kanji cards and resulting learn cards for box test.

var kanjiCards = []kanjis.Card{{
	ID:    1,
	Kanji: '方',
	Details: []kanjis.Detail{{
		Reading:     "HOO",
		ReadingKana: "ホー",
		Meanings:    []string{"Richtung", "Art und Weise, etwas zu tun"},
	}, {
		Reading:     "kata",
		ReadingKana: "かた",
		Meanings:    []string{"Person", "Art und Weise, etwas zu tun"},
	}},
}, {
	ID:    2,
	Kanji: '曜',
	Details: []kanjis.Detail{{
		Reading:  "yoo",
		Meanings: []string{"weekday"},
	}},
}}

var wantKanji2Native = []learn.Card{{
	ID:          learn.CardID{ContentID: 1},
	Question:    "方",
	Answer:      "Richtung, Art und Weise, etwas zu tun, Person",
	MoreAnswers: []string{"HOO, kata", "ホー, かた"},
}, {
	ID:          learn.CardID{ContentID: 2},
	Question:    "曜",
	Answer:      "weekday",
	MoreAnswers: []string{"yoo"},
}}

var wantKanji2NativeWithLesson = []learn.Card{{
	ID: learn.CardID{
		ContentID: 1,
		LessonID: learn.LessonID{
			Title:       "lesson 1",
			BookTitle:   "book 1",
			SeriesTitle: "book",
			Volume:      1,
		},
	},
	Question:    "方",
	Answer:      "Richtung, Art und Weise, etwas zu tun, Person",
	MoreAnswers: []string{"HOO, kata", "ホー, かた"},
}, {
	ID: learn.CardID{
		ContentID: 2,
		LessonID: learn.LessonID{
			Title:       "lesson 1",
			BookTitle:   "book 1",
			SeriesTitle: "book",
			Volume:      1,
		},
	},
	Question:    "曜",
	Answer:      "weekday",
	MoreAnswers: []string{"yoo"},
}}

var wantNative2Kanji = []learn.Card{{
	ID:          learn.CardID{ContentID: 1},
	Question:    "Richtung, Art und Weise, etwas zu tun, Person",
	Answer:      "方",
	MoreAnswers: []string{"HOO, kata", "ホー, かた"},
}, {
	ID:          learn.CardID{ContentID: 2},
	Question:    "weekday",
	Answer:      "曜",
	MoreAnswers: []string{"yoo"},
}}

var wantKana2Kanji = []learn.Card{{
	ID:       learn.CardID{ContentID: 1},
	Question: "ホー, かた",
	Answer:   "方",
	MoreAnswers: []string{
		"Richtung, Art und Weise, etwas zu tun, Person",
		"HOO, kata",
	},
}, {
	ID:          learn.CardID{ContentID: 2},
	Question:    "yoo",
	Answer:      "曜",
	MoreAnswers: []string{"weekday"},
}}
