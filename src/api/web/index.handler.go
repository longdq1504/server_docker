package web

import (
	"github.com/gin-gonic/gin"
	"miagi/lib/common"
)

func showIndexPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "Miagi",
	}, "index.html")
}
