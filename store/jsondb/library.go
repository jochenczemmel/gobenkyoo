package jsondb

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

type Library struct {
	baseDir string
	library books.Library
}

func NewLibrary(dir string, library books.Library) Library {
	return Library{
		baseDir: dir,
		library: library,
	}
}

func (l Library) Store() error {
	dirName := filepath.Join(l.baseDir, libraryPath, url.PathEscape(l.library.Name))
	for _, book := range l.library.Books {
		err := storeBook(dirName, book)
		if err != nil {
			return fmt.Errorf("store library: %w", err)
		}
	}

	return nil
}
