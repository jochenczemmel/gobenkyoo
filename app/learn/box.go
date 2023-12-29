package learn

import "github.com/jochenczemmel/gobenkyoo/content/words"

// Box provides organizing a set of cards in levels.
// Each card has multiple levels, one for each learn mode.
type Box struct {
	containers map[string]*container
}

// NewWordBox creates a Box from a list of word cards.
func NewWordBox(cards ...*words.Card) *Box {
	result := &Box{containers: make(map[string]*container)}
	for _, mode := range GetWordModes() {
		result.containers[mode] = newContainer(makeWordCards(mode, cards...)...)
	}
	return result
}

// Cards returns a sorted list of cards in the given learn mode
// that match the given level.
// If the learn mode is not known, an empty list is returned.
func (b *Box) Cards(mode string, level int) []*Card {
	container, ok := b.containers[mode]
	if !ok {
		return []*Card{}
	}
	return container.cards(level)
}

// Advance shifts the card on the next level for the given learning mode.
func (b *Box) Advance(mode string, card *Card) {
	container, ok := b.containers[mode]
	if !ok {
		return
	}
	container.advance(card)
}

// Reset puts the card on the start level for the given learning mode.
func (b *Box) Reset(mode string, card *Card) {
	container, ok := b.containers[mode]
	if !ok {
		return
	}
	container.setLevel(card, MinLevel)
}
