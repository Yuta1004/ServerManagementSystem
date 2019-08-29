package controller

import (
	"github.com/gin-gonic/gin"
	"server-manage/authfunc"
)

// TopPageController : Topページのコントローラー
func TopPageController(c *gin.Context) {
	nowLogin, _ := authfunc.CheckOKSession(c)
	c.HTML(200, "top", gin.H{
		"pageTitle": "SMS Top",
		"nowLogin": nowLogin,
	})
}
