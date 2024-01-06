package learncards

// Box provides access to learning cards with different learn modes.
type Box struct {
	Title      string
	BookTitle  string
	modes      []string
	containers map[string]container
}

// NewBox returns an initialized Box with the provided titles.
func NewBox(title, booktitle string) Box {
	return Box{
		Title:      title,
		BookTitle:  booktitle,
		modes:      []string{},
		containers: map[string]container{},
	}
}

// Set fills a new container for the given mode with the given cards.
func (b *Box) Set(mode string, cards ...Card) {
	if _, ok := b.containers[mode]; !ok {
		b.modes = append(b.modes, mode)
	}
	b.containers[mode] = newContainer(cards...)
}

// Cards returns a list of cards in the given level for the given mode.
func (b Box) Cards(mode string, level int) []Card {
	return b.containers[mode].cards(level)
}

// NCards returns the number of cards in the given level for the given mode.
func (b Box) NCards(mode string, level int) int {
	return len(b.containers[mode].cards(level))
}

// SetCardLevel sets the level of the given card for the given mode.
func (b *Box) SetCardLevel(mode string, card Card, level int) {
	container := b.containers[mode]
	container.setLevel(card, level)
}
