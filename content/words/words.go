// Package words handles vocabulary data.
//
// A single Card object contains the japanese meaning
// as romaji, kana (hiragana or Katakana) and nihongo (Kanji, Kana, Romaji).
// Additionally it contains the meaning in the learners language (meaning).
//
// It can contain a hint and an explanation.
// In the learning process, the hint is supposed to be presented
// with the question, the explanation is supposed to be presented
// with the answer. This hint can be used to add precision to the quest,
// the explanation can be used to display additional variants,
// special cases, caveats etc.
//
// Each entry has a unique id.
package words

import "fmt"

// Card represents a single word or a sentence.
// If the word is a verb, the three forms should be filled.
// To create a Card, the words.Builder can be used.
type Card struct {
	// unique Identifier of the card
	Identifier string

	// content
	Nihongo string // content as written in Japanese
	Kana    string // content written in Kana
	Romaji  string // content written in Romaji
	Meaning string // meaning in the learners language

	// additional infos, might be empty
	Hint        string // hint
	Explanation string // explanantion
	ContentType string // free text

	// only filled for verbs
	// nihongo contains the masu-form
	DictForm string // dictionary-form
	TeForm   string // te-form
	NaiForm  string // nai-form
}

// New returns a New word with the given id.
func New(id string) Card {
	return Card{Identifier: id}
}

// ID returns the id.
func (c Card) ID() string {
	return c.Identifier
}

// String returns the id.
func (c Card) String() string {
	if c.Identifier == "" {
		return ""
	}

	return fmt.Sprintf("%s: %q", c.Identifier, c.Nihongo)
}

// IsEmpty returns true if it is an empty object.
// func (c Card) IsEmpty() bool {
// return c.Identifier == ""
// }
