package checkins

import (
	"miagi/database/models"
	"miagi/lib/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CheckIn = models.CheckIn
type User = models.User
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Image string  `json:"image" binding:"required"`
		Lng   float64 `json:"lng" binding:"required"`
		Lat   float64 `json:"lat" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var existedCheckIn []CheckIn
	now := time.Now()
	user := c.MustGet("user").(User)
	db.Preload("User").Where("owner_id = ? and created_at >= ? and created_at <= ?", user.ID, common.Bod(now), common.Eod(now)).Find(&existedCheckIn)

	if len(existedCheckIn) > 0 {
		c.JSON(409, common.GenerateResponse(1, "Bạn đã check-in hôm nay!", nil))
		return
	}

	checkin := CheckIn{
		Image:   requestBody.Image,
		Lng:     requestBody.Lng,
		Lat:     requestBody.Lat,
		OwnerID: user.ID,
		User:    user,
	}

	if common.Distance(21.016298, 105.795210, checkin.Lat, checkin.Lng, "K") > 0.1 {
		c.JSON(400, common.GenerateResponse(1, "Hãy check-in khi bạn đã có mặt ở công ty!", nil))
	} else {
		db.NewRecord(checkin)
		db.Create(&checkin)
		c.JSON(200, common.GenerateResponse(0, "Check-in thành công!", checkin.Serialize()))
	}
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	offset := common.GetPage(c)
	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var checkins []CheckIn

	if cursor == "" {
		if err := db.Limit(common.PageSize).Offset(offset).Order("id desc").Find(&checkins).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Limit(common.PageSize).Offset(offset).Order("id desc").Where(condition, cursor).Find(&checkins).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(checkins)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = checkins[i].Serialize()
	}

	c.JSON(200, common.GenerateResponse(0, "", common.JSON{
		"checkins": serialized,
		"total":    len(serialized),
	}))
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var checkin CheckIn

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&checkin).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, checkin.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	var checkin CheckIn
	if err := db.Where("id = ?", id).First(&checkin).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if checkin.OwnerID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&checkin)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	type RequestBody struct {
		Image string  `json:"image" binding:"required"`
		Lng   float64 `json:"lng" binding:"required"`
		Lat   float64 `json:"lat" binding:"required"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var checkin CheckIn
	if err := db.Where("id = ?", id).First(&checkin).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if checkin.OwnerID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	checkin.Image = requestBody.Image
	checkin.Lng = requestBody.Lng
	checkin.Lat = requestBody.Lat
	db.Save(&checkin)
	c.JSON(200, checkin.Serialize())
}

func listAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	offset := common.GetPage(c)
	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var checkins []CheckIn

	if cursor == "" {
		if err := db.Limit(common.PageSize).Offset(offset).Order("id desc").Find(&checkins).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Limit(common.PageSize).Offset(offset).Order("id desc").Where(condition, cursor).Find(&checkins).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(checkins)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = checkins[i].Serialize()
	}

	c.JSON(200, common.GenerateResponse(0, "", common.JSON{
		"checkins": serialized,
		"total":    len(serialized),
	}))
}
