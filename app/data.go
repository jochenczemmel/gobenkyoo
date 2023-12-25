package app

import "github.com/jochenczemmel/gobenkyoo/content/books"

type Data struct {
	library books.Library
	// shelf learn.Shelf
}

// Loader defines methods for loading data from the storage
// or importing data from external sources.
type Loader interface {
	Load() (*Data, error)
}

// Storer defines methods for storing data in the storage.
type Storer interface {
	Store(*Data) error
}
