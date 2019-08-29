package controller

import (
	"github.com/gin-gonic/gin"
	"server-manage/authfunc"
)


func RegisterPageController(c *gin.Context) {
	errorMsg := c.Query("error")
	c.HTML(200, "register", gin.H{
		"error": errorMsg,
	})
}

func RegisterPagePostController(c *gin.Context) {
	// 入力確認
	c.Request.ParseForm()
	userID := c.Request.Form["userid"][0]
	password := c.Request.Form["password"][0]
	passwordConf := c.Request.Form["password-conf"][0]
	if userID == "" || password == "" || passwordConf == "" {
		redirectWithError(c, "register", "InvalidInput")
		return
	}

	// 登録処理
	regResult := authfunc.Register(userID, password, passwordConf)
	if !regResult {
		redirectWithError(c, "register", "WrongInput")
		return
	}

	authfunc.Login(c, userID, password)
	c.Redirect(302, "")
	c.Next()
}