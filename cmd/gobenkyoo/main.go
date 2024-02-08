package main

import (
	"flag"
	"fmt"
	"os"
)

// command line flags.
var (
	optDbPath string // data base path specification
	optDbType string // type of data base
	optUI     string // type of user interface
)

const ERROR_RC = 2

// main program.
func main() {

	getOptions()
	err := doBenkyoo()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(ERROR_RC)
	}
}

// doBenkyoo executes the program.
func doBenkyoo() error {
	// TODO:
	fmt.Printf("ui: %s, path: %s, type: %s\n", optUI, optDbPath, optDbType)
	return nil
}

// getOptions parses the command line options.
func getOptions() {

	// set usage note
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s - learn japanese\n", os.Args[0])
		flag.PrintDefaults()
	}

	// define flags
	flag.StringVar(&optDbType, "dbtype", "", "config file")
	// flag.StringVar(&optDbPath, "db", "", "database path")

	flag.StringVar(&optUI, "ui", "", "user interface mode")

	// parse options
	flag.Parse()
}
