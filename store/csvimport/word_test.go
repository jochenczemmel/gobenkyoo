package csvimport_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/words"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

const testDataDir = "testdata"

func TestWordImport(t *testing.T) {

	format, _ := csvimport.NewWordFormat(
		"KANA", "NIHONGO", "ROMAJI", "MEANING", "HINT",
		"EXPLANATION", "DICTFORM", "TEFORM", "NAIFORM",
	)
	minimalFormat, _ := csvimport.NewWordFormat(
		"", "", "ROMAJI", "MEANING")

	testCases := []struct {
		name      string
		fileName  string
		importer  csvimport.Words
		want      []words.Card
		wantError bool
	}{{
		name:     "ok",
		fileName: "words1.csv",
		importer: csvimport.Words{
			Format:     format,
			Separator:  ',',
			HeaderLine: true,
		},
		want: word1Cards,
	}, {
		name:     "ok no header",
		fileName: "words1noheader.csv",
		importer: csvimport.Words{
			Format:    format,
			Separator: ',',
		},
		want: word1Cards,
	}, {
		name:     "selective fields",
		fileName: "words1.csv",
		importer: csvimport.Words{
			Format:     minimalFormat,
			Separator:  ',',
			HeaderLine: true,
		},
		want: word1CardsMinimal,
	}, {
		name:     "file not found",
		fileName: "does not exist",
		importer: csvimport.Words{
			Format:     minimalFormat,
			Separator:  ',',
			HeaderLine: true,
		},
		wantError: true,
	}, {
		name:     "invalid quotes due to wrong separator",
		fileName: "words1.csv",
		importer: csvimport.Words{
			Format:     format,
			Separator:  ';',
			HeaderLine: true,
		},
		wantError: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			got, err := c.importer.Import(
				filepath.Join(testDataDir, c.fileName))

			if c.wantError {
				if err == nil {
					t.Fatalf("ERROR: wanted error not detected")
				}
				t.Logf("INFO: error message: %v", err)
				return
			}

			if err != nil {
				t.Fatalf("ERROR: got error: %v", err)
			}
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}

var word1Cards = []words.Card{
	{
		Nihongo:     "先生",
		Kana:        "せんせい",
		Romaji:      "sensei",
		Meaning:     "Lehrer",
		Hint:        "andere Personen",
		Explanation: "für sich selbst anderer Ausdruck",
	},
	{
		Nihongo: "医者",
		Kana:    "いしゃ",
		Romaji:  "isha",
		Meaning: "Arzt, Ärztin"},
	{
		Nihongo: "お名前\u3000は\u3000「何\u3000です\u3000か」。",
		Kana:    "お\u3000なまえ\u3000は\u3000「なん\u3000です\u3000か」。",
		Romaji:  "onamae wa (nan desu ka).",
		Meaning: "Wie heißen Sie bitte?",
	},
	{
		Nihongo:  "起きます",
		Kana:     "おきます",
		Romaji:   "okimasu",
		Meaning:  "aufstehen",
		DictForm: "おきる",
		TeForm:   "おきて",
		NaiForm:  "おきない",
	},
}

var word1CardsMinimal = []words.Card{
	{
		Romaji:  "sensei",
		Meaning: "Lehrer",
	},
	{
		Romaji:  "isha",
		Meaning: "Arzt, Ärztin"},
	{
		Romaji:  "onamae wa (nan desu ka).",
		Meaning: "Wie heißen Sie bitte?",
	},
	{
		Romaji:  "okimasu",
		Meaning: "aufstehen",
	},
}
