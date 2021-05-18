package main

import (
	"fmt"
)

const (
	binary = `##PROJECT_NAME##`

	// TODO: edit usage
	usage = "\n" + binary + " does nothing.\n\n" +
		"Usage: %s [OPTIONS] <args>...\n\nOPTIONS:\n"
)

// set via -ldflags
var version string

func main() {
	parseFlags()

	// TODO: implement your stuff here
	fmt.Println("##PROJECT_NAME## is doing nothing.")
}
