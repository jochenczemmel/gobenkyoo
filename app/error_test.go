package app_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app"
)

func TestConfigurationError(t *testing.T) {
	want := "error message"
	err := app.ConfigurationError(want)
	got := err.Error()
	if got != want {
		t.Errorf("ERROR: got %v, want %v", got, want)
	}
}
