package kanjis

import (
	"fmt"
	"strings"
)

// Details holds a single pair of reading and meanings.
type Detail struct {
	// reading in Romaji
	reading string
	// reading in Hiragana/Katakana
	readingKana string
	// meanings in target language
	meanings       []string
	uniqueMeanings map[string]bool
}

// newDetail returns a new Details object.
func newDetail(reading string, meanings ...string) Detail {
	result := Detail{
		reading:        reading,
		uniqueMeanings: map[string]bool{},
	}
	result.addMeanings(meanings...)

	return result
}

// newDetailKana returns a new Details object with explicit
// kana specification.
func newDetailKana(reading, kana string, meanings ...string) Detail {
	result := newDetail(reading, meanings...)
	result.readingKana = kana

	return result
}

// addMeanings adds meanings if the do not yet exist.
func (d *Detail) addMeanings(meanings ...string) {
	for _, m := range meanings {
		if !d.uniqueMeanings[m] {
			d.uniqueMeanings[m] = true
			d.meanings = append(d.meanings, m)
		}
	}
}

// String returns a string representation containing the reading
// in romaji, the reading in kana and the list of meanings.
// Mainly useful for test debugging.
func (d Detail) String() string {
	return fmt.Sprintf("%s (%s): %s",
		d.reading,
		d.ReadingKana(),
		strings.Join(d.meanings, ", "),
	)
}

// Meanings returns the kanji meanings in the target language.
func (d Detail) Meanings() []string {
	return d.meanings
}

// Readings returns the Readings as Romaji.
func (d Detail) Reading() string {
	return d.reading
}

// ReadingKana returns the Readings as Hiragana or Katakana.
// If no kana readings are provided, they are derived from the
// romaji readings.
func (d Detail) ReadingKana() string {
	if d.readingKana == "" {
		return ToKana(d.reading)
	}

	return d.readingKana
}
