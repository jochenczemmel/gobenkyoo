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

// StoreLibrary stores all the books that are in the specified library.
func (l DB) StoreLibrary(library books.Library) error {
	dirName := filepath.Join(l.baseDir, libraryPath, url.PathEscape(library.Name))
	for _, id := range library.SortedBookIDs() {
		book := library.Book(id)
		err := storeBook(dirName, book)
		if err != nil {
			return fmt.Errorf("store library: %w", err)
		}
	}

	return nil
}

// LoadLibrary loads all the books in the specified library.
// If the library is not found, found is set to false.
func (l DB) LoadLibrary(name string) (lib books.Library, found bool, err error) {
	dirName := filepath.Join(l.baseDir, libraryPath, url.PathEscape(name))
	lib, err = readLibrary(name, dirName)
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
			return lib, false, err
		}
		return lib, true, fmt.Errorf("load library: %w", err)
	}

	return lib, true, nil
}

// readLibrary reads the json book files from the given directory
// and returns a library with the provided name.
func readLibrary(name, dirname string) (books.Library, error) {
	library := books.NewLibrary(name)

	dir, err := os.Open(dirname)
	if err != nil {
		return library, fmt.Errorf("open library directory: %w", err)
	}
	defer dir.Close()

	errorList := []error{}
	files, err := dir.ReadDir(readAllFiles)
	if err != nil {
		return library, fmt.Errorf("read directory files: %w", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), jsonExtension) {
			continue
		}
		book, err := readBook(filepath.Join(dirname, file.Name()))
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		library.SetBooks(book)
	}

	if len(errorList) > 0 {
		return library, fmt.Errorf("read book: %w", errors.Join(errorList...))
	}

	return library, nil
}
