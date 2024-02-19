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
	ID          string
	Kanji       rune
	Hint        string
	Explanation string
	Details     []Detail
}

/*
// New returns a New initialized kanji object with
// the provided rune.
func New(kanji rune) Card {
	return Card{
		Kanji: kanji,
	}
}
*/

// String returns the kanji as a string.
func (c Card) String() string {
	if c.Kanji == '\x00' || c.Kanji == ' ' {
		return ""
	}
	return string(c.Kanji)
}

// Descriptor returns the classification for the 79 radical system.
func (c Card) Descriptor() string {
	return kanji2Descriptor[c.Kanji]
}

// Radicals returns the list of the radicals
// for the kanji on the card.
func (c Card) Radicals() string {
	return radicals.ForKanji(c.Kanji)
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

// Description returns a string representation containing the kanji,
// the descriptor and optionally the number.
func (c Card) Description() string {
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

// HasRadical returns true if the given radical
// is part of the kanji on the card.
func (c Card) HasRadical(radical rune) bool {
	return radicals.Radical(radical).IsPartOf(c.Kanji)
}

// Meanings returns a distinct list of all Meanings.
func (c Card) Meanings() []string {
	result := []string{}
	found := map[string]bool{}
	for _, detail := range c.Details {
		for _, meaning := range detail.Meanings {
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
	for _, detail := range c.Details {
		reading := detail.Reading
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

	for _, detail := range c.Details {
		reading := detail.ReadingKana
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
