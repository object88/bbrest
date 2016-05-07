package dtos

import "gopkg.in/mgo.v2/bson"

// BaseDto is used.
type BaseDto struct {
	ID bson.ObjectId `json:"id"`
}
