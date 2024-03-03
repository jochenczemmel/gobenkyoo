package jsondb

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
)

// StoreClassroom stores all boxes that the specified classroom contains.
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

// LoadClassroom loads all boxes that the specified classroom contains.
func (l DB) LoadClassroom(name string) (learn.Classroom, error) {
	var room learn.Classroom

	baseDirName := filepath.Join(l.baseDir, classroomPath)
	kanjiBoxes, err := readBoxes(name,
		filepath.Join(baseDirName, kanjiPath, url.PathEscape(name)))
	if err != nil {
		return room, fmt.Errorf("load classroom: %w", err)
	}

	wordBoxes, err := readBoxes(name,
		filepath.Join(baseDirName, wordPath, url.PathEscape(name)))
	if err != nil {
		return room, fmt.Errorf("load classroom: %w", err)
	}

	room = learn.NewClassroom(name)
	room.SetKanjiBoxes(kanjiBoxes...)
	room.SetWordBoxes(wordBoxes...)

	return room, nil
}

// readBoxes reads all learn boxes that are found in the given directory.
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

	if len(errorList) > 0 {
		return result, fmt.Errorf("read box: %w", errors.Join(errorList...))
	}

	return result, nil
}
