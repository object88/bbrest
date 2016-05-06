package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

// PhotoRouter sets up routes for the PhotoHandler
type PhotoRouter struct {
	pH *PhotoHandler
}

// AddRouter adds the routes to the mux Router
func (pR *PhotoRouter) AddRouter(r *mux.Router, pC *PhotoController) {
	fmt.Printf("Adding photo router...\n")

	pH := NewPhotoHandler(pC)
	r.HandleFunc("/photo", pH.Handle).Methods("GET")
	r.HandleFunc("/photo/{id}", pH.HandleSingleGet).Methods("GET")
	r.HandleFunc("/photo", pH.HandleCreate).Methods("POST")

	fmt.Printf("Added photo router...\n")
}
