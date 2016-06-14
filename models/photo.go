package models

import (
	"fmt"
	"time"

	"github.com/object88/bbrest/dtos"
	"github.com/object88/bbrest/singleton"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PhotoMetadata is...
type photoMetadata struct {
}

// Photo is used.
type Photo struct {
	ID         bson.ObjectId `bson:"_id"`
	OwnerID    bson.ObjectId `bson:"ownerId,omitempty"`
	OwnerName  string        `bson:"ownerName"`
	UploadedOn time.Time     `bson:"uploadedOn"`
	// ID         bson.ObjectId
	// OwnerID    bson.ObjectId
	// OwnerName  string
	// UploadedOn time.Time
}

var photoMetadataSingleton *photoMetadata
var once singleton.Once

// GetID does...
func (p Photo) GetID() string {
	return p.ID.Hex()
}

// Read is...
func (Photo) Read(c *mgo.Collection, oid *bson.ObjectId) (Modeler, error) {
	p := Photo{}

	query := c.Find(bson.M{"_id": oid})
	err := query.One(&p)
	if err != nil {
		return nil, err
	}

	if p.ID != *oid {
		return nil, nil
	}

	return p, nil
}

// Create is...
func (p Photo) Create(c *mgo.Collection) (Modeler, error) {
	p.ID = bson.NewObjectId()
	p.UploadedOn = time.Now().UTC()

	err := c.Insert(p)
	return p, err
}

// Update is...
func (p Photo) Update(c *mgo.Collection) (Modeler, error) {

	err := c.Update(bson.M{"_id": p.ID}, p)
	return p, err
}

// // GetOwnerID returns...
// func (p *Photo) GetOwnerID() string {
// 	return p.ownerName
// }
//
// // GetOwnerName returns...
// func (p *Photo) GetOwnerName() string {
// 	return p.ownerName
// }
//
// // GetUploadedOn does...
// func (p *Photo) GetUploadedOn() time.Time {
// 	return p.uploadedOn
// }

// GetMetadata is...
func (Photo) GetMetadata() Metadatar {
	once.Do(func() {
		photoMetadataSingleton = &photoMetadata{}
	})
	return photoMetadataSingleton
}

func (p *Photo) String() string {
	s := fmt.Sprintf("Photo: id='%s', ownername='%s', uploadedOn='%s'\n", p.ID.Hex(), p.OwnerName, p.UploadedOn)
	// uj, _ := json.MarshalIndent(p, "", "\t")
	// s := string(uj)
	// fmt.Printf("Serialized Photo with id '%s' and uploaded date %s into\n%s\n", p.id, p.uploadedOn, s)
	return s
}

// FromDto ...
func (p *Photo) FromDto(source *dtos.Photo) {
	p.ID = bson.ObjectIdHex(source.ID)
	p.OwnerID = bson.ObjectIdHex(source.OwnerID)
	p.OwnerName = source.OwnerName
	p.UploadedOn = source.UploadedOn
}

// ToDto converts a Photo model into a Photo DTO
func (p *Photo) ToDto() *dtos.Photo {
	result := dtos.Photo{
		ID:         p.ID.Hex(),
		OwnerID:    p.OwnerID.Hex(),
		OwnerName:  p.OwnerName,
		UploadedOn: p.UploadedOn,
	}
	return &result
}

// GetBSON serializes...
// func (p Photo) GetBSON() (interface{}, error) {
// 	fmt.Printf("Translating photo '%s' to BSON structure.\n", p.GetID())
// 	return struct {
// 		ID         bson.ObjectId `bson:"_id"`
// 		OwnerID    bson.ObjectId `bson:"ownerId,omitempty"`
// 		OwnerName  string        `bson:"ownerName"`
// 		UploadedOn time.Time     `bson:"uploadedOn"`
// 	}{
// 		ID:         p.ID,
// 		OwnerID:    p.OwnerID,
// 		OwnerName:  p.OwnerName,
// 		UploadedOn: p.UploadedOn,
// 	}, nil
// }
//
// // SetBSON deserializes...
// func (p *Photo) SetBSON(raw bson.Raw) error {
// 	decoded := new(struct {
// 		ID         bson.ObjectId `bson:"_id"`
// 		OwnerID    bson.ObjectId `bson:"ownerId,omitempty"`
// 		OwnerName  string        `bson:"ownerName"`
// 		UploadedOn time.Time     `bson:"uploadedOn"`
// 	})
//
// 	fmt.Printf("Received %d bytes to decode.\n", len(string(raw.Data)))
// 	bsonErr := raw.Unmarshal(decoded)
//
// 	if bsonErr != nil {
// 		fmt.Printf("Received error while unmarshaling: '%s'\n", bsonErr)
// 		return bsonErr
// 	}
//
// 	fmt.Printf("Decoded: %s\n", decoded)
// 	p.ID = decoded.ID
// 	p.OwnerID = decoded.OwnerID
// 	p.OwnerName = decoded.OwnerName
// 	p.UploadedOn = decoded.UploadedOn
// 	return nil
// }

// GetCollectionName is...
func (photoMetadata) GetCollectionName() string {
	return "photos"
}

func (photoMetadata) Instantiate() interface{} {
	return &Photo{}
}
