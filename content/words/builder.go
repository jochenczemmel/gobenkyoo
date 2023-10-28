package words

// Builder provides functions to build a words Card object.
type Builder struct {
	word Card
}

// NewBuilder creates a new words Card object.
func NewBuilder(id string) *Builder {
	return &Builder{
		word: newCard(id),
	}
}

// Build returns the new words Card object.
func (b Builder) Build() Card {
	return b.word
}

// SetContent sets content to the new word.
func (b *Builder) SetContent(nihongo, kana, romaji, meaning string) *Builder {

	b.word.nihongo = nihongo
	b.word.kana = kana
	b.word.romaji = romaji
	b.word.meaning = meaning

	return b
}

// SetHint sets a hint to the new word.
func (b *Builder) SetHint(hint string) *Builder {
	b.word.hint = hint

	return b
}

// SetExplanation sets an explanation to the new word.
func (b *Builder) SetExplanation(explanation string) *Builder {
	b.word.explanation = explanation

	return b
}

// SetVerbForms sets the dictionary form,
// the te-form and the nai-form of the verb.
func (b *Builder) SetVerbForms(dictform, teform, naiform string) *Builder {
	b.word.dictForm = dictform
	b.word.teForm = teform
	b.word.naiForm = naiform

	return b
}

// SetContentType sets the content type.
func (b *Builder) SetContentType(newtype string) *Builder {
	b.word.contentType = newtype

	return b
}
