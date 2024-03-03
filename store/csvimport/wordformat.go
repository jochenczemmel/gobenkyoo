package csvimport

import (
	"fmt"
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

type WordFormat struct {
	fields []string
}

func NewWordFormat(keys ...string) (WordFormat, error) {

	result := WordFormat{
		fields: make([]string, 0, len(keys)),
	}
	if len(keys) < 1 {
		return result, fmt.Errorf("no keys defined")
	}

	for i, key := range keys {
		key = strings.ToUpper(key)
		switch key {
		case "NIHONGO", "KANA", "ROMAJI", "MEANING", "HINT",
			"EXPLANATION", "DICTFORM", "TEFORM", "NAIFORM", "":
		default:
			return result, fmt.Errorf("invalid key: %q", keys[i])
		}
		result.fields = append(result.fields, key)
	}

	return result, nil
}

func (f WordFormat) lineToWordCard(line []string) words.Card {
	var card words.Card
	for i, field := range line {
		switch f.field(i) {
		case "NIHONGO":
			card.Nihongo = field
		case "KANA":
			card.Kana = field
		case "ROMAJI":
			card.Romaji = field
		case "MEANING":
			card.Meaning = field
		case "HINT":
			card.Hint = field
		case "EXPLANATION":
			card.Explanation = field
		case "DICTFORM":
			card.DictForm = field
		case "TEFORM":
			card.TeForm = field
		case "NAIFORM":
			card.NaiForm = field
		}
	}
	return card
}

func (f WordFormat) field(i int) string {
	if i < 0 || i >= len(f.fields) {
		return ""
	}
	return f.fields[i]
}
