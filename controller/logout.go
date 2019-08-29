package controller

import (
	"github.com/gin-gonic/gin"
	"server-manage/authfunc"
)

func LogoutPageController(c *gin.Context) {
	checkLogin, _ := authfunc.CheckOKSession(c)
	authfunc.Logout(c)
	c.HTML(200, "logout", gin.H{
		"pageTitle": "SMS-Logout",
		"loginNow": checkLogin,
	})
}