// Package store handles importing, reading and storing of all
// kinds of application content.
package store

/*
import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app"
)

// NewLoader returns a loader for the requested database type.
func NewLoader(dbtype, path string) app.Loader {
	switch dbtype {
	case "json":
		return jsdondb.NewLoader(path)
	}

	return UnknownLoader{dbtype: dbtype}
}

// UnknownLoader is used to handle unknown
type UnknownLoader struct {
	dbtype string
}

// Load always returns an error.
func (n UnknownLoader) Load() (*app.Data, error) {
	return nil, fmt.Errorf("unknown db type: %q", n.dbtype)
}
*/
