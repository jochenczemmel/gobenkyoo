package learn

// Define card levels.
const (
	// InvalidLevel represents the level of a card
	// that is not found.
	InvalidLevel = -2
	// AllLevel represents cards on all Levels.
	// A card can never be on this level.
	AllLevel = -1
	// MinLevel is the default (intial) level for all cards
	// Reset should put the card back to this level.
	MinLevel = 0
	// MaxLevel is the final level.
	// Advance should not exceed this value.
	MaxLevel = 5
)

/*
// Levels returns the levels from MinLevel to MaxLevel
// (both included) as int list.
func Levels() []int {
	result := []int{}
	for i := MinLevel; i <= MaxLevel; i++ {
		result = append(result, i)
	}
	return result
}
*/
