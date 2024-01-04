// Package kanjis handles kanji related data.
//
// Kanji data consist of a kanji character (a rune in go),
// some metadata as the kanji descriptor or the
// spahn-hadamitzky-number,
// and associated readings and meanings.
package kanjis

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals"
)

// Card holds the kanji as a rune, a list of the readings and
// their meanings, an optional hint and an optional explanation.
// Use kanjis.Builder to create a new kanjis card.
type Card struct {
	kanji       rune
	Hint        string
	Explanation string
	details     []detail
	uniqDetails map[string]detail
}

// newCard returns a newCard initialized kanji object with
// the provided rune.
func newCard(kanji rune) Card {
	return Card{
		kanji:       kanji,
		uniqDetails: map[string]detail{},
	}
}

// Kanji returns the kanji as a string.
func (c *Card) Kanji() string {
	// avoid display of \x00
	if c.kanji == '\x00' || c.kanji == ' ' {
		return ""
	}
	return string(c.kanji)
}

// Descriptor returns the classification for the 79 radical system.
func (c Card) Descriptor() string {
	return kanji2Descriptor[c.kanji]
}

// Radicals returns the list of the radicals
// for the kanji on the card.
func (c Card) Radicals() string {
	return radicals.ForKanji(c.kanji)
}

// StrokeCount returns the number of strokes of the kanji.
func (c Card) StrokeCount() int {
	first, last, unused := 0, 0, ""
	fmt.Sscanf(kanji2Descriptor[c.kanji],
		"%d%1s%d.%d", &first, &unused, &last)
	return first + last
}

// Number returns the Hadamitzky Number.
func (c Card) Number() int {
	number, ok := kanji2Nummer[c.kanji]
	if !ok {
		return 0
	}
	return number
}

// Description returns a string representation containing the kanji,
// the descriptor and optionally the number.
func (c Card) Description() string {
	result := ""
	if c.kanji != '\x00' && c.kanji != ' ' {
		result = string(c.kanji)
	}

	if d, ok := kanji2Descriptor[c.kanji]; ok {
		result += fmt.Sprintf(" (%s", d)
		if number, ok := kanji2Nummer[c.kanji]; ok {
			result += fmt.Sprintf("/%d", number)
		}
		result += ")"
	}

	return result
}

// HasRadical returns true if the given radical
// is part of the kanji on the card.
func (c Card) HasRadical(radical rune) bool {
	rad := radicals.ForKanji(c.kanji)
	if rad == "" {
		return false
	}
	for _, r := range rad {
		if r == radical {
			return true
		}
	}

	return false
}

// Meanings returns a distinct list of all Meanings.
func (c Card) Meanings() []string {
	result := []string{}
	found := map[string]bool{}
	for _, detail := range c.details {
		for _, meaning := range detail.meanings {
			if !found[meaning] {
				found[meaning] = true
				result = append(result, meaning)
			}
		}
	}

	return result
}

// Readings returns a distinct list of all readings as romaji.
func (c Card) Readings() []string {
	result := []string{}
	found := map[string]bool{}
	for _, detail := range c.details {
		reading := detail.reading
		if !found[reading] {
			found[reading] = true
			result = append(result, reading)
		}
	}

	return result
}

// ReadingsKana returns a distinct list of all Readings as kana.
// Missing kana readings are not added to the list.
func (c Card) ReadingsKana() []string {

	result := []string{}
	found := map[string]bool{}

	for _, detail := range c.details {
		reading := detail.readingKana
		if reading == "" {
			continue
		}
		if !found[reading] {
			found[reading] = true
			result = append(result, reading)
		}
	}

	return result
}

// addDetails adds a list of details.
// Duplicate readings are not added.
func (c *Card) addDetails(details ...detail) {
	for _, detail := range details {
		if detail.reading == "" {
			continue
		}
		if _, ok := c.uniqDetails[detail.reading]; ok {
			continue
		}
		c.uniqDetails[detail.reading] = detail
		c.details = append(c.details, detail)
	}
}
