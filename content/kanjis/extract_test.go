package kanjis_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

func TestExtract(t *testing.T) {
	testCases := []struct {
		name     string
		in, want string
	}{{
		name: "only kanji",
		in:   "日本語",
		want: "日本語",
	}, {
		name: "kanji and hiragana",
		in:   "わたし は 日本語が むずかしい と 思います",
		want: "日本語思",
	}, {
		name: "duplicate kanjis",
		in:   "日本語は 毎日",
		want: "日本語毎",
	}, {
		name: "only kana",
		in:   "むずかしい ですね",
		want: "",
	}, {
		name: "not japanese",
		in:   "not japanese",
		want: "",
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := kanjis.Extract(c.in)
			if string(got) != c.want {
				t.Errorf("ERROR: got %s, want %s", string(got), c.want)
			}
		})
	}
}

func TestKana(t *testing.T) {

	candidates := []struct {
		name     string
		in, want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "hiragana", in: "hiragana", want: "ひらがな"},
		{name: "KATAKANA", in: "KATAKANA", want: "カタカナ"},
		{name: "kanji", in: "日本語", want: "日本語"},
		{name: "hiragana long a", in: "raamen", want: "らあめん"},
		{name: "KATAKANA KYUU", in: "KYUURI", want: "キューリ"},
		{name: "not japanese", in: "murks", want: "むrks"},
		{name: "hiragana long o with u", in: "koufukuji", want: "こうふくじ"},
		// not correct: use other package or fix package kana
		{name: "hiragana kyuu", in: "kyuuri", want: "きゅーり"},    // - only used in katakana
		{name: "hiragana oo", in: "koofukuji", want: "こおふくじ"},  // "こうふくじ"}
		{name: "hiragana ō", in: "kō", want: "こお"},             // should be "こう"
		{name: "KATAKANA OO", in: "KOO", want: "コオ"},           // should be "コー"
		{name: "KATAKANA Ō", in: "KŌ", want: "コオ"},             // should be "コー"
		{name: "KATAKANA AA", in: "RAAMEN", want: "ラアメン"},      // should be "ラーメン"
		{name: "KATAKANA EE", in: "EREBEETAA", want: "エレベイタア"}, // should be "エレベーター"}
		{name: "KATAKANA II", in: "PAATII", want: "パアティイ"},     // should be "パーティー"}
	}

	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			got := kanjis.ToKana(c.in)
			if got != c.want {
				t.Errorf("ERROR: got %q, want %q", got, c.want)
			}
		})
	}
}
