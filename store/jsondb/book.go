package jsondb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"unicode/utf8"

	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

// storeBook stores a book as json file in the given directory.
func storeBook(dirname string, book books.Book) error {

	err := os.MkdirAll(dirname, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("store book: create directory: %w", err)
	}

	jsonBook := bookJSON{
		ID: bookIDJSON{
			Title:       book.ID.Title,
			SeriesTitle: book.ID.SeriesTitle,
			Volume:      book.ID.Volume,
		},
		LessonNames:   book.LessonNames(),
		LessonsByName: map[string]lessonJSON{},
	}

	for _, name := range jsonBook.LessonNames {
		lesson, ok := book.Lesson(name)
		if ok {
			jsonBook.LessonsByName[name] = lesson2json(lesson)
		}
	}

	fileName := filepath.Join(dirname, jsonBook.ID.fileName())
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("store book: create file: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if !Minify {
		enc.SetIndent("", "\t")
	}
	err = enc.Encode(jsonBook)
	if err != nil {
		return fmt.Errorf("store book %v: encode json: %w", jsonBook.ID, err)
	}

	return nil
}

// readBook reads a single book file and returrns the content.
func readBook(filename string) (books.Book, error) {
	var book books.Book
	file, err := os.Open(filename)
	if err != nil {
		return book, fmt.Errorf("open book file: %w", err)
	}
	defer file.Close()

	var jsonBook bookJSON
	err = json.NewDecoder(file).Decode(&jsonBook)
	if err != nil {
		return book, fmt.Errorf("json book %q: decode: %w", filename, err)
	}

	book = books.New(books.ID{
		Title:       jsonBook.ID.Title,
		SeriesTitle: jsonBook.ID.SeriesTitle,
		Volume:      jsonBook.ID.Volume,
	})

	for _, ln := range jsonBook.LessonNames {
		lesson := books.NewLesson(ln)
		lesson.AddKanjis(json2KanjiCards(jsonBook.LessonsByName[ln].KanjiCards)...)
		lesson.AddWords(json2WordCards(jsonBook.LessonsByName[ln].WordCards)...)
		book.SetLessons(lesson)
	}

	return book, nil
}

// lesson2json converts a book lesson to a jsondb lesson.
func lesson2json(lesson books.Lesson) lessonJSON {
	return lessonJSON{
		Name:       lesson.Name,
		KanjiCards: kanjiCards2Json(lesson.KanjiCards()),
		WordCards:  wordCards2Json(lesson.WordCards()),
	}
}

type bookJSON struct {
	ID            bookIDJSON            `json:"id"`
	LessonNames   []string              `json:"lessonNames,omitempty"`
	LessonsByName map[string]lessonJSON `json:"lessonsByName,omitempty"`
}

type bookIDJSON struct {
	Title       string `json:"title,omitempty"`
	SeriesTitle string `json:"seriesTitle,omitempty"`
	Volume      int    `json:"volume,omitempty"`
}

func (b bookIDJSON) fileName() string {
	return url.PathEscape(
		b.Title+"\n"+b.SeriesTitle+"\n"+strconv.Itoa(b.Volume)) + jsonExtension
}

type lessonJSON struct {
	Name       string          `json:"name"`
	WordCards  []wordCardJSON  `json:"wordCards,omitempty"`
	KanjiCards []kanjiCardJSON `json:"kanjiCards,omitempty"`
}

type wordCardJSON struct {
	ID          string `json:"id"`
	Nihongo     string `json:"nihongo,omitempty"`
	Kana        string `json:"kana,omitempty"`
	Romaji      string `json:"romaji,omitempty"`
	Meaning     string `json:"meaning"`
	Hint        string `json:"hint,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	DictForm    string `json:"dictForm,omitempty"`
	TeForm      string `json:"teForm,omitempty"`
	NaiForm     string `json:"naiForm,omitempty"`
}

type kanjiCardJSON struct {
	ID           string            `json:"id"`
	Kanji        string            `json:"kanji"`
	KanjiDetails []kanjiDetailJSON `json:"kanjiDetails"`
}

type kanjiDetailJSON struct {
	Reading     string   `json:"reading"`
	ReadingKana string   `json:"readingKana,omitempty"`
	Meanings    []string `json:"meanings"`
}

// kanjiCards2Json converts a list of kanji cards to json cards.
func kanjiCards2Json(cards []kanjis.Card) []kanjiCardJSON {
	result := make([]kanjiCardJSON, 0, len(cards))
	for _, card := range cards {
		jsonCard := kanjiCardJSON{
			ID:    card.ID,
			Kanji: card.String(),
		}
		for _, details := range card.Details {
			jsonDetail := kanjiDetailJSON{
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

// json2KanjiCards converts a list of cards from json cards.
func json2KanjiCards(jsoncards []kanjiCardJSON) []kanjis.Card {
	result := make([]kanjis.Card, 0, len(jsoncards))
	for _, jsoncard := range jsoncards {
		kanji, _ := utf8.DecodeRuneInString(jsoncard.Kanji)
		card := kanjis.Card{
			ID:    jsoncard.ID,
			Kanji: kanji,
		}
		for _, jsonDetail := range jsoncard.KanjiDetails {
			detail := kanjis.Detail{
				Reading:     jsonDetail.Reading,
				ReadingKana: jsonDetail.ReadingKana,
				Meanings:    jsonDetail.Meanings,
			}
			card.Details = append(card.Details, detail)
		}
		result = append(result, card)
	}

	return result
}

// wordCards2Json converts a list of word cards to json cards.
func wordCards2Json(cards []words.Card) []wordCardJSON {
	result := make([]wordCardJSON, 0, len(cards))
	for _, card := range cards {
		jsonCard := wordCardJSON{
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

// json2WordCards converts a list of json cards to word cards.
func json2WordCards(jsoncards []wordCardJSON) []words.Card {
	result := make([]words.Card, 0, len(jsoncards))

	for _, jsonCard := range jsoncards {
		card := words.Card{
			ID:          jsonCard.ID,
			Nihongo:     jsonCard.Nihongo,
			Kana:        jsonCard.Kana,
			Romaji:      jsonCard.Romaji,
			Meaning:     jsonCard.Meaning,
			Hint:        jsonCard.Hint,
			Explanation: jsonCard.Explanation,
			DictForm:    jsonCard.DictForm,
			TeForm:      jsonCard.TeForm,
			NaiForm:     jsonCard.NaiForm,
		}
		result = append(result, card)
	}

	return result
}
