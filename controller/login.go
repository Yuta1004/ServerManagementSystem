package controller

import (
	"github.com/gin-gonic/gin"
)

// LoginPageController : ログイン画面のコントローラー
func LoginPageController(c *gin.Context) {
	c.HTML(200, "login", gin.H{})
}