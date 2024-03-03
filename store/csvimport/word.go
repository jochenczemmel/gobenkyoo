package csvimport

import (
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

type Words struct {
	Separator  rune
	HeaderLine bool
	Format     WordFormat
}

func (l Words) Import(filename string) ([]words.Card, error) {
	table, err := getLines(filename, l.Separator, l.HeaderLine)
	if err != nil {
		return nil, err
	}
	return l.linesToWords(table), nil
}

func (l Words) linesToWords(table [][]string) []words.Card {
	result := make([]words.Card, 0, len(table))
	for i, line := range table {
		card := l.Format.lineToWordCard(line)
		card.ID = strconv.Itoa(i + 1)
		result = append(result, card)
	}
	return result
}
