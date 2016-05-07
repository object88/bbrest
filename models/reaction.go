package models

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// Reaction is unused.
type Reaction struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Photo    bson.ObjectId `json:"photoId" bson:"photoId"`
	User     bson.ObjectId `json:"userId" bson:"userId"`
	Favorite bool          `json:"favorite" bson:"favorite"`
	Reaction int           `json:"reaction" bson:"reaction"`
}

func (p *Reaction) String() string {
	uj, _ := json.MarshalIndent(p, "", "\t")
	s := string(uj)
	return s
}
