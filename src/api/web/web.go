package web

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine) {
	r.GET("/login", showLoginPage)
	r.GET("/", showIndexPage)
}
