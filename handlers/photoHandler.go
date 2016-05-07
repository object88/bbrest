package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/object88/bbrest/controllers"
	"github.com/object88/bbrest/models"
)

// PhotoHandler class!
type PhotoHandler struct {
	Handler
	photoController *controllers.PhotoController
}

// AddPhotoHandler creates a new instance
func AddPhotoHandler(pC *controllers.PhotoController, r *mux.Router) {
	fmt.Printf("Adding photo router...\n")

	pH := &PhotoHandler{Handler{}, pC}

	r.HandleFunc("/photo", pH.Handle).Methods("GET")
	r.HandleFunc("/photo/{id}", pH.HandleSingleGet).Methods("GET")
	r.HandleFunc("/photo", pH.HandleCreate).Methods("POST")

	fmt.Printf("Added photo router...\n")
}

// Handle processes a request for a set of photos.
func (pH *PhotoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered Handle\n")
}

// HandleCreate handles photo creation
func (pH *PhotoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered HandleCreate\n")

	p := &models.Photo{}
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
