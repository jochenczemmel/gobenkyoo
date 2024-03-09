package csvimport

import (
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// Kanji provides importing csv word files.
type Kanji struct {
	separator      rune
	fieldSeparator rune
	headerLine     bool
	fields         []string
}

func NewKanji(sep, fieldsep rune, header bool, fields []string) Kanji {
	return Kanji{
		separator:      sep,
		fieldSeparator: fieldsep,
		headerLine:     header,
		fields:         fields,
	}
}

func (l Kanji) ImportKanji(filename string) ([]kanjis.Card, error) {
	var result []kanjis.Card
	format, err := newKanjiFormat(l.fields...)
	if err != nil {
		return result, err
	}
	table, err := getLines(filename, l.separator, l.headerLine)
	if err != nil {
		return nil, err
	}
	return l.linesToKanjis(format, table), nil
}

// linesToKanjis converts a list of lines to a list of kanji cards.
func (l Kanji) linesToKanjis(format kanjiFormat, table [][]string) []kanjis.Card {

	result := make([]kanjis.Card, 0, len(table))
	sep := ""
	if l.fieldSeparator != 0 && l.fieldSeparator != ' ' {
		sep = string(l.fieldSeparator)
	}

	seen := make(map[rune]*kanjis.Card, len(table))
	order := make([]rune, 0, len(table))

	i := 0
	for _, line := range table {
		card := format.lineToKanjiCard(sep, line)

		seenCard, ok := seen[card.Kanji]
		if ok {
			seenCard.Details = append(seenCard.Details, card.Details...)
			continue
		}
		i++
		card.ID = strconv.Itoa(i)
		seen[card.Kanji] = &card
		order = append(order, card.Kanji)
	}

	for _, k := range order {
		result = append(result, *seen[k])
	}

	return result
}
