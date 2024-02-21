package jsondb

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

// StoreClassroom stores the specified classroom in the base dir
// of the json library object.
func (l DB) StoreClassroom(classroom learn.Classroom) error {

	dirName := filepath.Join(l.baseDir, libraryPath, kanjiPath,
		url.PathEscape(classroom.Name))
	for _, box := range classroom.KanjiBoxes() {
		err := storeBox(dirName, box)
		if err != nil {
			return fmt.Errorf("store classroom: %w", err)
		}
	}

	dirName = filepath.Join(l.baseDir, libraryPath, wordPath,
		url.PathEscape(classroom.Name))
	for _, box := range classroom.WordBoxes() {
		err := storeBox(dirName, box)
		if err != nil {
			return fmt.Errorf("store classroom: %w", err)
		}
	}

	return nil
}
