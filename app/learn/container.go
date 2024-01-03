package learn

// container provides management of a set of cards
// for a single learn mode.
// Each card has an level between MinLevel and MaxLevel (including).
// The sort order of the cards is preserved.
type container struct {
	cardList []Card
	levels   map[string]int
}

// newContainer returns a container with the cards.
// All cards are stored in level MinLevel.
func newContainer(cards ...Card) *container {
	result := &container{
		cardList: cards,
		levels:   make(map[string]int, len(cards)),
	}
	for _, card := range cards {
		result.levels[card.Identity] = MinLevel
	}
	return result
}

// cards returns a sorted list of cards that match the given level.
// If level is AllLevel, all cards are returned.
func (c container) cards(level int) []Card {
	result := []Card{}
	for _, card := range c.cardList {
		if c.levels[card.Identity] == level || level == AllLevel {
			result = append(result, card)
		}
	}
	return result
}

// setLevel sets the level for the given card.
// If the level is lower than MinLevel, larger than MaxLevel
// or if the card is unknown, nothing happens.
func (c *container) setLevel(card Card, level int) {
	// level too low
	if level < MinLevel {
		return
	}
	// level too high
	if level > MaxLevel {
		return
	}
	// card not in box
	if _, ok := c.levels[card.Identity]; !ok {
		return
	}
	c.levels[card.Identity] = level
}

// advance puts the card in the next level.
// If it is already in the highest level or if it is
// not known, nothing happens.
func (c *container) advance(card Card) {
	// card not in box
	if _, ok := c.levels[card.Identity]; !ok {
		return
	}
	level := c.levels[card.Identity]
	// level too high
	if level >= MaxLevel {
		return
	}
	c.levels[card.Identity] = level + 1
}
