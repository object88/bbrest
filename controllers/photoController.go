package controllers

import (
	"fmt"
	"time"

	"github.com/object88/bbrest/dtos"
	"github.com/object88/bbrest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const photoCollectionName string = "photos"

// IPhotoController is...
type IPhotoController interface {
	Create(d *dtos.Photo) *dtos.Photo
	Get(id string) *dtos.Photo
}

// PhotoController dictates CRUD operations
type PhotoController struct {
	BaseController
}

// NewPhotoController instantiates a new instance of the controller
func NewPhotoController(s *mgo.Session, databaseName string) *PhotoController {
	return &PhotoController{BaseController{s, databaseName, photoCollectionName}}
}

// Create accepts a Photo and places it in the repository
func (pC *PhotoController) Create(p *dtos.Photo) *dtos.Photo {
	fmt.Printf("Inserting photo '%s' into repository...\n", p)

	photo := (&models.Photo{}).FromDto(p)
	photo.ID = bson.NewObjectId()
	photo.UploadedOn = time.Now().UTC()

	fmt.Printf("Converted to '%s'", photo)

	err := pC.GetCollection().Insert(photo)
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
	oid := bson.ObjectIdHex(id)

	fmt.Printf("Requesting photo with id '%s'...\n", id)

	photo := &models.Photo{}
	query := pC.GetCollection().FindId(oid)
	err := query.One(photo)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	result := &dtos.Photo{
		BaseDto:        dtos.BaseDto{ID: photo.ID},
		OwnerID:        bson.NewObjectId(),
		OwnerName:      "Bob Roberts",
		Favorited:      false,
		CameraSettings: nil,
	}

	fmt.Printf("Request for photo with id '%s' complete.\n", id)

	return result
}
