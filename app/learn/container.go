package learn

// container provides management of a set of cards
// for a single learn mode.
// Each card has an level between MinLevel and MaxLevel (including).
// The sort order of the cards is preserved.
type container struct {
	cardList []Card
	levels   map[CardID]int
}

// newContainer returns a container with the cards.
// All cards are stored in level MinLevel.
func newContainer(cards ...Card) container {
	result := container{levels: make(map[CardID]int, len(cards))}
	result.addCards(MinLevel, cards...)
	return result
}

// cards returns a sorted list of cards that match the given level.
// If level is AllLevel, all cards are returned.
func (c container) cards(level int) []Card {
	result := []Card{}
	for _, card := range c.cardList {
		if c.levels[card.ID] == level || level == AllLevel {
			result = append(result, card)
		}
	}

	return result
}

// addCards adds the cards at the specified level.
// If the card is already known, it is set to the new level.
// If it is not known, it is added to the list.
// If the level is lower than MinLevel or larger than MaxLevel,
// the value is adjusted to MinLevel respectively MaxLevel.
func (c *container) addCards(level int, cards ...Card) {
	for _, card := range cards {
		if _, ok := c.levels[card.ID]; !ok {
			c.cardList = append(c.cardList, card)
		}
		c.setLevel(level, card)
	}
}

// advance puts the card in the next level.
func (c *container) advance(card Card) {
	c.setCardLevel(c.levels[card.ID]+1, card)
}

// setCardLevel sets the level for the given card.
// If the card is unknown, nothing happens.
// If the level is lower than MinLevel or larger than MaxLevel,
// the value is adjusted to MinLevel respectively MaxLevel.
func (c *container) setCardLevel(level int, card Card) {
	if _, ok := c.levels[card.ID]; !ok {
		// card not in box
		return
	}
	c.setLevel(level, card)
}

// setLevel setzts the card to the adjusted level.
func (c *container) setLevel(level int, card Card) {
	level = adjustLevel(level)
	c.levels[card.ID] = level
}

// adjustLevel adjusts the given level to be within
// the lower and upper limit.
func adjustLevel(level int) int {
	// level too low
	if level < MinLevel {
		return MinLevel
	}
	// level too high
	if level > MaxLevel {
		return MaxLevel
	}
	return level
}
