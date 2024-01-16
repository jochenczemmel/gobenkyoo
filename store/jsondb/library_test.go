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
	testBook1            = books.New(
		testBookTitle1,
		testBookSeriesTitle1,
		testBookVolume1,
	)
	testBook2        = books.New("minna no nihongo sho 2", "minna no nihongo", 2)
	testLessonTitle1 = "lesson1"
	testLessonTitle2 = "lesson2"
)

var kanjiCardsLesson1 = []kanjis.Card{
	kanjis.NewBuilder('方').
		AddDetailsWithKana("HOO", "ホー", "Richtung", "Art und Weise, etwas zu tun").
		AddDetailsWithKana("kata", "かた", "Person", "Art und Weise, etwas zu tun").
		Build(),
	kanjis.NewBuilder('曜').AddDetails("yoo", "weekday").Build(),
}

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
	kanjis.NewBuilder('習').
		AddDetailsWithKana("narai(u)", "なら（う）", "learn").Build(),
	kanjis.NewBuilder('世').AddDetails("se", "world", "era", "generation").Build(),
	kanjis.NewBuilder('界').AddDetails("kai", "all").Build(),
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
