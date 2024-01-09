package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Define the valid word learning modes.
const (
	// ask native, answer in japanese
	Native2Japanese = "native_to_japanese"
	// ask japanese, answer in native
	Japanese2Native = "japanese_to_native"
	// ask native, answer in kana
	Native2Kana = "native_to_kana"
	// ask kana, answer in native
	Kana2Native = "kana_to_native"

	WordType        = "word"
	DefaultWordMode = Native2Japanese
)

// GetWordModes returns a list of the implemented (=valid) word learning modes.
func GetWordModes() []string {
	return []string{
		Native2Japanese,
		Japanese2Native,
		Native2Kana,
		Kana2Native,
	}
}

// makeWordCards transforms a list of word card to a learn card
// using the given learn mode.
func makeWordCards(mode string, cards ...words.Card) []Card {
	result := make([]Card, 0, len(cards))
	for _, card := range cards {
		result = append(result, makeWordCard(mode, card))
	}
	return result
}

// makeWordCard returns the learn.Card with the content of the
// card accordig to the given learn mode. If the mode is not
// known, the DefaultWordMode is used.
func makeWordCard(mode string, card words.Card) Card {
	result := Card{
		ID:          card.ID,
		Hint:        card.Hint,
		Explanation: card.Explanation,
		// Default is: Native2Japanese
		Question:    card.Meaning,
		Answer:      card.Nihongo,
		MoreAnswers: []string{card.Kana, card.Romaji},
	}

	switch mode {
	case Japanese2Native:
		result.Question = card.Nihongo
		result.Answer = card.Meaning
		result.MoreAnswers = []string{card.Kana, card.Romaji}
	case Native2Kana:
		result.Question = card.Meaning
		result.Answer = card.Kana
		result.MoreAnswers = []string{card.Romaji, card.Nihongo}
	case Kana2Native:
		result.Question = card.Kana
		result.Answer = card.Meaning
		result.MoreAnswers = []string{card.Romaji, card.Nihongo}
	}

	if card.DictForm != "" {
		result.MoreAnswers = append(result.MoreAnswers,
			card.DictForm, card.TeForm, card.NaiForm)
	}

	return result
}
