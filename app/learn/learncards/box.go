package learncards

type Box struct {
	Title      string
	BookTitle  string
	modes      []string
	containers map[string]container
}

func NewBox(title, booktitle string) *Box {
	return &Box{
		Title:      title,
		BookTitle:  booktitle,
		modes:      []string{},
		containers: map[string]container{},
	}
}

func (b *Box) Set(mode string, cards ...Card) {
	if _, ok := b.containers[mode]; !ok {
		b.modes = append(b.modes, mode)
	}
	b.containers[mode] = newContainer(cards...)
}

func (b *Box) Modes() []string {
	return b.modes
}

func (b *Box) AllCards(mode string) []Card {
	container, ok := b.containers[mode]
	if !ok {
		return []Card{}
	}
	return container.cards(AllLevel)
}
