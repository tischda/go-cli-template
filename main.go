package main

import (
	"flag"
	"fmt"
	"log"
)

const PROG_NAME string = "PROJECT_NAME"

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
var version string

// command line flags
var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("%s version %s\n", PROG_NAME, version)
	} else {
		// TODO: implement your stuff here...
		fmt.Println(PROG_NAME, "is doing nothing.")
	}
}
