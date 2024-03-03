package csvimport_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

func TestFormatErrors(t *testing.T) {
	testCases := []struct {
		name         string
		input        []string
		wantWordErr  bool
		wantKanjiErr bool
	}{{
		name:  "ok",
		input: []string{"HINT", "", "EXPLANATION"},
	}, {
		name:         "word ok",
		input:        []string{"NIHONGO", "", "KANA"},
		wantKanjiErr: true,
	}, {
		name:        "kanji ok",
		input:       []string{"KANJI", "", "READINGKANA"},
		wantWordErr: true,
	}, {
		name:         "wrong content",
		input:        []string{"HINT", "CHINESE"},
		wantWordErr:  true,
		wantKanjiErr: true,
	}, {
		name:         "no fields",
		input:        []string{},
		wantWordErr:  true,
		wantKanjiErr: true,
	}, {
		name:         "nil fields",
		wantWordErr:  true,
		wantKanjiErr: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("word", func(t *testing.T) {
				_, err := csvimport.NewWordFormat(c.input...)
				if c.wantWordErr {
					if err == nil {
						t.Fatalf("ERROR: wanted error not detected")
					}
					t.Logf("INFO: got error: %v", err)
					return
				}
				if err != nil {
					t.Fatalf("ERROR: got error: %v", err)
				}
			})
			t.Run("kanji", func(t *testing.T) {
				_, err := csvimport.NewKanjiFormat(c.input...)
				if c.wantKanjiErr {
					if err == nil {
						t.Fatalf("ERROR: wanted error not detected")
					}
					t.Logf("INFO: got error: %v", err)
					return
				}
				if err != nil {
					t.Fatalf("ERROR: got error: %v", err)
				}
			})
		})
	}
}
