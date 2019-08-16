package view

import (
	"github.com/gin-gonic/gin"
)

func TopView(c *gin.Context) {
	c.HTML(200, "top", gin.H{
		"pageTitle": "SMS Top",
	})
}
