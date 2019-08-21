package view

import (
	"github.com/gin-gonic/gin"
)

// TopView : Topのビューを設定する
func TopView(c *gin.Context) {
	c.HTML(200, "top", gin.H{
		"pageTitle": "SMS Top",
	})
}
