// Package learn provides organzing cards in boxes,
// executing exams and tracking the learning process.
package learn

// Define card levels.
const (
	// InvalidLevel represents the level of a card
	// that is not found in the box.
	InvalidLevel = -2
	// AllLevel represents cards on all Levels.
	// A card can never be on this level.
	AllLevel = -1
	// MinLevel is the default level for all cards
	// Reset puts cards back to MinLevel
	MinLevel = 0
	// MaxLevel is the final level.
	// Advance lets the card in MaxLevel.
	MaxLevel = 5
)

// Levels returns the Levels from MinLevel to MaxLevel as int list.
func Levels() []int {
	result := []int{}
	for i := MinLevel; i <= MaxLevel; i++ {
		result = append(result, i)
	}
	return result
}
