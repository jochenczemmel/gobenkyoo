package ui

import (
	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/ui/cli"
)

// New returns a runner for the requested ui type.
// If the ui type is unknown, it returns an UnknownRunner.
func New(uitype string, application *app.App) app.Runner {
	switch uitype {
	case UITypeCLILearn:
		return cli.NewLearner(application)
	case UITypeCLISearch:
		// not yet implemented
		return UnknownRunner{uitype: uitype}
	}

	return UnknownRunner{uitype: uitype}
}

// UnknownRunner is used to handle unknown application types.
type UnknownRunner struct {
	uitype string
}

// Run always returns an error.
func (u UnknownRunner) Run() error {
	return UnknownTypeError(u.uitype)
}

type UnknownTypeError string

func (e UnknownTypeError) Error() string {
	return "unknown ui type: " + string(e)
}
