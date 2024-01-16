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

type Loader struct {
	path string
}

func NewLoader(path string) Loader {
	return Loader{
		path: path,
	}
}

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
	library = books.NewLibrary(jsonLibrary.Title)
	for _, jsonBook := range jsonLibrary.Books {
		book := books.New(jsonBook.Title, jsonBook.SeriesTitle, jsonBook.Volume)
		for _, lessonTitle := range jsonBook.LessonTitles {
			book.AddKanjis(lessonTitle,
				json2KanjiCards(jsonBook.LessonsByName[lessonTitle].KanjiCards)...)
			book.AddWords(lessonTitle,
				json2WordCards(jsonBook.LessonsByName[lessonTitle].WordCards)...)
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
		kanji, _ := utf8.DecodeRuneInString(jsonCard.Kanji)
		builder := kanjis.NewBuilder(kanji)
		for _, details := range jsonCard.KanjiDetails {
			if details.ReadingKana != "" {
				builder.AddDetailsWithKana(
					details.Reading,
					details.ReadingKana,
					details.Meanings...,
				)
			} else {
				builder.AddDetails(details.Reading, details.Meanings...)
			}
		}
		result = append(result, builder.Build())
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
