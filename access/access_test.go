package access

import (
	"fmt"
	"testing"

	"github.com/object88/bbrest/models"
	"gopkg.in/mgo.v2"
)

func TestSave(t *testing.T) {
	s, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		t.Fail()
	}
	a := NewAccess(s, "brighterblacker")

	p := &models.Photo{}
	p.OwnerName = "Sam"
	fmt.Printf("New photo:\n%s\n", p)

	var result models.Modeler
	result, err = a.Save(p)
	if err != nil {
		t.Logf("Recieved error from save: '%s'.", err)
		t.Fail()
	}
	if result == nil {
		t.Log("No object returned from Save")
		t.Fail()
	}
	fmt.Print("Test complete.\n")
}

func TestGet(t *testing.T) {
	s, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		t.Fail()
	}
	a := NewAccess(s, "brighterblacker")

	p := &models.Photo{}
	p.OwnerName = "Fred"

	var result models.Modeler
	result, err = a.Save(p)
	if err != nil {
		t.Errorf("Failed to save photo: %s\n", err)
	}
	fmt.Printf("Save completed, '%s'...\n", result)

	id := result.GetID()
	fmt.Printf("Creating photo with id '%s'\n", id)

	result, err = a.Get(result.GetMetadata(), id)
	if err != nil {
		t.Fail()
	}
	if result == nil {
		t.Error("Did not recieve a model from Access.Get")
	}

	fmt.Printf("result: %s\n", result)

	resultPhoto := result.(models.Photo)
	if resultPhoto.GetID() != id {
		t.Fail()
	}
}
