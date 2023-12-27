package store

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

// NewLoader returns a loader for the requested database type.
// If the dbtype is unknown, it returns an UnknownLoader.
func NewLoader(dbtype, path string) app.Loader {
	switch dbtype {
	case dbTypeJson:
		return jsondb.NewLoader(path)
	}

	return UnknownLoader{dbtype: dbtype}
}

// UnknownLoader is used to handle unknown database types.
type UnknownLoader struct {
	dbtype string
}

// LoadBooks always returns an error.
func (n UnknownLoader) LoadBooks() (*books.Library, error) {
	return nil, fmt.Errorf("unknown db type: %q", n.dbtype)
}

// LoadBoxes always returns an error.
func (n UnknownLoader) LoadBoxes() (*learn.Shelf, error) {
	return nil, fmt.Errorf("unknown db type: %q", n.dbtype)
}
