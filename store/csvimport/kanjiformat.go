package csvimport

import (
	"strings"
	"unicode/utf8"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// valid values for field names in kanji csv files.
const (
	KanjiFieldKanji       = "KANJI"
	KanjiFieldReading     = "READING"
	KanjiFieldReadingKana = "READINGKANA"
	KanjiFieldMeanings    = "MEANINGS"
	KanjiFieldHint        = "HINT"
	KanjiFieldExplanation = "EXPLANATION"
)

func KanjiFields() []string {
	return []string{
		KanjiFieldKanji,
		KanjiFieldReading,
		KanjiFieldReadingKana,
		KanjiFieldMeanings,
		KanjiFieldHint,
		KanjiFieldExplanation,
	}
}

// kanjiFormat defines the field order in the imported csv file.
type kanjiFormat struct {
	fields []string
}

// newKanjiFormat returns a new kanji format definition.
// If splitchar is not blank, the readings fields are split by the
// given char.
// The list of keys corresponds to the list of fields in the
// csv file.
// Each key must be one of the KanjiField constants (case insensitive)
// or a missing value if the field should be skipped.
func newKanjiFormat(keys ...string) (kanjiFormat, error) {

	result := kanjiFormat{
		fields: make([]string, 0, len(keys)),
	}
	if len(keys) < 1 {
		return result, ErrNoKey
	}

	for i, key := range keys {
		key = strings.ToUpper(key)
		switch key {
		case KanjiFieldKanji, KanjiFieldReading, KanjiFieldReadingKana,
			KanjiFieldMeanings, KanjiFieldHint, KanjiFieldExplanation, "":
		default:
			return result, InvalidKeyError(keys[i])
		}
		result.fields = append(result.fields, key)
	}

	return result, nil
}

// lineToKanjiCard creates a kanji card based on the fields in
// the line and the field definition.
func (f kanjiFormat) lineToKanjiCard(split string, line []string) kanjis.Card {
	card, detail, readings, readingsKana := f.fillFields(line, split)
	if len(readings) < 1 {
		card.Details = append(card.Details, detail)
	} else {
		for i, r := range readings {
			detail.Reading = r
			if i < len(readingsKana) {
				detail.ReadingKana = readingsKana[i]
			} else {
				detail.ReadingKana = ""
			}
			card.Details = append(card.Details, detail)
		}
	}

	return card
}

func (f kanjiFormat) fillFields(line []string, split string) (kanjis.Card, kanjis.Detail, []string, []string) {

	var card kanjis.Card
	var detail kanjis.Detail
	var readings []string
	var readingsKana []string

	for i, field := range line {
		switch f.field(i) {
		case KanjiFieldKanji:
			kanjiRune, _ := utf8.DecodeRuneInString(field)
			card.Kanji = kanjiRune
		case KanjiFieldReading:
			if split != "" {
				for _, f := range strings.Split(field, split) {
					readings = append(readings, strings.TrimSpace(f))
				}
				continue
			}
			detail.Reading = field
		case KanjiFieldReadingKana:
			for _, f := range strings.Split(field, split) {
				readingsKana = append(readingsKana, strings.TrimSpace(f))
			}
			continue
		case KanjiFieldMeanings:
			if split == "" {
				detail.Meanings = append(detail.Meanings, field)
				continue
			}
			for _, f := range strings.Split(field, split) {
				detail.Meanings = append(detail.Meanings, strings.TrimSpace(f))
			}
		case KanjiFieldHint:
			card.Hint = field
		case KanjiFieldExplanation:
			card.Explanation = field
		}
	}

	return card, detail, readings, readingsKana
}

// field returns the value of the definition or a missing
// value if i is out of bounds.
func (f kanjiFormat) field(i int) string {
	if i < 0 || i >= len(f.fields) {
		return ""
	}
	return f.fields[i]
}
