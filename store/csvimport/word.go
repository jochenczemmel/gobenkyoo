package csvimport

import (
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
	for _, line := range table {
		result = append(result, l.Format.lineToWordCard(line))
	}
	return result
}
