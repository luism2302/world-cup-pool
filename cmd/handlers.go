package main

import (
	"context"
	"net/http"

	"github.com/luism2302/world-cup-pool/components"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {
	components.Base().Render(context.Background(), w)
}
