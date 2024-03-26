package app

import (
	"errors"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

// BoxCreator provides creating or updating learn boxes.
type BoxCreator struct {
	loadStorer ClassroomLoadStorer
	Library    books.Library
	Classroom  learn.Classroom
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

	c.Library, err = c.loadStorer.LoadLibrary(libname)
	if err != nil {
		return false, err
	}

	c.Classroom, err = c.loadStorer.LoadClassroom(roomname)
	if err == nil {
		return true, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		c.Classroom = learn.NewClassroom(roomname)
		return false, nil
	}

	return false, err
}

// Store stores the classroom boxes.
func (c *BoxCreator) Store() error {
	if c.loadStorer == nil {
		return ConfigurationError("no ClassroomLoadStorer defined")
	}
	return c.loadStorer.StoreClassroom(c.Classroom)
}

// KanjiBox creates a new kanji box from the lesson id provided
// in the box id.
func (c *BoxCreator) KanjiBox(id learn.BoxID) {
	lesson := c.getLesson(id)
	c.Classroom.SetKanjiBoxes(learn.NewKanjiBox(id, lesson.KanjiCards()...))
}

func (c *BoxCreator) getLesson(id learn.BoxID) books.Lesson {
	return c.Library.Book(id.LessonID.ID).Lesson(id.LessonID.Name)
}

// WordBox creates a new word box from the lesson id provided
// in the box id.
func (c *BoxCreator) WordBox(id learn.BoxID) {
	lesson := c.getLesson(id)
	c.Classroom.SetWordBoxes(learn.NewWordBox(id, lesson.WordCards()...))
}

// KanjiBoxFromList creates a new kanji box from the lesson id provided
// in the from box id that matches the provided list.
func (c *BoxCreator) KanjiBoxFromList(kanjilist string, from books.ID, id learn.BoxID) {

	uniqueKanjis := make(map[rune]bool, len(kanjilist))
	for _, k := range kanjilist {
		uniqueKanjis[k] = true
	}

	box := learn.NewKanjiBox(id)
	book := c.Library.Book(from)

	// LOOP:
	for _, lesson := range book.Lessons() {
		var cards []kanjis.Card
		for _, card := range lesson.KanjiCards() {
			if uniqueKanjis[card.Kanji] {
				cards = append(cards, card)
			}
		}
		box.AddKanjiCards(books.LessonID{
			Name: lesson.Name,
			ID:   book.ID,
		}, cards...)
	}

	c.Classroom.SetKanjiBoxes(box)
}
