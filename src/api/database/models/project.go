package models

import (
	"miagi/lib/common"

	"github.com/jinzhu/gorm"
)

// Stadium data model
type Project struct {
	gorm.Model
	Name        string
	Description string `sql:"type:text;"`
	Type        uint
}

// Serialize serializes user data
func (project *Project) Serialize() common.JSON {
	return common.JSON{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"type":        project.Type,
	}
}

func (project *Project) Read(m common.JSON) {
	project.ID = uint(m["id"].(float64))
	project.Name = m["name"].(string)
	project.Description = m["description"].(string)
	project.Type = m["type"].(uint)
}
