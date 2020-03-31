package divisions

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	divisions := r.Group("/division")
	{
		divisions.GET("/", list)
		divisions.GET("/:id", read)
	}
}
