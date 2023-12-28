package learn_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordBox(t *testing.T) {

	inputCards := []*words.Card{{}}
	box := learn.NewWordBox(inputCards...)
	t.Logf("DEBUG: box: %v", box)

	// exam := box.StartExam(learn.Native2Japanese, learn.MinLevel)
}
