package ui

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/ui/cli"
)

// New returns a runner for the requested ui type.
// If the ui type is unknown, it returns an UnknownRunner.
func New(uitype string, application *app.App) app.Runner {
	switch uitype {
	case UITypeCLILearn:
		return cli.NewLearner(application)
	}

	return UnknownRunner{uitype: uitype}
}

// UnknownRunner is used to handle unknown application types.
type UnknownRunner struct {
	uitype string
}

// Run always returns an error.
func (n UnknownRunner) Run() error {
	return fmt.Errorf("unknown ui type: %q", n.uitype)
}
