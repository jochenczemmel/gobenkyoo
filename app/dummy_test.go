package app_test

import (
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// dummy is a test dummy ('test double').
type dummy struct {
	loadError, pathError, storeError string
	kanjiError, wordError            string
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

func (d dummy) ImportKanji(string) ([]kanjis.Card, error) {
	var result []kanjis.Card
	if d.kanjiError != "" {
		return result, fmt.Errorf("%s", d.kanjiError)
	}
	return result, nil
}

func (d dummy) ImportWord(string) ([]words.Card, error) {
	var result []words.Card
	if d.wordError != "" {
		return result, fmt.Errorf("%s", d.wordError)
	}
	return result, nil
}
