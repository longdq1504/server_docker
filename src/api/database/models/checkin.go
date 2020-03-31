package models

import (
	"github.com/jinzhu/gorm"
	"miagi/lib/common"
)

// Stadium data model
type CheckIn struct {
	gorm.Model
	Image   string `gorm:"type:longtext"`
	Lng     float64
	Lat     float64
	User    User `gorm:"foreignkey:OwnerID"`
	OwnerID uint `gorm:"column:owner_id" json:"owner_id"`
}

// Serialize serializes user data
func (checkIn *CheckIn) Serialize() common.JSON {
	return common.JSON{
		"id":       checkIn.ID,
		"image":    checkIn.Image,
		"lng":      checkIn.Lng,
		"lat":      checkIn.Lat,
		"owner_id": checkIn.OwnerID,
	}
}

func (checkIn *CheckIn) Read(m common.JSON) {
	checkIn.ID = uint(m["id"].(float64))
	checkIn.Image = m["name"].(string)
	checkIn.Lng = m["lng"].(float64)
	checkIn.Lat = m["lat"].(float64)
	checkIn.OwnerID = m["owner_id"].(uint)
}
