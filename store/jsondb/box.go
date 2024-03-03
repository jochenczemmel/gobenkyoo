package jsondb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// readBox reads a single learn box from the given file.
func readBox(filename string) (learn.Box, error) {
	var box learn.Box
	file, err := os.Open(filename)
	if err != nil {
		return box, fmt.Errorf("open box file: %w", err)
	}
	defer file.Close()

	var jsonBox boxJSON
	err = json.NewDecoder(file).Decode(&jsonBox)
	if err != nil {
		return box, fmt.Errorf("json box %q: decode: %w", filename, err)
	}

	boxID := learn.BoxID{
		Name: jsonBox.BoxID.Name,
		LessonID: books.LessonID{
			Name: jsonBox.BoxID.LessonID.Name,
			ID: books.ID{
				Title:       jsonBox.BoxID.LessonID.BookID.Title,
				SeriesTitle: jsonBox.BoxID.LessonID.BookID.SeriesTitle,
				Volume:      jsonBox.BoxID.LessonID.BookID.Volume,
			},
		},
	}
	if jsonBox.Type == learn.KanjiType {
		box = learn.NewKanjiBox(boxID)
	} else {
		box = learn.NewWordBox(boxID)
	}

	for mode, modeLevels := range jsonBox.Cards {
		for level, cards := range modeLevels {
			box.AddCards(mode, level, json2LearnCards(cards)...)
		}
	}

	return box, nil
}

// storeBox stores a single box in the given directory.
// The name of the file is determined by calling fileName()
// on the boxID.
func storeBox(dirname string, box learn.Box) error {
	err := os.MkdirAll(dirname, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("store book: create directory: %w", err)
	}

	jsonBox := boxJSON{
		Type: box.Type,
		BoxID: boxIDJSON{
			Name: box.BoxID.Name,
			LessonID: lessonIDJSON{
				Name: box.BoxID.LessonID.Name,
				BookID: bookIDJSON{
					Title:       box.BoxID.LessonID.ID.Title,
					SeriesTitle: box.BoxID.LessonID.ID.SeriesTitle,
					Volume:      box.BoxID.LessonID.ID.Volume,
				},
			},
		},
		Cards: map[string]map[int][]learnCardJSON{},
	}

	for _, mode := range box.Modes() {
		jsonBox.Cards[mode] = map[int][]learnCardJSON{}
		for _, level := range learn.Levels() {
			cards := learnCards2Json(box.Cards(mode, level))
			if len(cards) > 0 {
				jsonBox.Cards[mode][level] = cards
			}
		}
	}

	fileName := filepath.Join(dirname, jsonBox.BoxID.fileName())
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("store box: create file: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	err = enc.Encode(jsonBox)
	if err != nil {
		return fmt.Errorf("store box %v: encode json: %w", jsonBox.BoxID, err)
	}

	return nil
}

type boxJSON struct {
	BoxID boxIDJSON                          `json:"boxId"`
	Type  string                             `json:"type"`
	Cards map[string]map[int][]learnCardJSON `json:"cards"`
}

type boxIDJSON struct {
	Name     string       `json:"name"`
	LessonID lessonIDJSON `json:"lessonId"`
}

// fileName returns the name of the file for storage.
// It consists of the box name, the lesson name, the book title,
// the book series title and the book volume. The name is
// path escaped to ensure it is a valid file name.
func (b boxIDJSON) fileName() string {
	return url.PathEscape(
		b.Name+"\n"+
			b.LessonID.Name+"\n"+
			b.LessonID.BookID.Title+"\n"+
			b.LessonID.BookID.SeriesTitle+"\n"+
			strconv.Itoa(b.LessonID.BookID.Volume)) +
		jsonExtension
}

type lessonIDJSON struct {
	Name   string     `json:"name"`
	BookID bookIDJSON `json:"bookId"`
}

type cardIDJSON struct {
	ContentID string       `json:"contentId"`
	LessonID  lessonIDJSON `json:"lessonId"`
}

type learnCardJSON struct {
	ID          cardIDJSON `json:"id"`
	Question    string     `json:"question"`
	Hint        string     `json:"hint,omitempty"`
	Answer      string     `json:"answer"`
	MoreAnswers []string   `json:"moreAnswers,omitempty"`
	Explanation string     `json:"explanation,omitempty"`
}

// learnCardJSON converts a list of learn cards to json cards.
func learnCards2Json(cards []learn.Card) []learnCardJSON {
	jsonCards := make([]learnCardJSON, 0, len(cards))
	for _, c := range cards {
		jsonCards = append(jsonCards, learnCardJSON{
			ID: cardIDJSON{
				ContentID: c.ID.ContentID,
				LessonID: lessonIDJSON{
					Name: c.ID.LessonID.Name,
					BookID: bookIDJSON{
						Title:       c.ID.LessonID.ID.Title,
						SeriesTitle: c.ID.LessonID.ID.SeriesTitle,
						Volume:      c.ID.LessonID.ID.Volume,
					},
				},
			},
			Question:    c.Question,
			Hint:        c.Hint,
			Answer:      c.Answer,
			MoreAnswers: c.MoreAnswers,
			Explanation: c.Explanation,
		})
	}

	return jsonCards
}

// json2LearnCards converts a list of json cards to learn cards.
func json2LearnCards(cards []learnCardJSON) []learn.Card {
	jsonCards := make([]learn.Card, 0, len(cards))
	for _, c := range cards {
		card := learn.Card{
			ID: learn.CardID{
				ContentID: c.ID.ContentID,
				LessonID: books.LessonID{
					Name: c.ID.LessonID.Name,
					ID: books.ID{
						Title:       c.ID.LessonID.BookID.Title,
						SeriesTitle: c.ID.LessonID.BookID.SeriesTitle,
						Volume:      c.ID.LessonID.BookID.Volume,
					},
				},
			},
			Question:    c.Question,
			Hint:        c.Hint,
			Answer:      c.Answer,
			MoreAnswers: c.MoreAnswers,
			Explanation: c.Explanation,
		}
		jsonCards = append(jsonCards, card)
	}

	return jsonCards
}
