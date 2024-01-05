package learncards

// Modes returns a unique list of the known learn modes.
// Used as a test spy.
func (b *Box) Modes() []string {
	return b.modes
}
