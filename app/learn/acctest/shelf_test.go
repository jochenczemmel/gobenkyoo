//go:build acctest

package acctest

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func TestLearnShelf(t *testing.T) {
	// TODO: execute full test from shelf
	shelf := learn.NewShelf()
	boxtitle := "box1"
	mode := "native_to_japanese"
	level := 1
	exam := shelf.StartWordExam(mode, level, boxtitle)

	card, ok := exam.NextCard()

	t.Logf("DEBUG: card: %v", card)
	t.Logf("DEBUG: ok: %v", ok)

	exam.Advance()
	exam.Reset()

	err := exam.Save()
	t.Logf("DEBUG: err: %v", err)
}
