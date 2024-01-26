package jsondb_test

import (
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var (
	testLibraryName      = "japanology"
	testBookTitle1       = "minna no nihongo sho 1"
	testBookSeriesTitle1 = "minna no nihongo"
	testBookVolume1      = 1
	testBook1            = books.New(books.NewID(
		testBookTitle1,
		testBookSeriesTitle1,
		testBookVolume1,
	))
	testBook2 = books.New(books.NewID(
		"minna no nihongo sho 2", "minna no nihongo", 2,
	))
	testLessonName1 = "lesson1"
	testLessonName2 = "lesson2"
)

var kanjiCardsLesson1 = []kanjis.Card{{
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
	Kanji: '曜',
	Details: []kanjis.Detail{
		{Reading: "yoo", Meanings: []string{"weekday"}},
	},
}}

var wordCardsLesson1 = []words.Card{{
	Nihongo: "日曜日",
	Kana:    "にちようび",
	Romaji:  "nichyoobi",
	Meaning: "sunday",
}, {
	Nihongo: "月曜日",
	Kana:    "げつようび",
	Romaji:  "getsuyoobi",
	Meaning: "monday",
}, {
	Nihongo: "火曜日",
	Kana:    "かようび",
	Romaji:  "kayoobi",
	Meaning: "tuesday",
}}

var kanjiCardsLesson2 = []kanjis.Card{
	{
		Kanji: '習',
		Details: []kanjis.Detail{{
			Reading:     "narai(u)",
			ReadingKana: "なら（う）",
			Meanings:    []string{"learn"},
		}},
	},
	{
		Kanji: '世',
		Details: []kanjis.Detail{{
			Reading:  "se",
			Meanings: []string{"world", "era", "generation"},
		}},
	},
	{
		Kanji: '界',
		Details: []kanjis.Detail{{
			Reading:  "kai",
			Meanings: []string{"all"},
		}},
	},
}

var wordCardsLesson2 = []words.Card{{
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
