package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collectionName string = "photos"

// PhotoController dictates CRUD operations
type PhotoController struct {
	Controller
}

// NewPhotoController instantiates a new instance of the controller
func NewPhotoController(s *mgo.Session) *PhotoController {
	return &PhotoController{Controller{s, collectionName}}
}

// Create accepts a Photo and places it in the repository
func (pC *PhotoController) Create(p *Photo) *Photo {
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
func (pC *PhotoController) Get(id string) *Photo {
	oid := bson.ObjectIdHex(id)

	fmt.Printf("Requesting photo with id '%s'...\n", id)

	result := &Photo{}
	query := pC.GetCollection().FindId(oid)
	err := query.One(result)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	fmt.Printf("Request for photo with id '%s' complete.\n", id)

	return result
}
