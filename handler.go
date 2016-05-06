package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler abstracts out routes
type Handler struct {
}

func (h *Handler) writeSuccessResponse(p *Photo, httpStatus int, w http.ResponseWriter) {
	uj, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "%s", uj)
}
