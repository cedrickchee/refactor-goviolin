package handlers

import (
	"log"
	"net/http"

	"github.com/rosalita/goviolin/internal/render"
)

// Base represents the base handlers.
type Base struct {
	log *log.Logger
}

// Home handlerrenders the home.html page.
func (b *Base) Home(w http.ResponseWriter, req *http.Request) {
	pv := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(w, "home.html", pv); err != nil {
		b.log.Println(err)
		return
	}
}
