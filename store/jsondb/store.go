package jsondb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// Storer provides storing content.
type Storer struct {
	path string
}

// NewStorer returns a storer that uses the given path.
func NewStorer(path string) Storer {
	return Storer{
		path: path,
	}
}

// StoreLibrary stores the specified library.
func (s Storer) StoreLibrary(library books.Library) error {
	dirName := filepath.Join(s.path, libraryPath)
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
	err = enc.Encode(library2json(library))
	if err != nil {
		return fmt.Errorf("store library: encode json: %w", err)
	}

	return nil
}

func library2json(library books.Library) Library {
	result := Library{
		Title: library.Title,
	}
	for _, book := range library.Books() {
		jsonBook := Book{
			Title:       book.TitleInfo.Title,
			SeriesTitle: book.TitleInfo.SeriesTitle,
			Volume:      book.TitleInfo.Volume,
		}
		jsonBook.LessonTitles = book.Lessons()
		jsonBook.LessonsByName = make(map[string]Lesson, len(jsonBook.LessonTitles))
		for _, lesson := range jsonBook.LessonTitles {
			jsonLesson := Lesson{
				Title:      lesson,
				KanjiCards: kanjiCards2Json(book.KanjisFor(lesson)...),
				WordCards:  wordCards2Json(book.WordsFor(lesson)...),
			}
			jsonBook.LessonsByName[lesson] = jsonLesson
		}
		result.Books = append(result.Books, jsonBook)
	}
	return result
}

func wordCards2Json(cards ...words.Card) []WordCard {
	result := make([]WordCard, 0, len(cards))
	for _, card := range cards {
		jsonCard := WordCard{
			ID:          card.ID,
			Nihongo:     card.Nihongo,
			Kana:        card.Kana,
			Romaji:      card.Romaji,
			Meaning:     card.Meaning,
			Hint:        card.Hint,
			Explanation: card.Explanation,
			DictForm:    card.DictForm,
			TeForm:      card.TeForm,
			NaiForm:     card.NaiForm,
		}
		result = append(result, jsonCard)
	}
	return result
}

func kanjiCards2Json(cards ...kanjis.Card) []KanjiCard {
	result := make([]KanjiCard, 0, len(cards))
	for _, card := range cards {
		jsonCard := KanjiCard{
			Kanji: card.String(),
		}
		for _, details := range card.GetDetails() {
			jsonDetail := KanjiDetail{
				Reading:     details.Reading,
				ReadingKana: details.ReadingKana,
				Meanings:    details.Meanings,
			}
			jsonCard.KanjiDetails = append(jsonCard.KanjiDetails, jsonDetail)
		}
		result = append(result, jsonCard)
	}
	return result
}
