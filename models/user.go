package models

import "gopkg.in/mgo.v2/bson"

// User is...
type User struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}
