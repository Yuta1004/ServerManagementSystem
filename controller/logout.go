package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"server-manage/authfunc"
)

func LogoutPageController(c *gin.Context) {
	session := sessions.Default(c)
	loginNow := session.Get("UserID") != nil
	authfunc.Logout(c)
	c.HTML(200, "logout", gin.H{
		"pageTitle": "SMS-Logout",
		"loginNow": loginNow,
	})
}