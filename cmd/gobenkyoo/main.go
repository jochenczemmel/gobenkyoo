package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app"
)

// command line flags
var (
	optDbPath string // data base path specification
	optDbType string // type of data base
	optLearn  bool   // start learn cli
	optGui    bool   // start GUI
)

// main program
func main() {

	err := doBenkyoo()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}
}

// doBenkyoo executes the program.
func doBenkyoo() error {

	getOptions()

	// TODO: add options
	application := app.New()
	// app.WithLoader(store.NewLoader(optDbType, optDbPath)),

	// TODO: add options
	err := application.Load()
	if err != nil {
		return fmt.Errorf("load data: %v", err)
	}

	// TODO: add options
	return application.Run()
}

// getOptions parses the command line options.
func getOptions() {

	// set usage note
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s - learn japanese\n", os.Args[0])
		flag.PrintDefaults()
	}

	// define flags
	flag.StringVar(&optDbPath, "dbtype", "", "config file")
	flag.StringVar(&optDbPath, "db", "", "database path")

	flag.BoolVar(&optLearn, "learn", false, "learn in cli mode")
	flag.BoolVar(&optGui, "gui", false, "start app with GUI")

	// parse options
	flag.Parse()
}
