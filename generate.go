package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	modfile "golang.org/x/mod/modfile"
)

//go:generate echo Setting up your files...
//go:generate go run generate.go

func main() {

	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	moduleName := modfile.ModulePath(goModBytes)
	projectName := filepath.Base(moduleName)

	f, err := os.Create("README.generated")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readmeTemplate.Execute(f, struct {
		Project    string
		Repository string
	}{
		Project:    projectName,
		Repository: moduleName,
	})
}

var readmeTemplate = template.Must(template.New("").Parse(`# {{ .Project }} [![Test](https://{{ .Repository }}/actions/workflows/test.yml/badge.svg)](https://{{ .Repository }}/actions/workflows/test.yml)

Utility written in [Go](https://www.golang.org).

### Install

~~~
go install {{ .Repository }}@latest
~~~

### Usage

~~~

~~~

Example:

~~~

~~~
`))

//go:generate rm generate.go
