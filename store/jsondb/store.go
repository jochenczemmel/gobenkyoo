package jsondb

import "github.com/jochenczemmel/gobenkyoo/content/books"

type Storer struct {
	path string
}

func NewStorer(path string) Storer {
	return Storer{
		path: path,
	}
}

func (s Storer) StoreBook(book books.Book) error {
	return nil
}
