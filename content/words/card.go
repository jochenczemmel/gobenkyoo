// Package words handles vocabulary data.
// It handles single words, expressions, complete sentences
// or short dialogs.
package words

// A Card object contains the japanese meaning
//   - as romaji,
//   - kana (hiragana or katakana) and/or
//   - nihongo (kanji, kana, and romaji)
//
// and the meaning in the learners native language.
//
// It can contain a hint and an explanation.
// In the learning process, the hint is supposed to be presented
// with the question, the explanation is supposed to be presented
// with the answer. This hint can be used to add precision to the question,
// the explanation can be used to display additional variants,
// special cases, caveats etc.
//
// If the word is a verb, the three forms should be filled
// with the dictionary form, the -te-form and the -nai-form.
// The variable Nihongo contains the -masu-form
//
// To create a Card, the words.Builder can be used.
type Card struct {
	ID      string // unique identifier of the card
	Nihongo string // content as written in Japanese
	Kana    string // content written in Kana
	Romaji  string // content written in Romaji
	Meaning string // meaning in the learners language

	// additional infos, might be empty
	Hint        string // hint
	Explanation string // explanation

	// only filled for verbs
	DictForm string // dictionary-form
	TeForm   string // te-form
	NaiForm  string // nai-form
}
