package deviceTokens

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	deviceTokens := r.Group("/devicetoken")
	{
		deviceTokens.POST("/", create)
	}
}
