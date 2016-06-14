package access

import (
	"fmt"
	"reflect"

	"github.com/object88/bbrest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Access is...
type Access struct {
	s  *mgo.Session
	db string
}

// NewAccess is...
func NewAccess(s *mgo.Session, databaseName string) *Access {
	return &Access{s, databaseName}
}

// Get is...
func (a *Access) Get(m models.Metadatar, id string) (models.Modeler, error) {
	collectionName := m.GetCollectionName()
	oid := bson.ObjectIdHex(id)

	fmt.Printf("Getting '%s' from '%s'\n", id, collectionName)

	c := getCollection(a.s, a.db, collectionName)

	i := m.Instantiate().(models.Reader)

	// x := reflect.TypeOf(i)
	// fmt.Printf("model type:\n%s\n", x)

	model, err := i.Read(c, &oid)

	if err != nil {
		fmt.Printf("Failed: '%s'\n", err)
		return nil, err
	}

	if model == nil {
		return nil, nil
	}

	x2 := reflect.TypeOf(model)
	fmt.Printf("model type: %s\n", x2)

	return model, nil
}

// Save ...
func (a *Access) Save(m models.Modeler) (models.Modeler, error) {
	c := getCollection(a.s, a.db, m.GetMetadata().GetCollectionName())

	var err error
	m, err = m.Create(c)

	return m, err
}

// Update is...
func (a *Access) Update(m models.Modeler) error {
	c := getCollection(a.s, a.db, m.GetMetadata().GetCollectionName())
	id := m.GetID()

	err := c.UpdateId(id, m)
	return err
}

func getCollection(s *mgo.Session, dbName string, collectionName string) *mgo.Collection {
	collection := s.DB(dbName).C(collectionName)
	return collection
}
