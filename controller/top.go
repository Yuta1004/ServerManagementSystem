package controller

import (
	"github.com/gin-gonic/gin"
)

// TopPageController : Topページのコントローラー
func TopPageController(c *gin.Context) {
	c.HTML(200, "top", gin.H{
		"pageTitle": "SMS Top",
	})
}
