package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jochenczemmel/gobenkyoo/app"
	"github.com/jochenczemmel/gobenkyoo/cfg"
)

// command line flags
var (
	optDbPath string // data base path specification
	optDbType string // type of data base
	optLearn  bool   // start learn cli
	optGui    bool   // start GUI
)

func main() {

	err := doBenkyoo()
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(2)
	}
}

// doBenkyoo executes the program.
func doBenkyoo() error {

	err := getConfig()
	if err != nil {
		return fmt.Errorf("read config: %v", err)
	}
	getOptions()

	// TODO: add options
	application := app.New()

	// TODO: add options
	err = application.Load()
	if err != nil {
		return fmt.Errorf("load data: %v", err)
	}

	// TODO: add options
	return application.Run()
}

// getConfig reads the configuration file.
// TODO: better move to app.App?
func getConfig() error {
	// TODO: read config file if it exists
	// _, err := os.Stat(cfg.DefaultCfgFile)
	return nil
}

// getOptions parses the command line options.
func getOptions() {

	// set usage note
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s - learn japanese\n", os.Args[0])
		flag.PrintDefaults()
	}

	// TODO: add/modify parameters, defaults, ...

	// define flags
	flag.StringVar(&optDbPath, "dbtype", os.Getenv("GOBENKYOO_DB_TYPE"),
		"config file (default: $GOBENKYOO_DB_TYPE)")
	flag.StringVar(&optDbPath, "db", os.Getenv("GOBENKYOO_DB_PATH"),
		"config file (default: $GOBENKYOO_DB_PATH)")

	flag.BoolVar(&optLearn, "learn", false, "learn in cli mode")
	flag.BoolVar(&optGui, "gui", false, "start app with GUI")

	// parse options
	flag.Parse()

	// check parameter
	if optDbPath == "" {
		optDbPath = cfg.DefaultDbPath
	}
	if optDbType == "" {
		optDbType = cfg.DefaultDbType
	}

	if !optLearn && !optGui {
		fmt.Println("specify either -learn or -gui")
		flag.Usage()
		os.Exit(1)
	}
}
