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
	photoController controllers.IPhotoController
}

// CreatePhotoHandler creates a new instance
func CreatePhotoHandler(pC controllers.IPhotoController) *PhotoHandler {
	pH := &PhotoHandler{Handler{}, pC}
	return pH
}

// AddPhotoHandler creates a new instance
func (pH *PhotoHandler) AddPhotoHandler(parentMux *web.Mux) {
	fmt.Printf("Adding photo router...\n")

	mux := web.New()
	parentMux.Handle("/photo/*", mux)
	mux.Use(middleware.SubRouter)

	mux.Get("/", pH.Handle)
	mux.Get("/:id", pH.HandleSingleGet)
	mux.Post("/", pH.HandleCreate)
	mux.Patch("/:id", pH.HandleModify)

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

	p := dtos.Photo{}
	json.NewDecoder(r.Body).Decode(&p)

	result := pH.photoController.Create(&p)

	uj, _ := json.Marshal(result)
	pH.writeSuccessResponse(uj, 201, w)
}

// HandleModify ...
func (pH *PhotoHandler) HandleModify(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Entered HandleModify\n")

	id := c.URLParams["id"]

	// p := dtos.Photo{}
	// json.NewDecoder(r.Body).Decode(&p)

	// r.Body

	err := pH.photoController.Modify(id, "")
	if err != nil {
		// Failure!
	}

	pH.writeSuccessResponse(nil, 200, w)
}

// HandleSingleGet processes a request for a single photos.
func (pH *PhotoHandler) HandleSingleGet(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	fmt.Printf("Entered HandleSingleGet with id '%s'.\n", id)

	p := pH.photoController.Get(id)

	uj, _ := json.Marshal(p)
	pH.writeSuccessResponse(uj, 200, w)
}
