package controllers

import (
	"fmt"

	"github.com/object88/bbrest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const photoCollectionName string = "photos"

// PhotoController dictates CRUD operations
type PhotoController struct {
	Controller
}

// NewPhotoController instantiates a new instance of the controller
func NewPhotoController(s *mgo.Session, databaseName string) *PhotoController {
	return &PhotoController{Controller{s, photoCollectionName, databaseName}}
}

// Create accepts a Photo and places it in the repository
func (pC *PhotoController) Create(p *models.Photo) *models.Photo {
	fmt.Printf("Inserting photo '%s' into repository...\n", p)

	p.ID = bson.NewObjectId()

	err := pC.GetCollection().Insert(p)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	fmt.Printf("Inserted photo '%s' into repository.\n", p)

	return p
}

// Get returns a Photo with the matching id
func (pC *PhotoController) Get(id string) *models.Photo {
	oid := bson.ObjectIdHex(id)

	fmt.Printf("Requesting photo with id '%s'...\n", id)

	result := &models.Photo{}
	query := pC.GetCollection().FindId(oid)
	err := query.One(result)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	fmt.Printf("Request for photo with id '%s' complete.\n", id)

	return result
}
