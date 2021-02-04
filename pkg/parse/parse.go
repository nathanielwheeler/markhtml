package parse

import (
	"bytes"
	"io/ioutil"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// MarkdownToHTML takes a file, parses it, and returns HTML bytes.  It can also take any number of  parse options to pass into goldmark's Convert function.
func MarkdownToHTML(f string, opts ...parser.ParseOption) (*bytes.Buffer, error) {
	var (
		buf bytes.Buffer
		err error
	)

	src, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	err = goldmark.Convert(src, &buf, opts...)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

// MarkdownToHTMLWithYAML is like MarkdownToHTML, except that it also parses for a YAML header in the markdown file.
func MarkdownToHTMLWithYAML(f string, opts ...parser.ParseOption) (*bytes.Buffer, *map[string]interface{}, error) {
	var (
		buf bytes.Buffer
		md  goldmark.Markdown
		err error
	)

	src, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, nil, err
	}

	md = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
    ),
  )

  // Prepend parser context to parser options
  ctx := parser.NewContext()
	opts = append([]parser.ParseOption{parser.WithContext(ctx)}, opts...)

	err = md.Convert(src, &buf, opts...)
	if err != nil {
		return nil, nil, err
	}

	yaml := meta.Get(ctx)

	return &buf, &yaml, nil
}
