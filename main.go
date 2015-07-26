package main

import (
	"flag"
	"fmt"
	"log"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version $(git describe --tags)"
var version string
var name string = "go-template"

// command line flags
var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("%s version %s\n", name, version)
	} else {
		// TODO: implement your stuff here...
		fmt.Println("go-template is doing nothing.")
	}
}
