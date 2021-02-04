package parse

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestMarkdownToHTML(t *testing.T) {
	is := is.New(t)

	b, err := MarkdownToHTML("_test.md")
	if err != nil {
    fmt.Printf("%s\n", err)
    is.Fail()
  }

  var buf bytes.Buffer
  html := `<h1>Header</h1>
<p>Paragraph body, with <em>italics</em> and with <strong>bold</strong>.</p>
<ol>
<li>Ordered</li>
<li>List</li>
</ol>
<ul>
<li>Unordered</li>
<li>List</li>
</ul>
<blockquote>
<p>Quote</p>
</blockquote>
`
  buf.WriteString(html)
  
  is.Equal(b, &buf)
}

// func TestMarkdownToHTMLWithYAML(t *testing.T) {
//   is := is.New(t)

//   b, y, err := MarkdownToHTMLWithYAML("_testyaml.md")
//   if err != nil {
//     fmt.Printf("%s\n", err)
//     is.Fail()
//   }

// }
