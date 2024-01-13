package jsondb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/content/books"
)

const (
	libararyPath           = "library"
	jsonExtension          = ".json"
	defaultFilePermissions = 0750
)

type Storer struct {
	path string
}

func NewStorer(path string) Storer {
	return Storer{
		path: path,
	}
}

func (s Storer) StoreLibrary(library books.Library) error {
	dirName := filepath.Join(s.path, libararyPath)
	err := os.MkdirAll(dirName, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("store library: create directory: %w", err)
	}

	fileName := filepath.Join(dirName, library.Title+jsonExtension)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("store library: create file: %w", err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	err = enc.Encode(converLibrary(library))
	if err != nil {
		return fmt.Errorf("store library: encode json: %w", err)
	}
	return nil
}

func converLibrary(library books.Library) Library {
	result := Library{
		Title: library.Title,
	}
	for _, book := range library.Books() {
		jsonBook := Book{
			TitleInfo: book.TitleInfo,
		}
		result.Books = append(result.Books, jsonBook)
	}
	return result
}

type Library struct {
	Title string `json:",omitempty"`
	Books []Book `json:",omitempty"`
}

type Book struct {
	books.TitleInfo
	LessonTitles  []string          `json:",omitempty"`
	LessonsByName map[string]Lesson `json:",omitempty"`
}

type Lesson struct {
	Title      string
	WordCards  []WordCard  `json:",omitempty"`
	KanjiCards []KanjiCard `json:",omitempty"`
}

type KanjiCard struct {
}

type WordCard struct {
}
