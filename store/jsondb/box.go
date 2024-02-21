package jsondb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

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
	}

	fileName := filepath.Join(dirname, jsonBox.BoxID.fileName())
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("store box: create file: %w", err)
	}
	defer file.Close()

	// TODO: extract cards, levels, learn modes, store them...

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	err = enc.Encode(jsonBox)
	if err != nil {
		return fmt.Errorf("store box %v: encode json: %w", jsonBox.BoxID, err)
	}

	return nil
}

type Box struct {
	BoxID BoxID       `json:"boxId"`
	Type  string      `json:"type"`
	Cards []LearnCard `json:"cards"`
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
