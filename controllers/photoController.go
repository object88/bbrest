package controllers

import (
	"fmt"

	"github.com/object88/bbrest/access"
	"github.com/object88/bbrest/dtos"
	"github.com/object88/bbrest/models"

	"gopkg.in/mgo.v2"
)

const photoCollectionName string = "photos"

// IPhotoController is...
type IPhotoController interface {
	Create(d *dtos.Photo) *dtos.Photo
	Get(id string) *dtos.Photo
	Modify(id string, patch string) error
}

// PhotoController dictates CRUD operations
type PhotoController struct {
	BaseController
	pA *access.Access
}

// NewPhotoController instantiates a new instance of the controller
func NewPhotoController(s *mgo.Session, databaseName string) *PhotoController {
	pA := access.NewAccess(s, databaseName)
	return &PhotoController{BaseController{s, databaseName, photoCollectionName}, pA}
}

// Create accepts a Photo and places it in the repository
func (pC *PhotoController) Create(p *dtos.Photo) *dtos.Photo {
	fmt.Printf("Inserting photo '%s' into repository...\n", p)

	photo := models.NewPhoto()
	photo.FromDto(p)

	fmt.Printf("Converted to '%s'", photo)

	err := pC.pA.Save(photo)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	result := photo.ToDto()

	fmt.Printf("Inserted photo '%s' into repository.\n", result)

	return result
}

// Get returns a Photo with the matching id
func (pC *PhotoController) Get(id string) *dtos.Photo {
	fmt.Printf("Requesting photo with id '%s'...\n", id)

	var p models.Photo
	pM := p.GetMetadata()
	m, err := pC.pA.Get(pM, id)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	p = m.(models.Photo)

	result := p.ToDto()

	fmt.Printf("Request for photo with id '%s' complete.\n", id)

	return result
}

// Modify takes a thing and makes it another thing.
func (pC *PhotoController) Modify(id string, patch string) error {
	fmt.Printf("Modifying photo with id '%s'...\n", id)

	var p models.Photo
	pM := p.GetMetadata()
	m, err := pC.pA.Get(pM, id)
	if err != nil {
		return err
	}

	p = m.(models.Photo)

	return pC.pA.Save(p)
}
