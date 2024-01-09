package app

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// BookLoader defines a method for loading book data from the storage.
type BookLoader interface {
	LoadBooks() (*books.Library, error)
}

// BoxLoader defines a method for loading learn boxes from the storage.
type BoxLoader interface {
	LoadBoxes() (*learn.Library, error)
}

// Loader defines methods for loading data from the storage.
type Loader interface {
	BookLoader
	BoxLoader
}

// Runner defines a method to execute the application.
type Runner interface {
	Run() error
}

/*
// Storer defines methods for storing data in the storage.
type Storer interface {
	Store(*books.Library, *learn.Shelf) error
}

// Importer defines a method for importing data from external sources.
type Importer interface {
	Import() (*books.Library, error)
}

*/
