package learn_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

func TestLevels(t *testing.T) {
	got := learn.Levels()
	want := []int{0, 1, 2, 3, 4, 5}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("ERROR: got- want+\n%s", diff)
	}
}
