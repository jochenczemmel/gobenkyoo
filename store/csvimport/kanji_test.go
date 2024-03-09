package csvimport_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

func TestKanjiImport(t *testing.T) {

	// used for kanjis1.csv
	format1 := []string{
		csvimport.KanjiFieldKanji,
		"", "",
		csvimport.KanjiFieldReading,
		csvimport.KanjiFieldMeanings,
		csvimport.KanjiFieldHint,
		csvimport.KanjiFieldExplanation,
	}

	// used for kanjis1.csv
	minimalFormat1 := []string{
		csvimport.KanjiFieldKanji,
		"", "", "",
		csvimport.KanjiFieldMeanings,
	}

	format1kana := []string{
		csvimport.KanjiFieldKanji,
		"", "",
		csvimport.KanjiFieldReading,
		csvimport.KanjiFieldReadingKana,
		csvimport.KanjiFieldMeanings,
		csvimport.KanjiFieldHint,
		csvimport.KanjiFieldExplanation,
	}

	// used for kanjis2.csv
	format2 := []string{
		csvimport.KanjiFieldKanji,
		"",
		csvimport.KanjiFieldReading,
		csvimport.KanjiFieldReadingKana,
		csvimport.KanjiFieldMeanings,
	}

	testCases := []struct {
		name      string
		fileName  string
		importer  csvimport.Kanji
		want      []kanjis.Card
		wantError bool
	}{{
		name:     "ok",
		fileName: "kanjis1.csv",
		importer: csvimport.NewKanji(';', ' ', true, format1),
		want:     kanji1Cards,
	}, {
		name:     "ok no header",
		fileName: "kanjis1noheader.csv",
		importer: csvimport.NewKanji(';', ' ', false, format1),
		want:     kanji1Cards,
	}, {
		name:     "ok splitted",
		fileName: "kanjis1.csv",
		importer: csvimport.NewKanji(';', '/', true, format1),
		want:     kanji1CardsSplitted,
	}, {
		name:     "ok splitted kana",
		fileName: "kanjis1kana.csv",
		importer: csvimport.NewKanji(';', '/', true, format1kana),
		want:     kanji1CardsKana,
	}, {
		name:     "multiline and kana",
		fileName: "kanjis2.csv",
		importer: csvimport.NewKanji(';', ';', false, format2),
		want:     kanji2Cards,
	}, {
		name:      "file not found",
		fileName:  "does not exist",
		importer:  csvimport.NewKanji(';', ' ', true, format1),
		wantError: true,
	}, {
		name:      "invalid quotes due to wrong separator",
		fileName:  "kanjis1.csv",
		importer:  csvimport.NewKanji(',', ' ', true, format1),
		wantError: true,
	}, {
		name:     "selective fields",
		fileName: "kanjis1.csv",
		importer: csvimport.NewKanji(';', ' ', true, minimalFormat1),
		want:     kanji1CardsMinimal,
	}, {
		name:      "invalid fields",
		fileName:  "kanjis1.csv",
		importer:  csvimport.NewKanji(';', ' ', true, []string{"invalid"}),
		wantError: true,
	}, {
		name:      "empty fields",
		fileName:  "kanjis1.csv",
		importer:  csvimport.NewKanji(';', ' ', true, []string{}),
		wantError: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {

			got, err := c.importer.ImportKanji(
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

var kanji1CardsSplitted = []kanjis.Card{{
	ID:    "1",
	Kanji: '人',
	Details: []kanjis.Detail{{
		Reading:  "hito",
		Meanings: []string{"Mensch"},
	}, {
		Reading:  "NIN",
		Meanings: []string{"Mensch"},
	}, {
		Reading:  "JIN",
		Meanings: []string{"Mensch"},
	}},
	Explanation: "nicht: 入",
}, {
	ID:    "2",
	Kanji: '一',
	Details: []kanjis.Detail{{
		Reading:  "hito.tsu",
		Meanings: []string{"eins"},
	}, {
		Reading:  "ICHI",
		Meanings: []string{"eins"},
	}},
}, {
	ID:    "3",
	Kanji: '二',
	Details: []kanjis.Detail{{
		Reading:  "futa",
		Meanings: []string{"zwei"},
	}, {
		Reading:  "futa.tsu",
		Meanings: []string{"zwei"},
	}, {
		Reading:  "NI",
		Meanings: []string{"zwei"},
	}},
	Hint: "auch ein kana",
}, {
	ID:    "4",
	Kanji: '金',
	Details: []kanjis.Detail{{
		Reading:  "kane",
		Meanings: []string{"Metall", "Geld", "Gold"},
	}, {
		Reading:  "KIN",
		Meanings: []string{"Metall", "Geld", "Gold"},
	}, {
		Reading:  "KON",
		Meanings: []string{"Metall", "Geld", "Gold"},
	}},
}}

var kanji1CardsKana = []kanjis.Card{{
	ID:    "1",
	Kanji: '人',
	Details: []kanjis.Detail{{
		Reading:     "hito",
		ReadingKana: "ひと",
		Meanings:    []string{"Mensch"},
	}, {
		Reading:     "NIN",
		ReadingKana: "ニン",
		Meanings:    []string{"Mensch"},
	}, {
		Reading:     "JIN",
		ReadingKana: "ジン",
		Meanings:    []string{"Mensch"},
	}},
	Explanation: "nicht: 入",
}, {
	ID:    "2",
	Kanji: '一',
	Details: []kanjis.Detail{{
		Reading:     "hito.tsu",
		ReadingKana: "ひと（つ）",
		Meanings:    []string{"eins"},
	}, {
		Reading:     "ICHI",
		ReadingKana: "イチ",
		Meanings:    []string{"eins"},
	}},
}, {
	ID:    "3",
	Kanji: '二',
	Details: []kanjis.Detail{{
		Reading:     "futa",
		ReadingKana: "ふた",
		Meanings:    []string{"zwei"},
	}, {
		Reading:     "futa.tsu",
		ReadingKana: "ふた（つ）",
		Meanings:    []string{"zwei"},
	}, {
		Reading:     "NI",
		ReadingKana: "二",
		Meanings:    []string{"zwei"},
	}},
	Hint: "auch ein kana",
}, {
	ID:    "4",
	Kanji: '金',
	Details: []kanjis.Detail{{
		Reading:     "kane",
		ReadingKana: "かね",
		Meanings:    []string{"Metall", "Geld", "Gold"},
	}, {
		Reading:     "KIN",
		ReadingKana: "キン",
		Meanings:    []string{"Metall", "Geld", "Gold"},
	}, {
		Reading:  "KON",
		Meanings: []string{"Metall", "Geld", "Gold"},
	}},
}}

var kanji2Cards = []kanjis.Card{{
	ID:    "1",
	Kanji: '方',
	Details: []kanjis.Detail{{
		Reading:     "HŌ",
		ReadingKana: "ホウ",
		Meanings:    []string{"Quadrat", "Richtung", "Seite"},
	}, {
		Reading:     "kata",
		ReadingKana: "かた",
		Meanings:    []string{"Methode", "Person", "Richtung"},
	}, {
		Reading:     "masa (ni)",
		ReadingKana: "まさ (に)",
		Meanings:    []string{"wirklich, genau, in der Tat"},
	}},
}, {
	ID:    "2",
	Kanji: '朋',
	Details: []kanjis.Detail{{
		Reading:     "HOO",
		ReadingKana: "ホウ",
		Meanings:    []string{"Kollege, Kamerad"},
	}},
}, {
	ID:    "3",
	Kanji: '島',
	Details: []kanjis.Detail{{
		Reading:     "TOO",
		ReadingKana: "トウ",
		Meanings:    []string{"Insel"},
	}, {
		Reading:     "shima",
		ReadingKana: "しま",
		Meanings:    []string{"Insel"},
	}},
}, {
	ID:    "4",
	Kanji: '曜',
	Details: []kanjis.Detail{{
		Reading:     "YOO",
		ReadingKana: "ヨウ",
		Meanings:    []string{"Licht, leuchten", "Wochentag"},
	}},
}, {
	ID:    "5",
	Kanji: '戸',
	Details: []kanjis.Detail{
		{
			Reading:     "KO",
			ReadingKana: "コ",
			Meanings:    []string{"Haushalt", "Tür"},
		},
		{
			Reading:     "to",
			ReadingKana: "と",
			Meanings:    []string{"Tür"},
		},
	},
}}
