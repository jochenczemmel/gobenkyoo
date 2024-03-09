package csvimport

import (
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Word provides importing csv word files.
type Word struct {
	separator  rune
	headerLine bool
	fields     []string
}

func NewWord(sep rune, header bool, fields []string) Word {
	return Word{
		separator:  sep,
		headerLine: header,
		fields:     fields,
	}
}

func (l Word) ImportWord(filename string) ([]words.Card, error) {
	var result []words.Card
	format, err := newWordFormat(l.fields...)
	if err != nil {
		return result, err
	}
	table, err := getLines(filename, l.separator, l.headerLine)
	if err != nil {
		return nil, err
	}
	return l.linesToWords(format, table), nil
}

// linesToWords converts a list of lines to a list of word cards.
func (l Word) linesToWords(format wordFormat, table [][]string) []words.Card {
	result := make([]words.Card, 0, len(table))
	for i, line := range table {
		card := format.lineToWordCard(line)
		card.ID = strconv.Itoa(i + 1)
		result = append(result, card)
	}
	return result
}
