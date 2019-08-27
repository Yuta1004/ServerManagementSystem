package controller

import (
	"github.com/gin-gonic/gin"
)

// LoginPageController : ログイン画面のコントローラー
func LoginPageController(c *gin.Context) {
	errMsg := c.Query("error")
	c.HTML(200, "login", gin.H{
		"error": errMsg,
	})
}