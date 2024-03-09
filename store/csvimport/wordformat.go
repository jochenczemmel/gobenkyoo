package csvimport

import (
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// valid values for field names in word csv files.
const (
	WordFieldNihongo     = "NIHONGO"
	WordFieldKana        = "KANA"
	WordFieldRomaji      = "ROMAJI"
	WordFieldMeaning     = "MEANING"
	WordFieldHint        = "HINT"
	WordFieldExplanation = "EXPLANATION"
	WordFieldDictform    = "DICTFORM"
	WordFieldTeform      = "TEFORM"
	WordFieldNaiform     = "NAIFORM"
)

func WordFields() []string {
	return []string{
		WordFieldNihongo,
		WordFieldKana,
		WordFieldRomaji,
		WordFieldMeaning,
		WordFieldHint,
		WordFieldExplanation,
		WordFieldDictform,
		WordFieldTeform,
		WordFieldNaiform,
	}
}

// wordFormat defines the field order in the imported csv file.
type wordFormat struct {
	fields []string
}

// NewWordFormat returns a new word format definition.
// Keys must be one of the WordField constants (case insensitive)
// or a missing value if the field should be skipped.
func newWordFormat(keys ...string) (wordFormat, error) {

	result := wordFormat{
		fields: make([]string, 0, len(keys)),
	}
	if len(keys) < 1 {
		return result, ErrNoKey
	}

	for i, key := range keys {
		key = strings.ToUpper(key)
		switch key {
		case WordFieldNihongo, WordFieldKana, WordFieldRomaji,
			WordFieldMeaning, WordFieldHint, WordFieldExplanation,
			WordFieldDictform, WordFieldTeform, WordFieldNaiform, "":
		default:
			return result, InvalidKeyError(keys[i])
		}
		result.fields = append(result.fields, key)
	}

	return result, nil
}

// lineToWordCard creates a word card based on the fields in
// the line and the field definition.
func (f wordFormat) lineToWordCard(line []string) words.Card {
	var card words.Card
	for i, field := range line {
		switch f.field(i) {

		case WordFieldNihongo:
			card.Nihongo = field
		case WordFieldKana:
			card.Kana = field
		case WordFieldRomaji:
			card.Romaji = field
		case WordFieldMeaning:
			card.Meaning = field
		case WordFieldHint:
			card.Hint = field
		case WordFieldExplanation:
			card.Explanation = field
		case WordFieldDictform:
			card.DictForm = field
		case WordFieldTeform:
			card.TeForm = field
		case WordFieldNaiform:
			card.NaiForm = field
		}
	}

	return card
}

// field returns the value of the definition or a missing
// value if i is out of bounds.
func (f wordFormat) field(i int) string {
	if i < 0 || i >= len(f.fields) {
		return ""
	}
	return f.fields[i]
}
