package main

import (
	"fmt"
	"miagi/api"
	"miagi/database"

	"miagi/lib/middlewares"
	notification "miagi/lib/notification"
	"miagi/web"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func task() {
	fmt.Println("I am running task.")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app

	go func() {
		fmt.Println("app")
		gocron.Every(5).Minutes().DoSafely(notification.SendCheckInFCM, db, map[string]string{})
		// gocron.Every(5).Seconds().DoSafely(common.SendCheckInFCM, db, []string{"eNNFmY9FRdyUFqdNXypkJt:APA91bFZRgRixT-xwAf9QSQSWZTzuhGs5Gv63AwzAsQLvmK2ibB-wnl4bLr_FIBR7f0NyZSEtHBGhWC4KMemW81EKQmF-1lohoKcwBIvzcbqKuWsEfXMJ1L6ESxo3ZGXs0givs7Rvkym"})
		// gocron.Every(1).Monday().At("0:55").DoSafely(common.SendCheckInFCM, db, nil)
		// gocron.Every(1).Tuesday().At("0:55").DoSafely(common.SendCheckInFCM, db, nil)
		// gocron.Every(1).Wednesday().At("0:55").DoSafely(common.SendCheckInFCM, db, nil)
		// gocron.Every(1).Thursday().At("0:55").DoSafely(common.SendCheckInFCM, db, nil)
		// gocron.Every(1).Friday().At("0:55").DoSafely(common.SendCheckInFCM, db, nil)
		<-gocron.Start()
	}()

	app.Use(cors.Default())

	app.LoadHTMLGlob("web/template/*")
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	web.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port

}
