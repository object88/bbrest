package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/object88/bbrest/dtos"
)

// Handler abstracts out routes
type Handler struct {
}

func (h *Handler) writeSuccessResponse(p *dtos.Photo, httpStatus int, w http.ResponseWriter) {
	uj, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "%s", uj)
}
