package jsondb

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// Library provides loading and storing book libraries in json format.
type Library struct {
	baseDir string
}

func NewLibrary(dir string) Library {
	return Library{
		baseDir: dir,
	}
}

// Store stores the specified library in the base dir of the json library object.
func (l Library) Store(library books.Library) error {
	dirName := filepath.Join(l.baseDir, libraryPath, url.PathEscape(library.Name))
	for _, book := range library.Books {
		err := storeBook(dirName, book)
		if err != nil {
			return fmt.Errorf("store library: %w", err)
		}
	}

	return nil
}

// Store loads the specified library from the base dir of the json library object.
func (l Library) Load(name string) (books.Library, error) {
	dirName := filepath.Join(l.baseDir, libraryPath, url.PathEscape(name))
	library, err := readLibrary(name, dirName)
	if err != nil {
		return library, fmt.Errorf("load library: %w", err)
	}

	return library, nil
}

// readLibrary reads the json book files from the library directory.
func readLibrary(name, dirname string) (books.Library, error) {
	library := books.NewLibrary(name)

	dir, err := os.Open(dirname)
	if err != nil {
		return library, fmt.Errorf("open library directory: %w", err)
	}
	defer dir.Close()

	errorList := []error{}
	files, err := dir.ReadDir(readAllFiles)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), jsonExtension) {
			continue
		}
		book, err := readBook(filepath.Join(dirname, file.Name()))
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		library.AddBooks(book)
	}

	if len(errorList) > 0 {
		return library, fmt.Errorf("read book: %w", errors.Join(errorList...))
	}
	return library, nil
}
