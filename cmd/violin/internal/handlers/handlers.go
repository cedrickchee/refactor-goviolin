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
func (b *Base) Home(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	pv := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(w, "home.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR : %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}
