package handle

import (
	"html/template"
	"net/http"
)

// Handler will hold methods and values needed to serve HTML.
type Handler struct {
	HTML *template.HTML
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("web/md.tpl")
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(w, "markdown", h.HTML)
	if err != nil {
		panic(err)
	}
}
