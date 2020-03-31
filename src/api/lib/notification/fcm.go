package notification

import (
	"fmt"
	models "miagi/database/models"

	fcm "github.com/NaySoftware/go-fcm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DeviceToken = models.DeviceToken

const (
	serverKey = "AAAABQ8oTho:APA91bHnHZ6ePOuhsK6BnoTPUguOAMHzypTOoSXfTKFsnWnmi_qRagWyI7EEvU_WHpYAfnl9E3GtR1ROLNiPmURY9ZzSedmJvz87c61j3WYCEw2qbrEiyG0ENchc7ndrW__YcrJNfQvO"
)

func TestFCM(title string, body string) {
	fmt.Println(title, body)
}

func SendDemoFCM(ids []string) {
	var NP fcm.NotificationPayload
	NP.Title = "From Miagi with LOVE"
	NP.Body = "Sắp đến giờ làm việc rồi. Bạn nhớ check-in nhé!"

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.SetNotificationPayload(&NP)
	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}

func SendFCM(ids []string,
	title string,
	body string,
	data map[string]string) {
	var NP fcm.NotificationPayload
	NP.Title = title
	NP.Body = body

	xds := []string{
		"tokens",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)
	c.SetNotificationPayload(&NP)
	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}

func SendCheckInFCM(db *gorm.DB, data map[string]string) {
	fmt.Println("Push notification")
	var deviceTokens []DeviceToken

	if err := db.Order("id desc").Find(&deviceTokens).Error; err != nil {
		return
	}

	var ids []string
	if len(deviceTokens) > 0 {
		for _, token := range deviceTokens {
			fmt.Println(token.FCMID)
			ids = append(ids, token.FCMID)
		}
	}

	var NP fcm.NotificationPayload
	NP.Title = "From Miagi with LOVE"
	NP.Body = "Sắp đến giờ làm việc rồi. Bạn nhớ check-in nhé!"

	xds := []string{
		"tokens",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)
	c.SetNotificationPayload(&NP)
	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
