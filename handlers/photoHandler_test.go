package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/object88/bbrest/dtos"
	. "github.com/smartystreets/goconvey/convey"
)

type MockPhotoController struct {
}

func (m *MockPhotoController) Create(d *dtos.BaseDto) *dtos.BaseDto {
	return nil
}
func (m *MockPhotoController) Get(id string) *dtos.BaseDto {
	return nil
}

func TestCreate(t *testing.T) {
	Convey("", t, func() {
		m := &MockPhotoController{}
		pH := AddPhotoHandler(m, nil)

		// w http.ResponseWriter, r *http.Request
		req, _ := http.NewRequest("GET", "", nil)
		w := httptest.NewRecorder()

		pH.HandleCreate(req, &w)

	})
}
