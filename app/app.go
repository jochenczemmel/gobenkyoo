package app

import (
	"errors"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

type Controller struct {
	lib           books.Library
	room          learn.Classroom
	Loader        DataLoader
	Storer        DataStorer
	KanjiImporter KanjiImporter
	WordImporter  WordImporter
}

func NewLoadStoreController(loadstorer DataLoadStorer) Controller {
	return Controller{
		Loader: loadstorer,
		Storer: loadstorer,
	}
}

// LoadLibrary loads the library with the given name.
// It returns the library and true if it is found.
// It returns a new library and false if the library is not found.
// In case of another error, the error is returned.
func (c *Controller) LoadLibrary(name string) (found bool, err error) {
	c.lib, err = c.Loader.LoadLibrary(name)
	if err == nil {
		return true, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		return false, nil
	}

	return false, err
}

func (c Controller) StoreLibrary() error {
	return c.Storer.StoreLibrary(c.lib)
}

func (c Controller) Book(title, seriestitle string, volume int) books.Book {
	return c.lib.Book(books.ID{
		Title:       title,
		SeriesTitle: seriestitle,
		Volume:      volume,
	})
}

func (c Controller) SetBooks(book books.Book) {
	c.lib.SetBooks(book)
}

func (c Controller) ImportKanji(filename string) error {
	return c.KanjiImporter.ImportKanji(filename)
}

func (c Controller) ImportWord(filename string) error {
	return c.WordImporter.ImportWord(filename)
}
