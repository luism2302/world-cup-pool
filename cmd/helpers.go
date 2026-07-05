package main

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func renderTemplate(w http.ResponseWriter, t templ.Component) error {
	return t.Render(context.Background(), w)
}
