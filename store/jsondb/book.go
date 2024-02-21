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
func storeBook(dir string, book books.Book) error {

	err := os.MkdirAll(dir, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("store book: create directory: %w", err)
	}

	jsonBook := Book{
		ID: BookID{
			Title:       book.ID.Title,
			SeriesTitle: book.ID.SeriesTitle,
			Volume:      book.ID.Volume,
		},
		LessonNames:   book.LessonNames(),
		LessonsByName: map[string]Lesson{},
	}

	for _, name := range jsonBook.LessonNames {
		lesson, ok := book.Lesson(name)
		if ok {
			jsonBook.LessonsByName[name] = lesson2json(lesson)
		}
	}

	fileName := filepath.Join(dir, jsonBook.ID.fileName())
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("store book: create file: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
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

	var jsonBook Book
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
func lesson2json(lesson books.Lesson) Lesson {
	return Lesson{
		Name:       lesson.Name,
		KanjiCards: kanjiCards2Json(lesson.KanjiCards()),
		WordCards:  wordCards2Json(lesson.WordCards()),
	}
}

type Book struct {
	ID            BookID            `json:"id"`
	LessonNames   []string          `json:"lessonNames,omitempty"`
	LessonsByName map[string]Lesson `json:"lessonsByName,omitempty"`
}

type BookID struct {
	Title       string `json:"title,omitempty"`
	SeriesTitle string `json:"seriesTitle,omitempty"`
	Volume      int    `json:"volume,omitempty"`
}

func (b BookID) fileName() string {
	return url.PathEscape(
		b.Title+"\n"+b.SeriesTitle+"\n"+strconv.Itoa(b.Volume)) + jsonExtension
}

type Lesson struct {
	Name       string      `json:"name"`
	WordCards  []WordCard  `json:"wordCards,omitempty"`
	KanjiCards []KanjiCard `json:"kanjiCards,omitempty"`
	// BookID     BookID      `json:"id"`
}

type WordCard struct {
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

type KanjiCard struct {
	ID           string        `json:"id"`
	Kanji        string        `json:"kanji"`
	KanjiDetails []KanjiDetail `json:"kanjiDetails"`
}

type KanjiDetail struct {
	Reading     string   `json:"reading"`
	ReadingKana string   `json:"readingKana,omitempty"`
	Meanings    []string `json:"meanings"`
}

// kanjiCards2Json converts a list of kanjis.Card to json.KanjiCard.
func kanjiCards2Json(cards []kanjis.Card) []KanjiCard {
	result := make([]KanjiCard, 0, len(cards))
	for _, card := range cards {
		jsonCard := KanjiCard{
			Kanji: card.String(),
		}
		for _, details := range card.Details {
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

// json2KanjiCards converts a list of json.KanjiCard to kanjis.Card.
func json2KanjiCards(jsoncards []KanjiCard) []kanjis.Card {
	result := make([]kanjis.Card, 0, len(jsoncards))
	for _, jsoncard := range jsoncards {
		kanji, _ := utf8.DecodeRuneInString(jsoncard.Kanji)
		card := kanjis.Card{Kanji: kanji}
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

// wordCards2Json converts a list of words.Card to json.WordCard.
func wordCards2Json(cards []words.Card) []WordCard {
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

// json2WordCards converts a list of json.WordCard to words.Card.
func json2WordCards(jsoncards []WordCard) []words.Card {
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