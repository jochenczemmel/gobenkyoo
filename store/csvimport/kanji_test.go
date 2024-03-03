package csvimport_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

func TestKanjiImport(t *testing.T) {

	format, _ := csvimport.NewKanjiFormat(
		csvimport.KanjiFieldKanji,
		"", "",
		csvimport.KanjiFieldReading,
		csvimport.KanjiFieldMeanings,
		csvimport.KanjiFieldHint,
		csvimport.KanjiFieldExplanation,
	)

	minimalFormat, _ := csvimport.NewKanjiFormat(
		csvimport.KanjiFieldKanji,
		"", "", "",
		csvimport.KanjiFieldMeanings,
	)

	testCases := []struct {
		name      string
		fileName  string
		importer  csvimport.Kanji
		want      []kanjis.Card
		wantError bool
	}{{
		name:     "ok",
		fileName: "kanjis1.csv",
		importer: csvimport.Kanji{
			Format:     format,
			Separator:  ';',
			HeaderLine: true,
		},
		want: kanji1Cards,
	}, {
		name:     "ok no header",
		fileName: "kanjis1noheader.csv",
		importer: csvimport.Kanji{
			Format:    format,
			Separator: ';',
		},
		want: kanji1Cards,
	}, {
		name:     "file not found",
		fileName: "does not exist",
		importer: csvimport.Kanji{
			Format:     format,
			Separator:  ';',
			HeaderLine: true,
		},
		wantError: true,
	}, {
		name:     "invalid quotes due to wrong separator",
		fileName: "kanjis1.csv",
		importer: csvimport.Kanji{
			Format:     format,
			Separator:  ',',
			HeaderLine: true,
		},
		wantError: true,
	}, {
		name:     "selective fields",
		fileName: "kanjis1.csv",
		importer: csvimport.Kanji{
			Format:     minimalFormat,
			Separator:  ';',
			HeaderLine: true,
		},
		want: kanji1CardsMinimal,
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

var kanji1Cards = []kanjis.Card{{
	ID:    "1",
	Kanji: '人',
	Details: []kanjis.Detail{{
		Reading:  "hito/NIN/JIN",
		Meanings: []string{"Mensch"},
	}},
	Explanation: "nicht: 入",
}, {
	ID:    "2",
	Kanji: '一',
	Details: []kanjis.Detail{{
		Reading:  "hito.tsu/ICHI",
		Meanings: []string{"eins"},
	}},
}, {
	ID:    "3",
	Kanji: '二',
	Details: []kanjis.Detail{{
		Reading:  "futa/futa.tsu/NI",
		Meanings: []string{"zwei"},
	}},
	Hint: "auch ein kana",
}, {
	ID:    "4",
	Kanji: '金',
	Details: []kanjis.Detail{{
		Reading:  "kane/KIN/KON",
		Meanings: []string{"Metall/Geld/Gold"},
	}},
}}

var kanji1CardsMinimal = []kanjis.Card{{
	ID:    "1",
	Kanji: '人',
	Details: []kanjis.Detail{{
		Meanings: []string{"Mensch"},
	}},
}, {
	ID:    "2",
	Kanji: '一',
	Details: []kanjis.Detail{{
		Meanings: []string{"eins"},
	}},
}, {
	ID:    "3",
	Kanji: '二',
	Details: []kanjis.Detail{{
		Meanings: []string{"zwei"},
	}},
}, {
	ID:    "4",
	Kanji: '金',
	Details: []kanjis.Detail{{
		Meanings: []string{"Metall/Geld/Gold"},
	}},
}}
