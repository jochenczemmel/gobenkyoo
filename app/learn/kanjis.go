package learn

import (
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// Available kanji learning modes.
const (
	// ask kanji, answer native (and all other infos)
	Kanji2Native = "kanji_to_native"
	// ask native, answer kanji (and all other infos)
	Native2Kanji = "native_to_kanji"
	// ask kana (=spellings), answer kanji (and all other infos)
	Kana2Kanji = "kana_to_kanji"

	DefaultKanjiMode = Kanji2Native
)

// GetKanjiModes returns a list of the implemented (=valid) kanji learning modes.
func GetKanjiModes() []string {
	return []string{Kanji2Native, Native2Kanji, Kana2Kanji}
}

// makeKanjiCards transforms a list of words.Card to learn.Card
// using the given learn mode.
func makeKanjiCards(mode string, cards ...kanjis.Card) []Card {
	result := make([]Card, 0, len(cards))
	for _, card := range cards {
		result = append(result, makeKanjiCard(mode, card))
	}
	return result
}

// makeKanjiCard returns the learn card with the content of the
// card according to the given learn mode. If the mode is not
// known, the default mode is used.
func makeKanjiCard(mode string, card kanjis.Card) Card {
	result := Card{
		ID:          card.Kanji(),
		Hint:        card.Hint,
		Explanation: card.Explanation,
		// default mode is: Kanji2Native
		Question:    card.Kanji(),
		Answer:      strings.Join(card.Meanings(), ", "),
		MoreAnswers: []string{strings.Join(card.Readings(), ", ")},
	}

	kana := card.ReadingsKana()
	if len(kana) > 0 {
		// add kana readings if available
		result.MoreAnswers = append(result.MoreAnswers, strings.Join(kana, ", "))
	}

	switch mode {
	case Native2Kanji:
		result.Question = strings.Join(card.Meanings(), ", ")
		result.Answer = card.Kanji()
		result.MoreAnswers = []string{strings.Join(card.Readings(), ", ")}
		if len(kana) > 0 {
			// add kana readings if available
			result.MoreAnswers = append(result.MoreAnswers, strings.Join(kana, ", "))
		}
	case Kana2Kanji:
		result.Question = strings.Join(kana, ", ")
		if len(kana) == 0 {
			// use romaji if no kana readings available
			result.Question = strings.Join(card.Readings(), ", ")
		}
		result.Answer = card.Kanji()
		result.MoreAnswers = []string{strings.Join(card.Meanings(), ", ")}
		if len(kana) > 0 {
			result.MoreAnswers = append(result.MoreAnswers,
				// show romaji if question was in kana
				strings.Join(card.Readings(), ", "),
			)
		}
	}

	return result
}
