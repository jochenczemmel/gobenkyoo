package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func main() {
	getOptions()
	err := execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(ERROR_RC)
	}
}

// execute does the main work.
func execute() error {

	database := jsondb.New(filepath.Join(optConfigDir, jsondb.BaseDir))

	lib, room, err := load(database)
	if err != nil {
		return err
	}

	bookID := books.NewID(optBookTitle, optSeriesTitle, optVolume)

	boxID := learn.BoxID{
		Name: optLessonTitle,
		LessonID: books.LessonID{
			Name: optLessonTitle,
			ID:   bookID,
		},
	}

	if optFromFileName != "" {
		cards, err := getKanjiCards(lib)
		if err != nil {
			return err
		}
		room.SetKanjiBoxes(learn.NewKanjiBox(boxID, cards...))

		return database.StoreClassroom(room)
	}

	lesson, ok := lib.Book(bookID).Lesson(optLessonTitle)
	if !ok {
		return fmt.Errorf("lesson %q not found in book %q",
			optLessonTitle, bookID)
	}

	switch optType {
	case learn.KanjiType:
		room.SetKanjiBoxes(learn.NewKanjiBox(boxID, lesson.KanjiCards()...))
	case learn.WordType:
		room.SetWordBoxes(learn.NewWordBox(boxID, lesson.WordCards()...))
	default:
		room.SetKanjiBoxes(learn.NewKanjiBox(boxID, lesson.KanjiCards()...))
		room.SetWordBoxes(learn.NewWordBox(boxID, lesson.WordCards()...))
	}

	return database.StoreClassroom(room)
}

func getKanjiCards(lib books.Library) ([]kanjis.Card, error) {
	var result []kanjis.Card
	data, err := os.ReadFile(optFromFileName)
	if err != nil {
		return result, err
	}

	book := lib.Book(books.NewID(optFromBook, optFromSeries, optFromVolume))
	cardsByKanji := map[rune]kanjis.Card{}
	for _, lesson := range book.Lessons() {
		for _, card := range lesson.KanjiCards() {
			cardsByKanji[card.Kanji] = card
		}
	}

	for _, wantKanji := range string(data) {
		if found, ok := cardsByKanji[wantKanji]; ok {
			result = append(result, found)
		}
	}

	fmt.Printf("DEBUG: cards: %#v\n", result)

	return result, nil
}

// load loads the library and the classroom.
func load(db jsondb.DB) (books.Library, learn.Classroom, error) {
	lib, err1 := loadLib(db)
	room, err2 := loadClassroom(db)
	if err1 != nil {
		return lib, room, err1
	}
	if err2 != nil {
		return lib, room, err2
	}
	return lib, room, nil
}

// loadLib loads the library and checks the error status.
// It is not considered an error if the library does not exist.
func loadLib(db jsondb.DB) (books.Library, error) {

	lib, err := db.LoadLibrary(cfg.DefaultLibrary)
	if err == nil {
		return lib, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		fmt.Fprintln(os.Stderr, "no library found, create new")
		return lib, nil
	}

	return lib, err
}

// loadClassroom loads the classroom and checks the error status.
// It is not considered an error if the classroom does not exist.
func loadClassroom(db jsondb.DB) (learn.Classroom, error) {

	room, err := db.LoadClassroom(cfg.DefaultClassroom)
	if err == nil {
		return room, nil
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) && os.IsNotExist(pathErr) {
		fmt.Fprintln(os.Stderr, "no classroom found, create new")
		return learn.NewClassroom(cfg.DefaultClassroom), nil
	}

	return room, err
}

// getOptions gets the command line options and stores them
// in global variables. Some plausibility checks are made.
// If an error occurrs, a usage note is displayed and the
// program exits with value 1.
func getOptions() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`%s - create kanji and word learn boxes from lessons.
kanji and word boxes can be created directly from lessons.
kanji boxes can be created from a list of kanji in a file and
the kanji lesson reference

`, os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "options:")
		flag.PrintDefaults()
	}

	flag.StringVar(&optConfigDir, "configdir", "", "configuration directory")

	flag.StringVar(&optType, "type", "", "data type: kanji or word\n"+
		"if missing, do both")

	flag.StringVar(&optLessonTitle, "lesson", "", "name of the lesson")
	flag.StringVar(&optBookTitle, "book", "", "name of the book")
	flag.StringVar(&optSeriesTitle, "series", "", "name of the book series")
	flag.IntVar(&optVolume, "volume", 0, "volume of the book in the series")

	flag.StringVar(&optFromFileName, "fromfile", "", "file containing kanjis\n"+
		"type is set to 'kanji'")
	flag.StringVar(&optFromBook, "frombook", "", "book containing kanjis")
	flag.StringVar(&optFromSeries, "fromseries", "", "book series of 'frombook'")
	flag.IntVar(&optFromVolume, "fromvolume", 0, "book series volume of 'frombook'")

	flag.Parse()

	exitIfEmpty("lesson", optLessonTitle)
	exitIfEmpty("book", optBookTitle)

	if optType != learn.WordType &&
		optType != learn.KanjiType &&
		optType != "" {
		fmt.Fprintf(flag.CommandLine.Output(), "invalid type: %q", optType)
		flag.Usage()
		os.Exit(USAGE_RC)
	}

	if optFromFileName != "" {
		exitIfEmpty("frombook", optFromBook)
		optType = learn.KanjiType
	}

	var err error
	if optConfigDir == "" {
		optConfigDir, err = cfg.UserDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "can not determine config dir: %v", err)
			os.Exit(USAGE_RC)
		}
	}
}

// exitIfEmpty exits the program if the string value is empty.
func exitIfEmpty(label, value string) {
	if value == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "%s missing\n", label)
		flag.Usage()
		os.Exit(USAGE_RC)
	}
}

var (
	optConfigDir    string
	optType         string
	optLessonTitle  string
	optBookTitle    string
	optSeriesTitle  string
	optVolume       int
	optFromFileName string
	optFromBook     string
	optFromSeries   string
	optFromVolume   int
)

const (
	// ERROR_RC is the return code in case of an error.
	ERROR_RC = 2
	// USAGE_RC is the return code in case of an invalid call.
	USAGE_RC = 1
)
