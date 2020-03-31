package checkins

import (
	"miagi/lib/middlewares"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	checkins := r.Group("/checkin")
	{
		checkins.POST("/", middlewares.Authorized, create)
		checkins.GET("/", list)
		checkins.GET("/:id", read)
		checkins.DELETE("/:id", middlewares.IsAdmin, remove)
		checkins.PATCH("/:id", middlewares.IsAdmin, update)
	}
}
