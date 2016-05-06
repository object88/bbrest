package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PhotoHandler class!
type PhotoHandler struct {
	Handler
	photoController *PhotoController
}

// NewPhotoHandler creates a new instance
func NewPhotoHandler(pC *PhotoController) *PhotoHandler {
	return &PhotoHandler{Handler{}, pC}
}

// Handle processes a request for a set of photos.
func (pH *PhotoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered Handle")
}

// HandleCreate handles photo creation
func (pH *PhotoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered HandleCreate\n")

	p := &Photo{}
	json.NewDecoder(r.Body).Decode(p)

	pH.photoController.Create(p)

	pH.writeSuccessResponse(p, 201, w)
}

// HandleSingleGet processes a request for a single photos.
func (pH *PhotoHandler) HandleSingleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Printf("Entered HandleSingleGet with id '%s'.\n", id)

	p := pH.photoController.Get(id)

	pH.writeSuccessResponse(p, 200, w)
}
