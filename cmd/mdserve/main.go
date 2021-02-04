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
  // TODO add flag option for YAML

	// get file from args
	if len(os.Args) < 2 {
		return errors.New("not enough arguments")
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

  b, err := parse.MarkdownToHTML(f.Name())
  
  fmt.Println(b)

	return nil
}
