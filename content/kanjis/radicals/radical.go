// Package radicals provides information about
// the elements (radicals) of kanjis.
package radicals

import "sort"

func init() {
	// always create inverted map
	fillInvertedMaps()
}

// AllKanjisContaining returns all kanjis that contain the given radical.
func AllKanjisContaining(radical rune) string {
	return string(radical2Kanji[radical])
}

// StrokeCount returns the stroke count of the radical.
func StrokeCount(radical rune) int {
	return kanjiStrokeCount[radical]
}

// ForStrokeCount returns a list of all radicals with the
// given stroke count.
func ForStrokeCount(strokecount int) string {
	return strokecount2Radical[strokecount]
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

// radical2Kanji contains the inverted data
// radikal -> kanji list.
var radical2Kanji = map[rune][]rune{}

// kanjiStrokeCount contains the inverted data
// radical -> stroke count.
var kanjiStrokeCount = map[rune]int{}

// fillInvertedMaps creates some inverted maps:
// - radical2Kanji
// - kanjiStrokeCount
func fillInvertedMaps() {
	fillRadical2Kanji()
	sortRadical2Kanjis()
	fillKanjiStrokeCount()
}

// fillRadical2Kanji appends the kanji list of the kradfile and kradfile2 files.
func fillRadical2Kanji() {
	// kradfile
	for kanji, list := range kanji2Radical {
		for _, radical := range list {
			radical2Kanji[radical] = append(radical2Kanji[radical], kanji)
		}
	}
	for kanji, list := range kanji2Radical2 {
		for _, radical := range list {
			radical2Kanji[radical] = append(radical2Kanji[radical], kanji)
		}
	}
}

// sortRadical2Kanjis sorts the kanjis in the radical2Kanji map.
func sortRadical2Kanjis() {
	// sort kanjis to have a stable output
	for radical, kanjis := range radical2Kanji {
		sort.Slice(kanjis, func(i, j int) bool {
			return kanjis[i] < kanjis[j]
		})
		radical2Kanji[radical] = kanjis
	}
}

// fillKanjiStrokeCount fills the map that holds the stroke counts
// for the kanjis.
func fillKanjiStrokeCount() {
	// stroke count
	for count, list := range strokecount2Radical {
		for _, kanji := range list {
			kanjiStrokeCount[kanji] = count
		}
	}
}

// Radical provides methods for runes that are radicals.
type Radical rune

// StrokeCount returns the stroke count of the radical.
func (r Radical) StrokeCount() int {
	return kanjiStrokeCount[rune(r)]
}

// Descriptor returns the descriptor prefix
// for the given Radical.
func (r Radical) Descriptor() string {
	return radical2Descriptor[rune(r)]
}

// AllKanjisContaining returns all kanjis that contain the given radical.
func (r Radical) AllKanjisContaining() string {
	return string(radical2Kanji[rune(r)])
}
