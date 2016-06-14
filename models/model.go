package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Modeler is...
type Modeler interface {
	Creater
	Reader
	Updater
	// Deleter
	GetID() string
	GetMetadata() Metadatar
}

// Creater is...
type Creater interface {
	Create(*mgo.Collection) (Modeler, error)
}

// Reader is...
type Reader interface {
	Read(*mgo.Collection, *bson.ObjectId) (Modeler, error)
}

// Updater is...
type Updater interface {
	Update(*mgo.Collection) (Modeler, error)
}

// Deleter is...
type Deleter interface {
	Delete()
}

// Metadatar is...
type Metadatar interface {
	GetCollectionName() string
	Instantiate() interface{}
}
