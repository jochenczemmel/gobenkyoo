package learn

import (
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Box provides access to learning cards with different learn modes.
type Box struct {
	BoxID             // id of the box
	Type       string // type of the box ("word" or "kanji")
	modes      []string
	containers map[string]container
}

// NewKanjiBox returns an initialized box with banji cards.
func NewKanjiBox(id BoxID, cards ...kanjis.Card) Box {
	box := Box{
		BoxID:      id,
		Type:       KanjiType,
		modes:      KanjiModes(),
		containers: map[string]container{},
	}
	for _, mode := range box.modes {
		box.containers[mode] = newContainer(
			box.makeKanjiCards(mode, cards...)...)
	}

	return box
}

// NewWordBox returns an initialized box with word cards.
// The cards are initially set to the minimum level on each learn mode.
func NewWordBox(id BoxID, cards ...words.Card) Box {
	box := Box{
		BoxID:      id,
		Type:       WordType,
		modes:      WordModes(),
		containers: map[string]container{},
	}
	for _, mode := range box.modes {
		box.containers[mode] = newContainer(
			box.makeWordCards(mode, cards...)...)
	}

	return box
}

// AddCards adds new cards to the given mode in the given level.
// If the mode does not match the box modes, nothing happens.
func (b *Box) AddCards(mode string, level int, cards ...Card) {
	container, ok := b.containers[mode]
	if !ok {
		return
	}
	container.addCards(level, cards...)
	b.containers[mode] = container
}

// Cards returns a list of cards in the given level for the given mode.
// The cards are initially set to the minimum level on each learn mode.
func (b Box) Cards(mode string, level int) []Card {
	return b.containers[mode].cards(level)
}

// NCards returns the number of cards in the given level for the given mode.
func (b Box) NCards(mode string, level int) int {
	return len(b.Cards(mode, level))
}

// Modes returns the list of learn modes for the box.
func (b Box) Modes() []string {
	if b.Type == KanjiType {
		return KanjiModes()
	}
	return WordModes()
}

// SetCardLevel moves an existing Card to the specified level
// in the specified learn mode.
func (b *Box) SetCardLevel(mode string, level int, cards ...Card) {
	container, ok := b.containers[mode]
	if !ok {
		return
	}
	for _, card := range cards {
		container.setCardLevel(level, card)
	}
}
