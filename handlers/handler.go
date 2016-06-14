package handlers

import (
	"fmt"
	"net/http"
)

// Handler abstracts out routes
type Handler struct {
}

func (h *Handler) writeSuccessResponse(uj []byte, httpStatus int, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)

	if uj != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", uj)
	}
}
