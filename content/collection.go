package content

// Collection represents a set of word.Content objects.
type Collection[T Card] struct {
	Title         string
	cardList      []T
	uniqueContent map[string]T
}

// NewCollection returns a new Collection object with the given titles
// and with the optionally provided Cards.
func NewCollection[T Card](title string, cards ...T) Collection[T] {
	result := Collection[T]{
		Title:         title,
		uniqueContent: map[string]T{},
	}
	result.Add(cards...)
	return result
}

// Add adds new Content to the collection.
// If the id exists already in the Collection, it is not added.
func (c *Collection[T]) Add(cards ...T) {
	for _, card := range cards {
		if _, ok := c.uniqueContent[card.ID()]; ok {
			continue
		}
		c.cardList = append(c.cardList, card)
		c.uniqueContent[card.ID()] = card
	}
}

// Content returns the content of the Collection as a list.
func (c Collection[T]) Content() []T {
	return c.cardList
}

// Get returns the content with the provided id.
func (c Collection[T]) Get(id string) T {
	return c.uniqueContent[id]
}
