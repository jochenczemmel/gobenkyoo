package csvimport

import (
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Word provides importing csv word files.
type Word struct {
	Separator  rune
	HeaderLine bool
	Format     WordFormat
}

// Import reads a csv file and returns a list of word cards.
func (l Word) Import(filename string) ([]words.Card, error) {
	table, err := getLines(filename, l.Separator, l.HeaderLine)
	if err != nil {
		return nil, err
	}
	return l.linesToWords(table), nil
}

// linesToWords converts a list of lines to a list of word cards.
func (l Word) linesToWords(table [][]string) []words.Card {
	result := make([]words.Card, 0, len(table))
	for i, line := range table {
		card := l.Format.lineToWordCard(line)
		card.ID = strconv.Itoa(i + 1)
		result = append(result, card)
	}
	return result
}
