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

// ID returns the uniq id of the kanji.
// It is the string representation of the kanji.
func (c Card) ID() string {
	return c.Kanji()
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
		for _, meaning := range detail.Meanings() {
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
		reading := detail.Reading()
		if !found[reading] {
			found[reading] = true
			result = append(result, reading)
		}
	}

	return result
}

// ReadingsKana returns a distinct list of all Readings as kana.
// If the kana have not been specified, they are derived from
// the romaji reading.
//
// The type of kana (hiragana or katakana) depends on the case of the
// first romaji character:
//   - upper case is returned as katakana
//   - lower case is returned as hiragana
//
// There are some bugs in the automatic conversion.
func (c Card) ReadingsKana() []string {

	result := []string{}
	found := map[string]bool{}

	for _, detail := range c.details {
		reading := detail.ReadingKana()
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
		if detail.Reading() == "" {
			continue
		}
		if _, ok := c.uniqDetails[detail.Reading()]; ok {
			continue
		}
		c.uniqDetails[detail.Reading()] = detail
		c.details = append(c.details, detail)
	}
}

// TODO: delete
// Details returns the list of readings and meanings.
func (c *Card) Details() []detail {
	return c.details
}
