package csvimport

import (
	"encoding/csv"
	"fmt"
	"os"
)

// getLines returns the lines from a csv file.
// The number of fields per line is data dependent.
// If header is true, the first line is omitted.
func getLines(filename string, separator rune, header bool) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open csv file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = separator
	reader.FieldsPerRecord = -1

	table, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv file: %w", err)
	}
	if header {
		return table[1:], nil
	}

	return table, nil
}
