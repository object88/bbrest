package handlers

import (
	"fmt"
	"net/http"
)

// Handler abstracts out routes
type Handler struct {
}

func (h *Handler) writeSuccessResponse(uj []byte, httpStatus int, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, "%s", uj)
}
