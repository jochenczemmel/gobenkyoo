package learncards

// Modes returns a unique list of the known learn modes.
// Only used in testing, used as a spy.
func (b *Box) Modes() []string {
	return b.modes
}
