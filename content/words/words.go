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

// Card represents a single word or a sentence.
// If the word is a verb, the three forms should be filled.
type Card struct {
	// unique id of the card
	id string

	// content
	nihongo string // content as written in Japanese
	kana    string // content written in Kana
	romaji  string // content written in Romaji
	meaning string // meaning in the learners language

	// additional infos, might be empty
	hint        string // hint
	explanation string // explanantion
	contentType string // free text

	// only filled for verbs
	// nihongo contains the masu-form
	dictForm string // dictionary-form
	teForm   string // te-form
	naiForm  string // nai-form
}

// newCard returns a newCard word with the given id.
func newCard(id string) Card {
	return Card{id: id}
}

// IsEmpty checks if the it is an empty Content object.
// func (c Card) IsEmpty() bool {
// return c.id == ""
// }

// ID returns the id.
func (c Card) ID() string {
	return c.id
}

// Nihongo returns the Nihongo value.
func (c Card) Nihongo() string {
	return c.nihongo
}

// Kana returns the Kana value.
func (c Card) Kana() string {
	return c.kana
}

// Romaji returns the Romaji value.
func (c Card) Romaji() string {
	return c.romaji
}

// Meaning returns the Meaning.
func (c Card) Meaning() string {
	return c.meaning
}

// Hint returns the hint.
func (c Card) Hint() string {
	return c.hint
}

// Explanation returns the explanation.
func (c Card) Explanation() string {
	return c.explanation
}

// DictForm returns the dictionary form
// (verbs only).
func (c Card) DictForm() string {
	return c.dictForm
}

// TeForm returns the te-form (verbs only).
func (c Card) TeForm() string {
	return c.teForm
}

// GetNaiForm returns the nai-form (verbs only).
func (c Card) NaiForm() string {
	return c.naiForm
}

// ContentType returns the type of the content.
func (c Card) ContentType() string {
	return c.contentType
}
