package parse

import (
	"bytes"
	"io/ioutil"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	meta "github.com/yuin/goldmark-meta"
  mathjax "github.com/litao91/goldmark-mathjax"
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

// MarkdownToHTMLWithExt is like MarkdownToHTML, except that it also parses for a YAML header in the markdown file, as well as MathJax.
func MarkdownToHTMLWithExt(f string, opts ...parser.ParseOption) (*string, *map[string]interface{}, error) {
	var (
		buf bytes.Buffer
		md  goldmark.Markdown
		err error
	)

	src, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, nil, err
	}

  // Prepend parser context to parser options
  ctx := parser.NewContext()
  opts = append(opts, parser.WithContext(ctx))

	md = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
      mathjax.MathJax,
		),
    goldmark.WithParserOptions(
      parser.WithAutoHeadingID(),
    ),
    goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)


	err = md.Convert([]byte(src), &buf, opts...)
	if err != nil {
		return nil, nil, err
	}

	// FIXME parses an empty map
	yaml := meta.Get(ctx)
	html := buf.String()

	return &html, &yaml, nil
}
