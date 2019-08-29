package controller

import (
	"github.com/gin-gonic/gin"
	"server-manage/authfunc"
)

func LogoutPageController(c *gin.Context) {
	authfunc.Logout(c)
	c.HTML(200, "logout", gin.H{})
}