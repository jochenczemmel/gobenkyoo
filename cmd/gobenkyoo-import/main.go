package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
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

func execute() error {

	db := jsondb.New(filepath.Join(optConfigDir, jsondb.BaseDir))
	lib, err := loadLib(db)
	if err != nil {
		return err
	}

	book := lib.Book(books.ID{
		Title:       optBookTitle,
		SeriesTitle: optSeriesTitle,
		Volume:      optVolume,
	})

	lesson, err := fillLesson(book)
	if err != nil {
		return err
	}

	book.SetLessons(lesson)
	lib.SetBooks(book)
	err = db.StoreLibrary(lib)
	if err != nil {
		return err
	}

	return nil
}

func fillLesson(book books.Book) (books.Lesson, error) {

	lesson, ok := book.Lesson(optLessonTitle)
	if !ok {
		lesson.Name = optLessonTitle
	}

	if optType == learn.KanjiType {
		format, err := csvimport.NewKanjiFormat(
			strings.Split(optFields, fieldSplitChar)...)
		if err != nil {
			return lesson, err
		}
		importer := csvimport.Kanji{
			Format:         format,
			Separator:      optSeparatorRune,
			FieldSeparator: optFieldSeparatorRune,
			HeaderLine:     optHeaderLine,
		}
		cards, err := importer.Import(optFileName)
		if err != nil {
			return lesson, err
		}
		lesson.AddKanjis(cards...)
	} else {
		format, err := csvimport.NewWordFormat(
			strings.Split(optFields, fieldSplitChar)...)
		if err != nil {
			return lesson, err
		}
		importer := csvimport.Word{
			Format:     format,
			Separator:  optSeparatorRune,
			HeaderLine: optHeaderLine,
		}
		cards, err := importer.Import(optFileName)
		if err != nil {
			return lesson, err
		}
		lesson.AddWords(cards...)
	}

	return lesson, nil
}

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

func getOptions() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`%s - import kanji and word (vocabulary) data

the order of the fields in the csv file must be specified
by a list of field names. Valid field names are (case insensitive):

word: %v
kanji: %v

Skipping fields in csv can be achieved by missing field names, e.g.
Meaning,Nihongo,,,Dictform,Teform,Naiform

`, os.Args[0],
			strings.Join(csvimport.WordFields(), ","),
			strings.Join(csvimport.KanjiFields(), ","),
		)
		fmt.Fprintln(flag.CommandLine.Output(), "options:")
		flag.PrintDefaults()
	}

	flag.StringVar(&optConfigDir, "configdir", "", "configuration directory")

	flag.StringVar(&optFileName, "file", "", "name of the csv file")
	flag.StringVar(&optType, "type", "", "import data type: kanji or word")
	flag.StringVar(&optFields, "fields", "", "comma separated list of fields")
	flag.StringVar(&optSeparator, "sep", ";", "csv column separator")
	flag.StringVar(&optFieldSeparator, "fieldsep", "", "content separator (kanji only)")
	flag.BoolVar(&optHeaderLine, "header", false, "skip header line")

	flag.StringVar(&optLessonTitle, "lesson", "", "name of the lesson")
	flag.StringVar(&optBookTitle, "book", "", "name of the book")
	flag.StringVar(&optSeriesTitle, "series", "", "name of the book series")
	flag.IntVar(&optVolume, "volume", 0, "volume of the book in the series")

	flag.Parse()

	exitIfEmpty("file", optFileName)
	exitIfEmpty("type", optType)
	exitIfEmpty("lesson", optLessonTitle)
	exitIfEmpty("book", optBookTitle)
	exitIfEmpty("fields", optFields)

	if optType != learn.WordType && optType != learn.KanjiType {
		fmt.Fprintf(flag.CommandLine.Output(), "invalid type: %q", optType)
		flag.Usage()
		os.Exit(1)
	}

	var err error
	if optConfigDir == "" {
		optConfigDir, err = cfg.UserDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "can not determine config dir: %v", err)
			os.Exit(1)
		}
	}

	optSeparatorRune, _ = utf8.DecodeRuneInString(optSeparator)
	optFieldSeparatorRune, _ = utf8.DecodeRuneInString(optFieldSeparator)
}

func exitIfEmpty(label, value string) {
	if value == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "%s missing\n", label)
		flag.Usage()
		os.Exit(1)
	}
}

var (
	optConfigDir          string
	optType               string
	optFileName           string
	optFields             string
	optSeparator          string
	optFieldSeparator     string
	optSeparatorRune      rune
	optFieldSeparatorRune rune
	optHeaderLine         bool
	optLessonTitle        string
	optBookTitle          string
	optSeriesTitle        string
	optVolume             int
)

const ERROR_RC = 2

const fieldSplitChar = ","
