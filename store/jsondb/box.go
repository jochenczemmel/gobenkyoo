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

func readBox(filename string) (learn.Box, error) {
	var box learn.Box
	file, err := os.Open(filename)
	if err != nil {
		return box, fmt.Errorf("open box file: %w", err)
	}
	defer file.Close()

	var jsonBox Box
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

func storeBox(dirname string, box learn.Box) error {
	err := os.MkdirAll(dirname, defaultFilePermissions)
	if err != nil {
		return fmt.Errorf("store book: create directory: %w", err)
	}

	jsonBox := Box{
		Type: box.Type,
		BoxID: BoxID{
			Name: box.BoxID.Name,
			LessonID: LessonID{
				Name: box.BoxID.LessonID.Name,
				BookID: BookID{
					Title:       box.BoxID.LessonID.ID.Title,
					SeriesTitle: box.BoxID.LessonID.ID.SeriesTitle,
					Volume:      box.BoxID.LessonID.ID.Volume,
				},
			},
		},
		Cards: map[string]map[int][]LearnCard{},
	}

	for _, mode := range box.Modes() {
		jsonBox.Cards[mode] = map[int][]LearnCard{}
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

type Box struct {
	BoxID BoxID                          `json:"boxId"`
	Type  string                         `json:"type"`
	Cards map[string]map[int][]LearnCard `json:"cards"`
}

type BoxID struct {
	Name     string   `json:"name"`
	LessonID LessonID `json:"lessonId"`
}

func (b BoxID) fileName() string {
	return url.PathEscape(
		b.Name+"\n"+
			b.LessonID.Name+"\n"+
			b.LessonID.BookID.Title+"\n"+
			b.LessonID.BookID.SeriesTitle+"\n"+
			strconv.Itoa(b.LessonID.BookID.Volume)) +
		jsonExtension
}

type LessonID struct {
	Name   string `json:"name"`
	BookID BookID `json:"bookId"`
}

type CardID struct {
	ContentID string   `json:"contentId"`
	LessonID  LessonID `json:"lessonId"`
}

type LearnCard struct {
	ID          CardID   `json:"id"`
	Question    string   `json:"question"`
	Hint        string   `json:"hint,omitempty"`
	Answer      string   `json:"answer"`
	MoreAnswers []string `json:"moreAnswers,omitempty"`
	Explanation string   `json:"explanation,omitempty"`
}

func learnCards2Json(cards []learn.Card) []LearnCard {
	jsonCards := make([]LearnCard, 0, len(cards))
	for _, c := range cards {
		jsonCards = append(jsonCards, LearnCard{
			ID: CardID{
				ContentID: c.ID.ContentID,
				LessonID: LessonID{
					Name: c.ID.LessonID.Name,
					BookID: BookID{
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

func json2LearnCards(cards []LearnCard) []learn.Card {
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
