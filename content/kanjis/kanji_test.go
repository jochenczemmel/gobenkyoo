package kanjis_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

func TestKanjiInfo(t *testing.T) {

	testCases := []struct {
		name                        string
		card                        *kanjis.Card
		wantRune                    rune
		wantString, wantDescriptor  string
		wantNumber, wantStrokeCount int
	}{
		{
			name:     "empty",
			card:     kanjis.NewBuilder(' ').Build(),
			wantRune: ' ',
		},
		{
			name: "kata_hoo",
			card: kanjis.
				NewBuilder('方').
				Build(),
			wantRune:        '方',
			wantString:      "方 (4h0.1/70)",
			wantDescriptor:  "4h0.1",
			wantNumber:      70,
			wantStrokeCount: 4,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			gotRune := c.card.Rune()
			if gotRune != c.wantRune {
				t.Errorf("ERROR: got %c, want %c", gotRune, c.wantRune)
			}
			got := c.card.String()
			if got != c.wantString {
				t.Errorf("ERROR: got %v, want %v", got, c.wantString)
			}
			got = c.card.Descriptor()
			if got != c.wantDescriptor {
				t.Errorf("ERROR: got %v, want %v", got, c.wantDescriptor)
			}
			gotNum := c.card.Number()
			if gotNum != c.wantNumber {
				t.Errorf("ERROR: got %v, want %v", gotNum, c.wantNumber)
			}
			gotNum = c.card.StrokeCount()
			if gotNum != c.wantStrokeCount {
				t.Errorf("ERROR: got %v, want %v", gotNum, c.wantStrokeCount)
			}
		})
	}
}

func TestKanjiDetails(t *testing.T) {

	testCases := []struct {
		name                               string
		card                               *kanjis.Card
		wantLen                            int
		wantReading, wantKana, wantMeaning []string
	}{
		{
			name:        "empty",
			card:        kanjis.NewBuilder(' ').Build(),
			wantReading: []string{},
			wantKana:    []string{},
			wantMeaning: []string{},
		},
		{
			name: "kata_hoo",
			card: kanjis.
				NewBuilder('方').
				AddDetailsKana("HOO", "ホー", "Richtung", "Art und Weise, etwas zu tun").
				AddDetails("kata", "Person", "Art und Weise, etwas zu tun").
				Build(),
			wantLen:     2,
			wantReading: []string{"HOO", "kata"},
			wantKana:    []string{"ホー", "かた"},
			wantMeaning: []string{"Richtung", "Art und Weise, etwas zu tun", "Person"},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			gotNum := len(c.card.Details())
			if gotNum != c.wantLen {
				t.Errorf("ERROR: got %v, want %v", gotNum, c.wantLen)
			}
			if diff := cmp.Diff(c.card.Readings(), c.wantReading); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
			if diff := cmp.Diff(c.card.ReadingsKana(), c.wantKana); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
			if diff := cmp.Diff(c.card.Meanings(), c.wantMeaning); diff != "" {
				t.Errorf("ERROR: got-, want+\n%s", diff)
			}
		})
	}
}

func TestKanjiHasRadical(t *testing.T) {

	candidates := []struct {
		name    string
		kanji   rune
		radical rune
		want    bool
	}{
		{"radical", '山', '山', true},
		{"kanji with radical", '島', '山', true},
		{"kanji without radical", '元', '山', false},
		{"hiragana", 'ん', '山', false},
		{"romaji", 'x', '山', false},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			got := kanjis.NewBuilder(c.kanji).Build().HasRadical(c.radical)
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

func TestKanjiRadicals(t *testing.T) {

	candidates := []struct {
		name  string
		kanji rune
		want  string
	}{
		{"radical", '山', "山"},
		{"multiple radicals", '島', "山鳥白"},
		{"hiragana", 'は', ""},
		{"romaji", 'x', ""},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			got := kanjis.NewBuilder(c.kanji).Build().Radicals()
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}
