package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/object88/bbrest/controllers"
	"github.com/object88/bbrest/dtos"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// PhotoHandler class!
type PhotoHandler struct {
	Handler
	photoController *controllers.PhotoController
}

// AddPhotoHandler creates a new instance
func AddPhotoHandler(pC *controllers.PhotoController, parentMux *web.Mux) {
	fmt.Printf("Adding photo router...\n")

	pH := &PhotoHandler{Handler{}, pC}

	mux := web.New()
	parentMux.Handle("/photo/*", mux)
	mux.Use(middleware.SubRouter)

	mux.Get("/", pH.Handle)
	mux.Get("/:id", pH.HandleSingleGet)
	mux.Post("/", pH.HandleCreate)

	parentMux.Get("/photo", http.RedirectHandler("/photo/", 301))

	fmt.Printf("Added photo router...\n")
}

// Handle processes a request for a set of photos.
func (pH *PhotoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered Handle\n")
}

// HandleCreate handles photo creation
func (pH *PhotoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered HandleCreate\n")

	p := &dtos.Photo{}
	json.NewDecoder(r.Body).Decode(p)

	pH.photoController.Create(p)

	pH.writeSuccessResponse(p, 201, w)
}

// HandleSingleGet processes a request for a single photos.
func (pH *PhotoHandler) HandleSingleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	fmt.Printf("Entered HandleSingleGet with id '%s'.\n", id)

	p := pH.photoController.Get(id)

	pH.writeSuccessResponse(p, 200, w)
}
