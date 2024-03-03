package csvimport

import (
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// Kanji provides importing csv word files.
type Kanji struct {
	Separator      rune
	FieldSeparator rune
	HeaderLine     bool
	Format         KanjiFormat
}

// Import reads a csv file and returns a list of kanji cards.
func (l Kanji) Import(filename string) ([]kanjis.Card, error) {
	table, err := getLines(filename, l.Separator, l.HeaderLine)
	if err != nil {
		return nil, err
	}
	return l.linesToKanjis(table), nil
}

// linesToKanjis converts a list of lines to a list of kanji cards.
func (l Kanji) linesToKanjis(table [][]string) []kanjis.Card {
	result := make([]kanjis.Card, 0, len(table))
	sep := ""
	if l.FieldSeparator != 0 && l.FieldSeparator != ' ' {
		sep = string(l.FieldSeparator)
	}
	for i, line := range table {
		card := l.Format.lineToKanjiCard(sep, line)
		card.ID = strconv.Itoa(i + 1)
		result = append(result, card)
	}

	return result
}
