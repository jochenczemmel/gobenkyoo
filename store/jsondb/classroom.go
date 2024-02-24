package jsondb

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

// StoreClassroom stores the specified classroom in the base dir
// of the json library object.
func (l DB) StoreClassroom(classroom learn.Classroom) error {

	classroomDir := filepath.Join(l.baseDir, classroomPath)
	dirName := filepath.Join(classroomDir, kanjiPath, url.PathEscape(classroom.Name))
	for _, box := range classroom.KanjiBoxes() {
		err := storeBox(dirName, box)
		if err != nil {
			return fmt.Errorf("store classroom: %w", err)
		}
	}

	dirName = filepath.Join(classroomDir, wordPath, url.PathEscape(classroom.Name))
	for _, box := range classroom.WordBoxes() {
		err := storeBox(dirName, box)
		if err != nil {
			return fmt.Errorf("store classroom: %w", err)
		}
	}

	return nil
}

func (l DB) LoadClassroom(name string) (learn.Classroom, error) {
	room := learn.NewClassroom(name)
	baseDirName := filepath.Join(l.baseDir, classroomPath)
	boxes, err := readBoxes(name,
		filepath.Join(baseDirName, kanjiPath, url.PathEscape(name)))
	if err != nil {
		return room, fmt.Errorf("load classroom: %w", err)
	}
	room.SetKanjiBoxes(boxes...)

	boxes, err = readBoxes(name,
		filepath.Join(baseDirName, wordPath, url.PathEscape(name)))
	if err != nil {
		return room, fmt.Errorf("load classroom: %w", err)
	}
	room.SetWordBoxes(boxes...)

	return room, nil
}

func readBoxes(name, dirname string) ([]learn.Box, error) {
	result := []learn.Box{}

	dir, err := os.Open(dirname)
	if err != nil {
		return result, fmt.Errorf("open box directory: %w", err)
	}
	defer dir.Close()

	errorList := []error{}
	files, err := dir.ReadDir(readAllFiles)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), jsonExtension) {
			continue
		}
		box, err := readBox(filepath.Join(dirname, file.Name()))
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		result = append(result, box)
	}

	return result, nil
}
