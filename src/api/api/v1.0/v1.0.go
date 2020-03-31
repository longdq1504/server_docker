package apiv1

import (
	"github.com/gin-gonic/gin"
	"miagi/api/v1.0/auth"
	checkins "miagi/api/v1.0/checkin"
	deviceTokens "miagi/api/v1.0/devicetoken"
	divisions "miagi/api/v1.0/division"
	users "miagi/api/v1.0/user"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		users.ApplyRoutes(v1)
		checkins.ApplyRoutes(v1)
		divisions.ApplyRoutes(v1)
		deviceTokens.ApplyRoutes(v1)
	}
}
