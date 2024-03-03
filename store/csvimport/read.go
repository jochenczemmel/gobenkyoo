package csvimport

import (
	"encoding/csv"
	"os"
)

// getLines returns the consolidated lines from a csv file.
// If header is true, the first line is omitted.
func getLines(filename string, separator rune, header bool) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = separator
	reader.FieldsPerRecord = -1

	table, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if header {
		return table[1:], nil
	}
	return table, nil
}
