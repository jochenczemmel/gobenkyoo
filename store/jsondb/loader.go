package jsondb

import (
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

type Loader struct {
	path string
}

func NewLoader(path string) *Loader {
	return &Loader{path: path}
}

func (n *Loader) LoadBooks() (*books.Library, error) {
	return nil, nil
}

func (n *Loader) LoadBoxes() (*learn.Library, error) {
	// read data
	// shelf := learn.NewShelf()
	return nil, nil
}
