package models

import (
	"miagi/lib/common"

	"github.com/jinzhu/gorm"
)

// Stadium data model
type Division struct {
	gorm.Model
	Name string
	Rank uint
}

// Serialize serializes user data
func (division *Division) Serialize() common.JSON {
	return common.JSON{
		"id":   division.ID,
		"name": division.Name,
		"rank": division.Rank,
	}
}

func (division *Division) Read(m common.JSON) {
	division.ID = uint(m["id"].(float64))
	division.Name = m["name"].(string)
	division.Rank = m["rank"].(uint)
}
