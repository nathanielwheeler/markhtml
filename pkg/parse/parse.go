package parse

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// MarkdownToHTML takes a file, parses it, and returns HTML bytes.  It can also take any number of  parse options to pass into goldmark's Convert function.
// TODO return string instead of buffer
func MarkdownToHTML(f string, opts ...parser.ParseOption) (*template.HTML, error) {
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

  html := template.HTML(buf.String())
	return &html, nil
}

// MarkdownToHTMLWithYAML is like MarkdownToHTML, except that it also parses for a YAML header in the markdown file.
func MarkdownToHTMLWithYAML(f string, opts ...parser.ParseOption) (*template.HTML, map[string]interface{}, error) {
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
	opts = append(opts, parser.WithContext(ctx))

	err = md.Convert([]byte(src), &buf, opts...)
	if err != nil {
		return nil, nil, err
	}

	// FIXME parses an empty map
	yaml := meta.Get(ctx)
	html := template.HTML(buf.String())

	return &html, yaml, nil
}
