package kanjis

// Builder provides functions to create a new kanjis Content object.
type Builder struct {
	kanjiCard *Card
}

// NewBuilder creates a new Content object.
func NewBuilder(kanji rune) *Builder {
	return &Builder{kanjiCard: newCard(kanji)}
}

// AddDetails createas and adds a Details object to the kanji.
func (b *Builder) AddDetails(reading string, meanings ...string) *Builder {
	b.kanjiCard.addDetails(newDetail(reading, meanings...))

	return b
}

// AddDetailsKana createas and adds a Details object
// with kana to the kanji.
func (b *Builder) AddDetailsKana(reading, kana string, meanings ...string) *Builder {
	b.kanjiCard.addDetails(newDetailKana(reading, kana, meanings...))

	return b
}

// Build returns the created Content object.
func (b Builder) Build() *Card {
	return b.kanjiCard
}
