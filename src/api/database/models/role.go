package models

import (
	"miagi/lib/common"

	"github.com/jinzhu/gorm"
)

// Stadium data model
type Role struct {
	gorm.Model
	Name string
	Rank uint
}

// Serialize serializes user data
func (role *Role) Serialize() common.JSON {
	return common.JSON{
		"id":   role.ID,
		"name": role.Name,
		"rank": role.Rank,
	}
}

func (role *Role) Read(m common.JSON) {
	role.ID = uint(m["id"].(float64))
	role.Name = m["name"].(string)
	role.Rank = m["rank"].(uint)
}
