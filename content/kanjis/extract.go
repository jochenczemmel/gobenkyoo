package kanjis

import (
	"strings"
	"unicode"

	"github.com/gojp/kana"
)

// Extract returns a string of uniq kanjis found in the given text.
// The result does not contain duplicates.
// The order of the Kanjis is according to their first appearance in the text.
func Extract(text string) []rune {

	result := []rune{}
	seen := map[rune]bool{}

	for _, r := range text {
		if !isKanji(r) {
			continue
		}
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}

	return result
}

// isKanji returns true if the kanji is found in the predefined list.
// Unfortunately, Unicode has no simple way to identify a kanji character,
// so we use the data from edict.
func isKanji(kanji rune) bool {
	_, ok := kanji2Descriptor[kanji]
	return ok
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
