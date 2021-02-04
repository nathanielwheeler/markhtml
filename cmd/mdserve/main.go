package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"markhtml/pkg/handle"
	"markhtml/pkg/parse"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// TODO refactor
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

	var buf *bytes.Buffer
	if !parseYAML {
		buf, err = parse.MarkdownToHTML(f.Name())
		if err != nil {
			return err
		}

	} else {
		// FIXME returns empty map for y
		buf, y, err := parse.MarkdownToHTMLWithYAML(f.Name())
		if err != nil {
			return err
		}
		fmt.Println(buf)
		fmt.Println(y)
  }
  
  html := template.HTML(buf.String())

	h := handle.Handler{HTML: &html}
	port := 1729
	fmt.Printf("Listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), &h))

	return nil
}
