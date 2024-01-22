// Package app provides entry functions to the application.
package app

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// App provides access to the application
type App struct {
	BookLoader BookLoader
	BoxLoader  BoxLoader
	// Importer   Importer
	// Storer     Storer

	library   *books.Library
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
		return fmt.Errorf("no box loader")
	}
	a.classroom, err = a.BoxLoader.LoadBoxes()
	return err
}
