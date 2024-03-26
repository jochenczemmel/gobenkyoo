package store_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/store"
)

func TestConfigurationError(t *testing.T) {
	want := "error message"
	err := store.ConfigurationError(want)
	got := err.Error()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}
