//go:build ignore
// +build ignore

//go:generate echo Setting up your files...
//go:generate cmd /c go mod init github.com/%OWNER%/%PROJECT%
//go:generate cmd /c go run template.go --owner=%OWNER% --project=%PROJECT% --copyright=%COPYRIGHT%

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

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
	flag.StringVar(&cfg.copyright, "copyright", "2025", "copyright notice")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Parse()

	if cfg.owner == "" || cfg.project == "" {
		flag.Usage()
		log.Fatal("Missing required flags: --owner and --project")
	}
	if err := writeReadme(cfg); err != nil {
		log.Fatal(err)
	}
	if err := writeVersioninfo(cfg); err != nil {
		log.Fatal(err)
	}
}

func writeReadme(cfg *Config) error {
	t, err := template.ParseFiles("README.md.tpl")
	if err != nil {
		return err
	}

	f, err := os.Create("README.md")
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("Generating README.md")
	err = t.Execute(f, struct {
		Project    string
		Owner      string
		Repository string
	}{
		Project:    cfg.project,
		Owner:      cfg.owner,
		Repository: "github.com/" + cfg.owner + "/" + cfg.project,
	})
	if err != nil {
		return err
	}
	return nil
}

func writeVersioninfo(cfg *Config) error {
	t, err := template.ParseFiles("versioninfo.json.tpl")
	if err != nil {
		return err
	}

	f, err := os.Create("versioninfo.json")
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("Generating versioninfo.json")
	err = t.Execute(f, struct {
		Project   string
		Copyright string
	}{
		Project:   cfg.project,
		Copyright: cfg.copyright,
	})
	if err != nil {
		return err
	}
	return nil
}
