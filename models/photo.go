package models

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// Photo is used.
type Photo struct {
	Title  string        `json:"title" bson:"title"`
	Author string        `json:"author" bson:"author"`
	ISBN   string        `json:"isbn" bson:"isbn"`
	Genre  string        `json:"genre" bson:"genre"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}

func (p *Photo) String() string {
	uj, _ := json.MarshalIndent(p, "", "\t")
	s := string(uj)
	return s
}
