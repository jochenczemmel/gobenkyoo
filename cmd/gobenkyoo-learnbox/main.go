package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/app/learn"
	"github.com/jochenczemmel/gobenkyoo/cfg"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/store/jsondb"
)

func main() {
	getOptions()
	err := execute()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		os.Exit(ERROR_RC)
	}
}

// execute does the main work.
func execute() error {

	creator := app.NewBoxCreator(
		jsondb.New(filepath.Join(optConfigDir, jsondb.BaseDir)),
	)
	if !optNoMinify {
		jsondb.Minify = true
	}

	ok, err := creator.Load(cfg.DefaultLibrary, cfg.DefaultClassroom)
	if err != nil {
		return err
	}
	if !ok {
		log.Printf("classroom %q not found, create new", cfg.DefaultClassroom)
	}

	boxID := learn.BoxID{
		Name: optLessonTitle,
		LessonID: books.LessonID{
			Name: optLessonTitle,
			ID:   books.NewID(optBookTitle, optSeriesTitle, optVolume),
		},
	}

	if optFromFileName != "" {
		kanjiList, err := os.ReadFile(optFromFileName)
		if err != nil {
			return err
		}
		creator.KanjiBoxFromList(string(kanjiList),
			books.NewID(optFromBook, optFromSeries, optFromVolume), boxID)

		return creator.Store()
	}

	if optType == "" || optType == learn.KanjiType {
		creator.KanjiBox(boxID)
	}
	if optType == "" || optType == learn.WordType {
		creator.WordBox(boxID)
	}

	return creator.Store()
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
	flag.BoolVar(&optNoMinify, "nominify", false, "do not minify json files")

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

// global variables that represent command line options.
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
	optNoMinify     bool
)

const (
	// ERROR_RC is the return code in case of an error.
	ERROR_RC = 2
	// USAGE_RC is the return code in case of an invalid call.
	USAGE_RC = 1
)
