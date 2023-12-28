package learn

type container struct {
	cardList []*Card
	levels   map[*Card]int
}

func newContainer(cards ...*Card) *container {
	result := &container{
		cardList: cards,
		levels:   make(map[*Card]int, len(cards)),
	}
	for _, card := range cards {
		card.CurrentLevel = MinLevel
		result.levels[card] = MinLevel
	}
	return result
}

func (c container) cards(level int) []*Card {
	result := []*Card{}
	for _, card := range c.cardList {
		if c.levels[card] == level {
			result = append(result, card)
		}
	}
	return result
}

func (c *container) setLevel(card *Card, level int) {
	// level too low
	if level < MinLevel {
		return
	}
	// level too high
	if level > MaxLevel {
		return
	}
	// card not in box
	if _, ok := c.levels[card]; !ok {
		return
	}
	card.CurrentLevel = level
	c.levels[card] = level
}
