package learn

// Define the valid kanji learning modes.
const (
	// ask kanji, answer native (and all other infos)
	Kanji2Native = "kanji_to_native"
	// ask native, answer kanji (and all other infos)
	Native2Kanji = "native_to_kanji"
	// ask kana (=spellings), answer kanji (and all other infos)
	Kana2Kanji = "kana_to_kanji"

	KanjiType        = "kanji"
	DefaultKanjiMode = Kanji2Native
)

// GetKanjiModes returns a list of the implemented (=valid) kanji learning modes.
func GetKanjiModes() []string {
	return []string{Kanji2Native, Native2Kanji, Kana2Kanji}
}
