package csvimport_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

func TestFormatFields(t *testing.T) {
	testCases := []struct {
		name string
		call func() []string
		want []string
	}{{
		name: "kanji",
		call: csvimport.KanjiFields,
		want: []string{
			"KANJI",
			"READING",
			"READINGKANA",
			"MEANINGS",
			"HINT",
			"EXPLANATION",
		},
	}, {
		name: "word",
		call: csvimport.WordFields,
		want: []string{
			"NIHONGO",
			"KANA",
			"ROMAJI",
			"MEANING",
			"HINT",
			"EXPLANATION",
			"DICTFORM",
			"TEFORM",
			"NAIFORM",
		},
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := c.call()
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
