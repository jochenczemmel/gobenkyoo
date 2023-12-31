package kanjis_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

func TestBuilderThree(t *testing.T) {

	kanjiRune := '人'

	testCases := []struct {
		romaji, kana string
		meaning      []string
	}{{
		romaji:  "JIN",
		kana:    "ジン",
		meaning: []string{"Mensch", "Person", "Leute"},
	}, {
		romaji:  "NIN",
		kana:    "ニン",
		meaning: []string{"Mensch"},
	}, {
		romaji:  "hito",
		kana:    "ひと",
		meaning: []string{"Mensch", "Person", "Leute"},
	}}

	kb := kanjis.NewBuilder(kanjiRune)

	for _, c := range testCases {
		kb.AddDetails(c.romaji, c.meaning...)
	}
	kanji := kb.Build()

	details := kanji.Details()
	wantLen := 3
	if len(details) != wantLen {
		t.Fatalf("ERROR: got %v, want %v", len(details), wantLen)
	}

	for i, c := range testCases {
		if details[i].Reading() != c.romaji {
			t.Errorf("ERROR: got %v, want %v",
				details[i].Reading(), c.romaji)
		}
		if details[i].ReadingKana() != c.kana {
			t.Errorf("ERROR: got %v, want %v",
				details[i].ReadingKana(), c.kana)
		}
		if diff := cmp.Diff(details[i].Meanings(), c.meaning); diff != "" {
			t.Errorf("ERROR: got-, want+:\n%s\n", diff)
		}
	}

	gotPretty := details[2].String()
	wantPretty := "hito (ひと): Mensch, Person, Leute"
	if gotPretty != wantPretty {
		t.Errorf("ERROR: got %v, want %v", gotPretty, wantPretty)
	}
}

func TestBuilderNoDuplicated(t *testing.T) {

	kanjiRune := '人'

	testCases := []struct {
		romaji, kana string
		meaning      []string
	}{{
		romaji:  "JIN",
		kana:    "ジン",
		meaning: []string{"Mensch", "Person", "Leute"},
	}, {
		romaji:  "NIN",
		kana:    "ニン",
		meaning: []string{"Mensch"},
	}, {
		romaji:  "hito",
		kana:    "ひと",
		meaning: []string{"Mensch", "Person", "Leute"},
	}, {
		romaji:  "NIN",
		kana:    "ニン",
		meaning: []string{"Mensch"},
	}, {
		romaji:  "hito",
		kana:    "ひと",
		meaning: []string{"Mensch", "Person", "Leute"},
	}}

	kb := kanjis.NewBuilder(kanjiRune)

	for _, c := range testCases {
		kb.AddDetails(c.romaji, c.meaning...)
	}
	kanji := kb.Build()

	details := kanji.Details()
	wantLen := 3
	if len(details) != wantLen {
		t.Fatalf("ERROR: got %v, want %v", len(details), wantLen)
	}
}

func TestBuilderNoEmpty(t *testing.T) {

	kanjiRune := '人'

	testCases := []struct {
		romaji  string
		meaning []string
	}{{
		romaji:  "",
		meaning: []string{"Mensch", "Person", "Leute"},
	}, {
		romaji:  "NIN",
		meaning: []string{"Mensch"},
	}, {
		romaji:  "NIN",
		meaning: nil,
	}, {
		romaji:  "NIN",
		meaning: []string{},
	}}

	kb := kanjis.NewBuilder(kanjiRune)

	for _, c := range testCases {
		kb.AddDetails(c.romaji, c.meaning...)
	}
	kanji := kb.Build()

	details := kanji.Details()
	wantLen := 1
	if len(details) != wantLen {
		t.Fatalf("ERROR: got %v, want %v", len(details), wantLen)
	}
}

func TestBuilderReverted(t *testing.T) {

	kanjiRune := '人'

	testCases := []struct {
		romaji, kana string
		meaning      []string
	}{{
		romaji:  "JIN",
		kana:    "ジン",
		meaning: []string{"Mensch", "Person", "Leute"},
	}, {
		romaji:  "NIN",
		kana:    "ニン",
		meaning: []string{"Mensch"},
	}, {
		romaji:  "hito",
		kana:    "ひと",
		meaning: []string{"Mensch", "Person", "Leute"},
	}}

	kb := kanjis.NewBuilder(kanjiRune)

	for _, c := range testCases {
		kb.AddDetailsWithKana(c.romaji, c.kana, c.meaning...)
	}
	kanji := kb.Build()

	details := kanji.Details()
	wantLen := 3
	if len(details) != wantLen {
		t.Fatalf("ERROR: got %v, want %v", len(details), wantLen)
	}

	for i, c := range testCases {
		if details[i].ReadingKana() != c.kana {
			t.Errorf("ERROR: got %v, want %v",
				details[i].ReadingKana(), c.kana)
		}
	}
}
