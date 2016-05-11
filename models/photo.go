package models

import (
	"encoding/json"
	"time"

	"github.com/object88/bbrest/dtos"

	"gopkg.in/mgo.v2/bson"
)

// Photo is used.
type Photo struct {
	ID         bson.ObjectId `bson:"_id"`
	OwnerID    bson.ObjectId `bson:"ownerId,omitempty"`
	OwnerName  string        `bson:"ownerName"`
	Favorited  bool          `bson:"favorited"`
	UploadedOn time.Time     `bson:"uploadedOn"`
}

func (p *Photo) String() string {
	uj, _ := json.MarshalIndent(p, "", "\t")
	s := string(uj)
	return s
}

// ToDto converts a Photo model into a Photo DTO
func (p *Photo) ToDto() *dtos.Photo {
	result := dtos.Photo{
		BaseDto:    dtos.BaseDto{ID: p.ID},
		OwnerID:    p.OwnerID,
		OwnerName:  p.OwnerName,
		Favorited:  p.Favorited,
		UploadedOn: p.UploadedOn,
	}
	return &result
}

// FromDto ...
func (p *Photo) FromDto(source *dtos.Photo) *Photo {
	p.ID = source.ID
	p.OwnerID = source.OwnerID
	p.OwnerName = source.OwnerName
	p.Favorited = source.Favorited
	p.UploadedOn = source.UploadedOn
	return p
}
