package models

import (
	"miagi/lib/common"

	"github.com/jinzhu/gorm"
)

// Stadium data model
type DeviceToken struct {
	gorm.Model
	DeviceID string `gorm:"type:longtext"`
	FCMID    string `gorm:"type:longtext"`
	Status   uint
}

// Serialize serializes user data
func (deviceToken *DeviceToken) Serialize() common.JSON {
	return common.JSON{
		"id":        deviceToken.ID,
		"device_id": deviceToken.DeviceID,
		"fcm_id":    deviceToken.FCMID,
		"status":    deviceToken.Status,
	}
}

func (deviceToken *DeviceToken) Read(m common.JSON) {
	deviceToken.ID = uint(m["id"].(float64))
	deviceToken.DeviceID = m["device_id"].(string)
	deviceToken.FCMID = m["fcm_id"].(string)
	deviceToken.Status = m["status"].(uint)
}
