package divisions

import (
	"miagi/database/models"
	"miagi/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Division = models.Division
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var divisions []Division

	if err := db.Order("id desc").Find(&divisions).Error; err != nil {
		c.AbortWithStatus(500)
		return
	}

	length := len(divisions)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = divisions[i].Serialize()
	}

	c.JSON(200, common.GenerateResponse(0, "", common.JSON{
		"divisions": serialized,
		"total":     len(serialized),
	}))
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var division Division

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&division).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, common.GenerateResponse(0, "", division.Serialize()))
}
