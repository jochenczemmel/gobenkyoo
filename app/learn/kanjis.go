package learn

import (
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

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

// makeKanjiCards transforms a list of words.Card to learn.Card
// using the given learn mode.
func makeKanjiCards(mode string, cards ...kanjis.Card) []Card {
	result := make([]Card, 0, len(cards))
	for _, card := range cards {
		result = append(result, makeKanjiCard(mode, card))
	}
	return result
}

// makeKanjiCard returns the learn.Card with the content of the
// card according to the given learn mode. If the mode is not
// known, an empty card is returned.
func makeKanjiCard(mode string, card kanjis.Card) Card {
	result := Card{
		ID:          card.Kanji(),
		Hint:        card.Hint,
		Explanation: card.Explanation,
	}

	switch mode {
	case Kanji2Native:
		result.Question = card.Kanji()
		result.Answer = strings.Join(card.Meanings(), ", ")
		result.MoreAnswers = append(result.MoreAnswers,
			strings.Join(card.Readings(), ", "),
			strings.Join(card.ReadingsKana(), ", "),
		)

	case Native2Kanji:
		result.Question = strings.Join(card.Meanings(), ", ")
		result.Answer = card.Kanji()
		result.MoreAnswers = append(result.MoreAnswers,
			strings.Join(card.Readings(), ", "),
			strings.Join(card.ReadingsKana(), ", "),
		)
	case Kana2Kanji:
		result.Question = strings.Join(card.ReadingsKana(), ", ")
		result.Answer = card.Kanji()
		result.MoreAnswers = append(result.MoreAnswers,
			strings.Join(card.Meanings(), ", "),
			strings.Join(card.Readings(), ", "),
		)
	default:
		return emptyCard
	}

	return result
}
