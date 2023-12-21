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
package words

// Card represents a single word or a sentence.
// If the word is a verb, the three forms should be filled.
// To create a Card, the words.Builder can be used.
type Card struct {

	// content
	Nihongo string // content as written in Japanese
	Kana    string // content written in Kana
	Romaji  string // content written in Romaji
	Meaning string // meaning in the learners language

	// additional infos, might be empty
	Hint        string // hint
	Explanation string // explanation
	ContentInfo string // free text

	// only filled for verbs
	// nihongo contains the masu-form
	DictForm string // dictionary-form
	TeForm   string // te-form
	NaiForm  string // nai-form
}
