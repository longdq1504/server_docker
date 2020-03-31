package deviceTokens

import (
	"fmt"
	"miagi/database/models"
	"miagi/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DeviceToken = models.DeviceToken
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		DeviceID string `json:"device_id" binding:"required"`
		FCMID    string `json:"fcm_id" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	var existedDeviceToken []DeviceToken
	db.Where("device_id = ?", requestBody.DeviceID).Find(&existedDeviceToken)

	if len(existedDeviceToken) > 0 {
		for _, token := range existedDeviceToken {
			db.Delete(&token)
		}
	}

	deviceToken := DeviceToken{
		DeviceID: requestBody.DeviceID,
		FCMID:    requestBody.FCMID,
		Status:   1,
	}

	db.NewRecord(deviceToken)
	db.Create(&deviceToken)
	c.JSON(200, common.GenerateResponse(0, "Đăng ký fcm token thành công!", deviceToken.Serialize()))
}
