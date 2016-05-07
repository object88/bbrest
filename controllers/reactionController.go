package controllers

import (
	"fmt"

	"github.com/object88/bbrest/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const reactionCollectionName string = "photos"

// ReactionController dictates CRUD operations
type ReactionController struct {
	Controller
}

// NewReactionController instantiates a new instance of the controller
func NewReactionController(s *mgo.Session, databaseName string) *ReactionController {
	return &ReactionController{Controller{s, reactionCollectionName, databaseName}}
}

// Create accepts a Photo and places it in the repository
func (rC *ReactionController) Create(r *models.Reaction) *models.Reaction {
	fmt.Printf("Inserting reaction '%s' into repository...\n", r)

	r.ID = bson.NewObjectId()

	err := rC.GetCollection().Insert(r)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	fmt.Printf("Inserted reaction '%s' into repository.\n", r)

	return r
}
