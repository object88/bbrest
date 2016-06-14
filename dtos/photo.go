package dtos

import (
	"encoding/json"
	"time"
)

// Photo is used.
type Photo struct {
	ID             string          `json:"id"`
	OwnerID        string          `json:"ownerId"`
	OwnerName      string          `json:"ownerName"`
	Favorited      bool            `json:"favorited"`
	CameraSettings *CameraSettings `json:"cameraSettings"`
	UploadedOn     time.Time       `json:"uploadedOn"`
}

// CameraSettings is used
type CameraSettings struct {
	Fstop   float32 `json:"fstop"`
	Shutter int     `json:"shutter"`
	ISO     int     `json:"iso"`
}

func (p *Photo) String() string {
	uj, _ := json.Marshal(p)
	s := string(uj)
	return s
}

/*
{
  id: "abc012",
  owner: {
    id: "def345",
    name: "bob roberts"
  },
  favorited: true,
  venue: {

  },
  event: {

  },
  date: [date],
  tags: ["a", "b", "c"]
}
*/
