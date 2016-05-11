package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/object88/bbrest/dtos"
	. "github.com/smartystreets/goconvey/convey"
)

type MockPhotoController struct {
	createCalled bool
	getCalled    bool
}

func (m *MockPhotoController) Create(d *dtos.Photo) *dtos.Photo {
	m.createCalled = true
	result := &dtos.Photo{
		BaseDto:        dtos.BaseDto{ID: bson.NewObjectId()},
		OwnerID:        d.OwnerID,
		OwnerName:      d.OwnerName,
		Favorited:      d.Favorited,
		CameraSettings: d.CameraSettings,
		UploadedOn:     time.Now().UTC(),
	}
	return result
}

func (m *MockPhotoController) Get(id string) *dtos.Photo {
	m.getCalled = true
	return nil
}

func TestCreate(t *testing.T) {
	Convey("Can create", t, func() {
		m := &MockPhotoController{}
		pH := CreatePhotoHandler(m)

		body := `{"ownerName":"Bob"}`
		buffer := bytes.NewBufferString(body)
		req, _ := http.NewRequest("POST", "/photo", buffer)
		w := httptest.NewRecorder()

		pH.HandleCreate(w, req)

		So(w.Code, ShouldEqual, 201)

		result := &dtos.Photo{}
		json.NewDecoder(w.Body).Decode(result)
		So(result.OwnerName, ShouldEqual, "Bob")

		So(m.createCalled, ShouldEqual, true)
	})
}
