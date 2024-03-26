package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/jochenczemmel/gobenkyoo/app"
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
		log.Printf("ERROR: %v", err)
		os.Exit(ERROR_RC)
	}
}

// execute does the main work.
func execute() error {

	importer := app.NewLibraryImporter(cfg.DefaultLibrary,
		jsondb.New(filepath.Join(optConfigDir, jsondb.BaseDir)),
	)
	if !optNoMinify {
		jsondb.Minify = true
	}

	lessonID := books.LessonID{
		Name: optLessonTitle,
		ID: books.ID{
			Title:       optBookTitle,
			SeriesTitle: optSeriesTitle,
			Volume:      optVolume,
		},
	}

	if optType == learn.KanjiType {
		return importer.Kanji(
			csvimport.NewKanji(optSeparatorRune, optFieldSeparatorRune,
				optHeaderLine, strings.Split(optFields, fieldSplitChar),
			),
			optFileName, lessonID,
		)
	}

	return importer.Word(
		csvimport.NewWord(optSeparatorRune, optHeaderLine,
			strings.Split(optFields, fieldSplitChar),
		),
		optFileName, lessonID,
	)
}

// getOptions gets the command line options and stores them
// in global variables. Some plausibility checks are made.
// If an error occurrs, a usage note is displayed and the
// program exits with value 1.
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
	flag.BoolVar(&optNoMinify, "nominify", false, "do not minify json files")

	flag.Parse()

	exitIfEmpty("file", optFileName)
	exitIfEmpty("type", optType)
	exitIfEmpty("lesson", optLessonTitle)
	exitIfEmpty("book", optBookTitle)
	exitIfEmpty("fields", optFields)

	if optType != learn.WordType && optType != learn.KanjiType {
		fmt.Fprintf(flag.CommandLine.Output(), "invalid type: %q", optType)
		flag.Usage()
		os.Exit(USAGE_RC)
	}

	var err error
	if optConfigDir == "" {
		optConfigDir, err = cfg.UserDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "can not determine config dir: %v", err)
			os.Exit(USAGE_RC)
		}
	}

	optSeparatorRune, _ = utf8.DecodeRuneInString(optSeparator)
	optFieldSeparatorRune, _ = utf8.DecodeRuneInString(optFieldSeparator)
}

// exitIfEmpty exits the program if the string value is empty.
func exitIfEmpty(label, value string) {
	if value == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "%s missing\n", label)
		flag.Usage()
		os.Exit(USAGE_RC)
	}
}

// global variables that represent command line options.
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
	optNoMinify           bool
)

const (
	// ERROR_RC is the return code in case of an error.
	ERROR_RC = 2
	// USAGE_RC is the return code in case of an invalid call.
	USAGE_RC = 1
)

const fieldSplitChar = ","
