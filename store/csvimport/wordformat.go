package csvimport

import (
	"errors"
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

var ErrNoKey = errors.New("no keys defined")

type InvalidKeyError string

func (e InvalidKeyError) Error() string {
	return "invalid key: " + string(e)
}

type WordFormat struct {
	fields []string
}

func NewWordFormat(keys ...string) (WordFormat, error) {

	result := WordFormat{
		fields: make([]string, 0, len(keys)),
	}
	if len(keys) < 1 {
		return result, ErrNoKey
	}

	for i, key := range keys {
		key = strings.ToUpper(key)
		switch key {
		case "NIHONGO", "KANA", "ROMAJI", "MEANING", "HINT",
			"EXPLANATION", "DICTFORM", "TEFORM", "NAIFORM", "":
		default:
			return result, InvalidKeyError(keys[i])
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
