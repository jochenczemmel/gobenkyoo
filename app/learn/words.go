package learn

import "github.com/jochenczemmel/gobenkyoo/content/words"

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

func NewWordBox(cards ...*words.Card) *Box {
	return &Box{}
}

func makeWordCard(mode string, card *words.Card) *Card {
	result := &Card{
		Hint:        card.Hint,
		Explanation: card.Explanation,
		WordCard:    card,
	}

	switch mode {
	case Native2Japanese:
		result.Question = card.Meaning
		result.Answer = append(result.Answer, card.Nihongo, card.Kana, card.Romaji)
	case Japanese2Native:
		result.Question = card.Nihongo
		result.Answer = append(result.Answer, card.Meaning, card.Kana, card.Romaji)
	case Native2Kana:
		result.Question = card.Meaning
		result.Answer = append(result.Answer, card.Kana, card.Romaji, card.Nihongo)
	case Kana2Native:
		result.Question = card.Kana
		result.Answer = append(result.Answer, card.Meaning, card.Romaji, card.Nihongo)
	default:
		return &Card{}
	}

	if card.DictForm != "" {
		result.Answer = append(result.Answer, card.DictForm, card.TeForm, card.NaiForm)
	}

	return result
}

/*
*
* 	Question    string
	Hint        string
	Answer      []string
	Explanation string

func makeWordCards(mode string, cards ...*words.Card) []*Card {
	result := make([]*Card, 0, len(cards))
	for _, card := range cards {
		result = append(result, makeWordCard(mode, card))
	}
	return result
}

*/
