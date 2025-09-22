//go:build ignore
// +build ignore

//go:generate echo Setting up your files...

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

const HOST = "github.com"

// flags
type Config struct {
	owner     string
	project   string
	copyright string
}

func initFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.owner, "owner", "", "project owner (github user or org)")
	flag.StringVar(&cfg.project, "project", "", "project name")
	flag.StringVar(&cfg.copyright, "copyright", "", "copyright notice")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Parse()

	if cfg.owner == "" || cfg.project == "" || cfg.copyright == "" {
		flag.Usage()
		log.Fatal("Missing required flags: --owner, --project or --copyright")
	}
	if err := writeReadme(cfg); err != nil {
		log.Fatal(err)
	}
	if err := writeVersioninfo(cfg); err != nil {
		log.Fatal(err)
	}
}

func writeReadme(cfg *Config) error {
	data := struct {
		Project    string
		Owner      string
		Repository string
	}{
		Project:    cfg.project,
		Owner:      cfg.owner,
		Repository: HOST + "/" + cfg.owner + "/" + cfg.project,
	}
	return processTemplate("README.md", data)
}

func writeVersioninfo(cfg *Config) error {
	data := struct {
		Project   string
		Copyright string
	}{
		Project:   cfg.project,
		Copyright: cfg.copyright,
	}
	return processTemplate("versioninfo.json", data)
}

// processTemplate centralizes template parsing, file creation, and execution.
func processTemplate(outPath string, data any) error {
	t, err := template.ParseFiles(outPath + ".tpl")
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Printf("Generating %s\n", outPath)
	return t.Execute(f, data)
}
