package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// BoxCreator provides creating or updating learn boxes.
type BoxCreator struct {
	loadStorer ClassroomLoadStorer
	lib        books.Library
	room       learn.Classroom
}

// NewBoxCreator returns a creator that uses the given loadstorer.
func NewBoxCreator(loadstorer ClassroomLoadStorer) BoxCreator {
	return BoxCreator{
		loadStorer: loadstorer,
	}
}

// Load loads the library and classroom with the given names.
// It returns true if the classroom is found, false if it is not found.
// In case of another error (including not finding the library),
// the error is returned.
func (c *BoxCreator) Load(libname, roomname string) (found bool, err error) {

	if c.loadStorer == nil {
		return false, ConfigurationError("no ClassroomLoadStorer defined")
	}

	c.lib, err = c.loadStorer.LoadLibrary(libname)
	if err != nil {
		return false, err
	}

	c.room, err = c.loadStorer.LoadClassroom(roomname)
	if err == nil {
		return true, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		c.room = learn.NewClassroom(roomname)
		return false, nil
	}

	return false, err
}

// Store stores the classroom boxes.
func (c *BoxCreator) Store() error {
	if c.loadStorer == nil {
		return ConfigurationError("no ClassroomLoadStorer defined")
	}
	return c.loadStorer.StoreClassroom(c.room)
}

// KanjiBox creates a new kanji box from the lesson id provided
// in the box id.
func (c *BoxCreator) KanjiBox(id learn.BoxID) error {

	lesson, ok := c.lib.Book(id.LessonID.ID).Lesson(id.LessonID.Name)
	if !ok {
		return ConfigurationError(
			fmt.Sprintf("lesson %q not found in book %q",
				id.LessonID.Name, id.LessonID.ID),
		)
	}

	c.room.SetKanjiBoxes(learn.NewKanjiBox(id, lesson.KanjiCards()...))

	return nil
}

// WordBox creates a new word box from the lesson id provided
// in the box id.
func (c *BoxCreator) WordBox(id learn.BoxID) error {

	lesson, ok := c.lib.Book(id.LessonID.ID).Lesson(id.LessonID.Name)
	if !ok {
		return ConfigurationError(
			fmt.Sprintf("lesson %q not found in book %q",
				id.LessonID.Name, id.LessonID.ID),
		)
	}

	c.room.SetWordBoxes(learn.NewWordBox(id, lesson.WordCards()...))

	return nil
}

// KanjiBoxFromList creates a new kanji box from the lesson id provided
// in the from box id that matches the provided list.
func (c *BoxCreator) KanjiBoxFromList(kanjilist string, from books.ID, id learn.BoxID) error {

	cardsByKanji := map[rune]kanjis.Card{}

	for _, lesson := range c.lib.Book(from).Lessons() {
		for _, card := range lesson.KanjiCards() {
			cardsByKanji[card.Kanji] = card
		}
	}

	cards := make([]kanjis.Card, 0, len(kanjilist))
	for _, wantKanji := range kanjilist {
		if found, ok := cardsByKanji[wantKanji]; ok {
			cards = append(cards, found)
		}
	}

	c.room.SetKanjiBoxes(learn.NewKanjiBox(id, cards...))

	return nil
}
