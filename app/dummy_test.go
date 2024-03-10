package app_test

import (
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/content/books"
)

// dummy is a test dummy ('test double').
// It implements some of the interfaces defined in the app package.
type dummy struct {
	loadError      string
	pathError      string
	storeError     string
	loadRoomError  string
	roomPathError  string
	storeRoomError string
}

func (d dummy) LoadClassroom(string) (learn.Classroom, error) {
	result := learn.NewClassroom("")
	if d.loadRoomError != "" {
		return result, fmt.Errorf("%s", d.loadRoomError)
	}
	if d.roomPathError != "" {
		return result, &os.PathError{
			Op:   "open",
			Path: ".",
			Err:  os.ErrNotExist,
		}
	}

	return result, nil
}

func (d dummy) StoreClassroom(learn.Classroom) error {
	if d.storeRoomError != "" {
		return fmt.Errorf("%s", d.storeRoomError)
	}

	return nil
}

func (d dummy) LoadLibrary(string) (books.Library, error) {
	var result books.Library
	if d.loadError != "" {
		return result, fmt.Errorf("%s", d.loadError)
	}
	if d.pathError != "" {
		return result, &os.PathError{
			Op:   "open",
			Path: ".",
			Err:  os.ErrNotExist,
		}
	}

	return result, nil
}

func (d dummy) StoreLibrary(books.Library) error {
	if d.storeError != "" {
		return fmt.Errorf("%s", d.storeError)
	}
	return nil
}
