package jsondb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Loader provides loading content.
type Loader struct {
	path string
}

// NewLoader returns a loader that uses the given base path.
func NewLoader(path string) Loader {
	return Loader{
		path: path,
	}
}

// LoadLibraries loads all libraries that are stored in the provided location.
func (l Loader) LoadLibraries() ([]books.Library, error) {
	result := []books.Library{}
	fileNames, err := fileList(l.path)
	if err != nil {
		return result, fmt.Errorf("load library: list files: %w", err)
	}

	for _, name := range fileNames {
		lib, err := l.loadLibrary(name)
		if err != nil {
			return result, fmt.Errorf("load library: %w", err)
		}
		result = append(result, lib)
	}

	return result, nil
}

func (l Loader) loadLibrary(name string) (books.Library, error) {

	var library books.Library
	file, err := os.Open(filepath.Join(l.path, name))
	if err != nil {
		return library, fmt.Errorf("open library file: %w", err)
	}
	defer file.Close()

	var jsonLibrary Library
	err = json.NewDecoder(file).Decode(&jsonLibrary)
	if err != nil {
		return library, fmt.Errorf("json decode: %w", err)
	}
	library = books.NewLibrary(jsonLibrary.Name)
	for _, jsonBook := range jsonLibrary.Books {
		book := books.New(books.NewID(jsonBook.Title, jsonBook.SeriesTitle, jsonBook.Volume))
		for _, lessonName := range jsonBook.LessonNames {
			book.AddKanjis(lessonName,
				json2KanjiCards(jsonBook.LessonsByName[lessonName].KanjiCards)...)
			book.AddWords(lessonName,
				json2WordCards(jsonBook.LessonsByName[lessonName].WordCards)...)
		}
		library.AddBooks(book)
	}

	return library, nil
}

func json2WordCards(jsoncards []WordCard) []words.Card {
	result := []words.Card{}
	for _, jsoncard := range jsoncards {
		result = append(result, words.Card{
			ID:          jsoncard.ID,
			Nihongo:     jsoncard.Nihongo,
			Kana:        jsoncard.Kana,
			Romaji:      jsoncard.Romaji,
			Meaning:     jsoncard.Meaning,
			Hint:        jsoncard.Hint,
			Explanation: jsoncard.Explanation,
			DictForm:    jsoncard.DictForm,
			TeForm:      jsoncard.TeForm,
			NaiForm:     jsoncard.NaiForm,
		})
	}

	return result
}

func json2KanjiCards(jsoncards []KanjiCard) []kanjis.Card {
	result := []kanjis.Card{}
	for _, jsonCard := range jsoncards {
		kanjiRune, _ := utf8.DecodeRuneInString(jsonCard.Kanji)
		card := kanjis.Card{Kanji: kanjiRune}
		for _, details := range jsonCard.KanjiDetails {
			card.Details = append(card.Details, kanjis.Detail{
				Reading:     details.Reading,
				ReadingKana: details.ReadingKana,
				Meanings:    details.Meanings,
			})
		}
		result = append(result, card)
	}

	return result
}

func fileList(path string) ([]string, error) {
	result := []string{}
	dir, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("open dir: %w", err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(readAllFiles)
	if err != nil {
		return result, fmt.Errorf("read dir: %w", err)
	}

	for _, name := range names {
		if strings.HasSuffix(name, jsonExtension) {
			result = append(result, name)
		}
	}

	return result, nil
}
