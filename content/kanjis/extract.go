package kanjis

import (
	"strings"
	"unicode"

	"github.com/gojp/kana"
)

// Extract returns a string of uniq kanjis found in the given text.
// The resulting slice does not contain duplicates. The
// order of the Kanjis is according to the appearance in the text.
func Extract(text string) []rune {

	result := []rune{}
	seen := map[rune]bool{}

	for _, r := range text {
		if _, ok := kanji2Descriptor[r]; ok {
			if _, found := seen[r]; !found {
				seen[r] = true
				result = append(result, r)
			}
		}
	}
	return result
}

// ToKana converts romaji to kana. If the first Character ist Upcase,
// it returns katakana, else it returns hiragana.
// Currently, it is not correct with respect to katakana -,
// long o with ou, and maybe other errors.
func ToKana(romaji string) string {
	lowcase := strings.ToLower(romaji)
	if len(romaji) < 1 {
		return ""
	}
	if unicode.IsUpper([]rune(romaji)[0]) {
		return kana.RomajiToKatakana(kana.NormalizeRomaji(lowcase))
	}
	return kana.RomajiToHiragana(kana.NormalizeRomaji(lowcase))
}
