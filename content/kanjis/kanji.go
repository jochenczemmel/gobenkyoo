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

// Card holds the Kanji as a rune, the meanings and the
// readings, an optional hint and an optional explanation.
type Card struct {
	Kanji       rune
	details     []Detail
	uniqDetails map[string]Detail

	// additional infos, might be empty
	Hint        string // hint
	Explanation string // explanation
	ContentInfo string // free text
}

// newCard returns a newCard initialized Kanji object with
// the provided rune.
// Other packages should use Builder to build a newCard Kanji.
func newCard(kanji rune) *Card {
	return &Card{
		Kanji:       kanji,
		uniqDetails: map[string]Detail{},
	}
}

// IsEmpty checks if the kanji is empty.
// func (c Card) IsEmpty() bool {
// return c.kanji == 0
// }

// Rune returns the Kanji as a Rune.
func (c *Card) Rune() rune {
	return c.Kanji
}

// String returns a string representation containing the kanji,
// the descriptor and optionally the number.
func (c Card) String() string {
	result := ""
	if c.Kanji != '\x00' && c.Kanji != ' ' {
		result = string(c.Kanji)
	}

	if d, ok := kanji2Descriptor[c.Kanji]; ok {
		result += fmt.Sprintf(" (%s", d)
		if number, ok := kanji2Nummer[c.Kanji]; ok {
			result += fmt.Sprintf("/%d", number)
		}
		result += ")"
	}

	return result
}

// Descriptor returns the Classification for the 79 radical system.
func (c Card) Descriptor() string {
	return kanji2Descriptor[c.Kanji]
}

// StrokeCount returns the number of strokes of the kanji.
func (c Card) StrokeCount() int {

	first, last, unused := 0, 0, ""
	fmt.Sscanf(kanji2Descriptor[c.Kanji],
		"%d%1s%d.%d", &first, &unused, &last)

	return first + last
}

// Number returns the Hadamitzky Number.
func (c Card) Number() int {
	number, ok := kanji2Nummer[c.Kanji]
	if !ok {
		return 0
	}

	return number
}

// Details returns the list of details.
func (c *Card) Details() []Detail {
	return c.details
}

// addDetails adds a list of details.
// Duplicate readings are not added.
func (c *Card) addDetails(details ...Detail) {
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

// Meanings returns a distinct List of all Meanings.
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

// Readings returns a distinct List of all Readings as romaji.
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

// ReadingsKana returns a distinct List of all Readings as kana
// depending on the case of the first rune:
// upper case is returned as katakana,
// lower case is returned as hiragana.
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

// HasRadical returns true if the given radical
// is part of the kanji on the card.
func (c Card) HasRadical(radical rune) bool {
	rad := radicals.ForKanji(c.Kanji)
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

// Radicals returns the list of the radicals
// for the kanji on the card.
func (c Card) Radicals() string {
	return radicals.ForKanji(c.Kanji)
}
