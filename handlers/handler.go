package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/object88/bbrest/models"
)

// Handler abstracts out routes
type Handler struct {
}

func (h *Handler) writeSuccessResponse(p *models.Photo, httpStatus int, w http.ResponseWriter) {
	uj, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "%s", uj)
}
