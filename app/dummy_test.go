package app_test

import (
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// dummy is a test dummy ('test double').
// It implements some of the interfaces defined in the app package.
type dummy struct {
	loadError, pathError, storeError string
}

func (d dummy) LoadLibrary(string) (books.Library, error) {
	var result books.Library
	if d.loadError != "" {
		return result, fmt.Errorf("%s", d.loadError)
	}
	if d.pathError != "" {
		return result, &os.PathError{
			Op:   "open",
			Path: ".",
			Err:  os.ErrNotExist,
		}
	}
	return result, nil
}

func (d dummy) StoreLibrary(books.Library) error {
	if d.storeError != "" {
		return fmt.Errorf("%s", d.storeError)
	}
	return nil
}
