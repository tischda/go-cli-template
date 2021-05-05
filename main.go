package main

import (
	"flag"
	"fmt"
	"log"
)

// go build -ldflags=all="-X main.version=${BUILD_TAG} -s -w"
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
		fmt.Printf("PROJECT_NAME version %s\n", version)
	} else {
		// TODO: implement your stuff here...
		fmt.Println(PROJECT_NAME, "is doing nothing.")
	}
}
