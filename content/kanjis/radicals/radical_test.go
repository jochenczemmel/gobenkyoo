package radicals_test

import (
	"fmt"
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals"
)

func TestStrokeCount(t *testing.T) {
	testCases := []struct {
		name    string
		radical rune
		want    int
	}{
		{name: "no", radical: 'ノ', want: 1},
		{name: "gate", radical: '門', want: 8},
		{name: "ring", radical: '龠', want: 17},
		{name: "zero", radical: 0, want: 0},
		{name: "empty", radical: ' ', want: 0},
		{name: "not a kanji", radical: 'X', want: 0},
		{name: "not a radical", radical: '右', want: 0},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			r := radicals.Radical(c.radical)
			got := r.StrokeCount()
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

func TestForStrokeCount(t *testing.T) {
	testCases := []struct {
		count int
		want  string
	}{
		{count: 1, want: "一｜丶ノ乙亅"},
		{count: 10, want: "馬骨高髟鬥鬯鬲鬼竜韋"},
		{count: 15, want: ""},
	}
	for _, c := range testCases {
		t.Run(fmt.Sprintf("%d", c.count), func(t *testing.T) {
			got := radicals.ForStrokeCount(c.count)
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

func TestDescriptor(t *testing.T) {
	testCases := []struct {
		name           string
		in             rune
		wantDescriptor string
	}{
		{"mensch links", '⺅', "2a"},
		{"mensch oben", '𠆢', "2a"},
		{"wasser", '⺡', "3a"},
		{"not a kanji", 'x', ""},
		{"not a radical", '右', ""},
		{"empty", ' ', ""},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			r := radicals.Radical(c.in)
			got := r.Descriptor()
			if got != c.wantDescriptor {
				t.Errorf("ERROR: got %v, want %v", got, c.wantDescriptor)
			}
		})
	}
}

func TestForKanji(t *testing.T) {
	testCases := []struct {
		name string
		in   rune
		want string
	}{
		{name: "one", in: '一', want: "一"},
		{name: "wakai", in: '右', want: "ノ一口"},
		{name: "not a kanji", in: 'X', want: ""},
		{name: "empty", in: ' ', want: ""},
		{name: "radical", in: '⺡', want: ""},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := radicals.ForKanji(c.in)
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

func TestAllKanjis(t *testing.T) {
	testCases := []struct {
		name    string
		radical rune
		want    string
	}{
		{name: "zero", radical: 0, want: ""},
		{name: "empty", radical: ' ', want: ""},
		{name: "not a kanji", radical: 'X', want: ""},
		{name: "not a radical", radical: '右', want: ""},
		{name: "bird", radical: '鳥', want: "嗚塢嫣嬝島嶋嶌搗梟槝樢歍烏" +
			"瑦窵篶舃蔦螐裊贗鄔鄥隖隝靎靏鰞鳥鳦鳧鳩鳫鳬鳰鳲鳳鳴鳶鳷鳹鴂鴃鴆" +
			"鴇鴈鴉鴋鴎鴑鴒鴕鴗鴘鴛鴜鴝鴞鴟鴣鴦鴨鴪鴫鴬鴯鴰鴲鴳鴴鴺鴻鴼鴽鴾" +
			"鴿鵁鵂鵃鵄鵅鵆鵇鵈鵊鵐鵑鵓鵔鵙鵜鵝鵞鵟鵠鵡鵢鵣鵤鵥鵩鵪鵫鵬鵯鵰" +
			"鵲鵶鵷鵺鵻鵼鵾鶃鶄鶆鶇鶉鶊鶍鶎鶏鶒鶓鶕鶖鶗鶘鶚鶡鶤鶩鶪鶫鶬鶮鶯" +
			"鶱鶲鶴鶵鶸鶹鶺鶻鶼鶿鷁鷂鷃鷄鷆鷇鷉鷊鷏鷓鷔鷕鷖鷗鷙鷚鷞鷟鷠鷥鷦" +
			"鷧鷩鷫鷭鷮鷯鷰鷲鷳鷴鷸鷹鷺鷽鷾鸂鸇鸊鸎鸐鸑鸒鸕鸖鸙鸚鸛鸜鸝鸞"},
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			r := radicals.Radical(c.radical)
			got := r.AllKanjisWith()
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

func TestIsPartOf(t *testing.T) {
	testCases := []struct {
		name    string
		radical rune
		kanji   rune
		want    bool
	}{
		{"ok", '口', '右', true},
		{"not contained", '⺅', '右', false},
		{"not a radical", 'X', '右', false},
		{"not a kanji", '⺅', 'X', false},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			rad := radicals.Radical(c.radical)
			got := rad.IsPartOf(c.kanji)
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}
