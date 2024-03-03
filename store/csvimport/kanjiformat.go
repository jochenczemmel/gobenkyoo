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

// KanjiFormat defines the field order in the imported csv file.
type KanjiFormat struct {
	fields []string
}

// NewKanjiFormat returns a new kanji format definition.
// If splitchar is not blank, the readings fields are split by the
// given char.
// The list of keys corresponds to the list of fields in the
// csv file.
// Each key must be one of the KanjiField constants (case insensitive)
// or a missing value if the field should be skipped.
func NewKanjiFormat(keys ...string) (KanjiFormat, error) {

	result := KanjiFormat{
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
func (f KanjiFormat) lineToKanjiCard(split string, line []string) kanjis.Card {
	var card kanjis.Card
	var detail kanjis.Detail

	for i, field := range line {
		switch f.field(i) {
		case KanjiFieldKanji:
			kanjiRune, _ := utf8.DecodeRuneInString(field)
			card.Kanji = kanjiRune
		case KanjiFieldReading:
			detail.Reading = field
		case KanjiFieldReadingKana:
			detail.ReadingKana = field
		case KanjiFieldMeanings:
			if split == "" {
				detail.Meanings = append(detail.Meanings, field)
				continue
			}
			detail.Meanings = append(detail.Meanings,
				strings.Split(field, split)...)
		case KanjiFieldHint:
			card.Hint = field
		case KanjiFieldExplanation:
			card.Explanation = field
		}
	}
	card.Details = append(card.Details, detail)

	return card
}

// field returns the value of the definition or a missing
// value if i is out of bounds.
func (f KanjiFormat) field(i int) string {
	if i < 0 || i >= len(f.fields) {
		return ""
	}
	return f.fields[i]
}
