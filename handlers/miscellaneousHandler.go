package handlers

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// AddMiscellaneousHandler maps unknown routes
func AddMiscellaneousHandler(mux *web.Mux) {
	mux.NotFound(NotFound)
}

// NotFound is a 404 handler.
func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Off and on!", 404)
}
