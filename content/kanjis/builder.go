package kanjis

// Builder provides functions to create a new kanji card.
type Builder struct {
	kanjiCard Card
}

// NewBuilder returns a builder for the given kanji.
func NewBuilder(kanji rune) *Builder {
	return &Builder{kanjiCard: newCard(kanji)}
}

// AddDetails adds a single reading (in romaji) and the
// associated meanings the to kanji.
func (b *Builder) AddDetails(reading string, meanings ...string) *Builder {
	b.kanjiCard.addDetails(newDetail(reading, meanings...))

	return b
}

// AddDetailsWithKana adds a single reading (in romaji and kana) and  the
// associated meanings the to kanji.
func (b *Builder) AddDetailsWithKana(reading, kana string, meanings ...string) *Builder {
	b.kanjiCard.addDetails(newDetailWithlKana(reading, kana, meanings...))

	return b
}

// Build returns the created kanji card.
func (b Builder) Build() Card {
	return b.kanjiCard
}
