package middlewares

import (
	"miagi/database/models"

	"github.com/gin-gonic/gin"
)

type User = models.User

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
}

func IsAdmin(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists || (user.(User)).RoleID != 1 {
		c.AbortWithStatus(401)
		return
	}
}
