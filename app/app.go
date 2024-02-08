// Package app provides entry functions to the application.
package app

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

// App provides access to the application.
type App struct {
	BookLoader BookLoader
	BoxLoader  BoxLoader

	// library   *books.Library
	classroom *learn.Classroom
}

// New returns a configured App object.
func New(loader Loader) *App {
	return &App{
		BookLoader: loader,
		BoxLoader:  loader,
	}
}

// LoadBoxes loads the learn box data from the storage.
func (a *App) LoadBoxes() (err error) {
	if a.BoxLoader == nil {
		return MissingLoaderError("no box loader")
	}
	a.classroom, err = a.BoxLoader.LoadBoxes()

	return fmt.Errorf("load boxes: %w", err)
}

type MissingLoaderError string

func (e MissingLoaderError) Error() string {
	return string(e)
}
