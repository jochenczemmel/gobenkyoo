package app

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// Loader defines methods for loading data from the storage
// or importing data from external sources.
type Loader interface {
	Load() (*books.Library, *learn.Shelf, error)
}

/*
// Storer defines methods for storing data in the storage.
type Storer interface {
	Store(*books.Library, *learn.Shelf) error
}

type Runner interface {
	Run() error
}
*/

// TODO: do we need separate interfaces?
// - for books.Library and learn.Shelf?
// - for Loader and Storer?
// Example:
// type BookLoader interface {
// LoadBooks() (*books.Library, error)
// }
