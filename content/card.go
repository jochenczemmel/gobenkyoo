package content

// Card defines the minimal common behavior of kanji and word cards.
type Card interface {
	ID() string
}
