package main

import (
	"flag"

	"github.com/cuckflong/gophish-analytics/pkg/logging"
	"github.com/cuckflong/gophish-analytics/pkg/runner"
)

func main() {
	var inputFile string
	var mode int

	flag.StringVar(&inputFile, "if", "", "input CSV file")
	flag.IntVar(&mode, "mode", 0, "actions to perform")
	flag.BoolVar(&logging.Debug, "debug", false, "use debug mode")
	flag.Parse()

	newRunner := runner.NewRunner(inputFile, mode)
	newRunner.RunMode()
}
