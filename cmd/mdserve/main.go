package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"markhtml/pkg/parse"
	// "markhtml/pkg/serve"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// flags
	parseYAML := false

	// check if enough arguments
	if len(os.Args) < 2 {
		return errors.New("not enough arguments")
	}

	// check for flags
	if len(os.Args) >= 3 {
		for _, v := range os.Args[1:] {
			switch v {
			case "-y":
			case "--yaml":
				parseYAML = true
			}
		}
	}

	// try to open file and check if filetype is markdown
	f, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	ext := filepath.Ext(os.Args[1])
	if ext != ".md" {
		return errors.New("invalid file type")
	}

	if !parseYAML {
		b, err := parse.MarkdownToHTML(f.Name())
		if err != nil {
			return err
		}
		fmt.Println(b)
	} else {
    // FIXME returns empty map for y
		b, y, err := parse.MarkdownToHTMLWithYAML(f.Name())
		if err != nil {
			return err
    }
		fmt.Println(b)
    fmt.Println(y)
	}

	return nil
}
