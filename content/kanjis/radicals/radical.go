// Package radicals provides information about
// the elements (radicals) of kanjis.
package radicals

import "sort"

func init() {
	fillInvertedMaps()
}

// AllKanjisWith returns all kanjis that contain the given radical.
// The returned kanjis are sorted.
func AllKanjisWith(radical rune) string {
	return string(radical2Kanjis[radical])
}

// StrokeCount returns the stroke count of the radical.
func StrokeCount(radical rune) int {
	return kanjiStrokeCount[radical]
}

// ForStrokeCount returns a list of all radicals with the
// given stroke count.
func ForStrokeCount(strokecount int) string {
	return strokecount2Radicals[strokecount]
}

// ForKanji returns the list of radicals for the
// given Kanji.
func ForKanji(kanji rune) string {
	return kanji2Radical[kanji]
}

// Descriptor returns the descriptor prefix
// for the given Radical.
func Descriptor(radical rune) string {
	return radical2Descriptor[radical]
}

// radical2Kanjis contains the inverted data
// radikal -> kanji list.
var radical2Kanjis = map[rune][]rune{}

// kanjiStrokeCount contains the inverted data
// radical -> stroke count.
var kanjiStrokeCount = map[rune]int{}

// fillInvertedMaps creates some inverted maps from the
// predefined radical data:
//   - radical2Kanji
//   - kanjiStrokeCount
func fillInvertedMaps() {
	fillRadical2Kanjis()
	sortRadical2Kanjis()
	fillKanjiStrokeCount()
}

// fillRadical2Kanjis appends the kanji list of the kradfile and kradfile2 files.
func fillRadical2Kanjis() {
	// kradfile
	for kanji, list := range kanji2Radical {
		for _, radical := range list {
			radical2Kanjis[radical] = append(radical2Kanjis[radical], kanji)
		}
	}
	for kanji, list := range kanji2Radical2 {
		for _, radical := range list {
			radical2Kanjis[radical] = append(radical2Kanjis[radical], kanji)
		}
	}
}

// sortRadical2Kanjis sorts the kanjis in the radical2Kanji map.
func sortRadical2Kanjis() {
	// sort kanjis to have a stable output
	for radical, kanjis := range radical2Kanjis {
		sort.Slice(kanjis, func(i, j int) bool {
			return kanjis[i] < kanjis[j]
		})
		radical2Kanjis[radical] = kanjis
	}
}

// fillKanjiStrokeCount fills the map that holds the stroke counts
// for the kanjis.
func fillKanjiStrokeCount() {
	// stroke count
	for count, list := range strokecount2Radicals {
		for _, kanji := range list {
			kanjiStrokeCount[kanji] = count
		}
	}
}

// Radical provides methods for runes that are kanji radicals.
type Radical rune

// StrokeCount returns the stroke count of the radical.
func (r Radical) StrokeCount() int {
	return kanjiStrokeCount[rune(r)]
}

// DescriptorPrefix returns the descriptor prefix
// for the given Radical.
func (r Radical) DescriptorPrefix() string {
	return radical2Descriptor[rune(r)]
}

// AllKanjisWith returns all kanjis that contain the given radical.
func (r Radical) AllKanjisWith() string {
	return string(radical2Kanjis[rune(r)])
}
