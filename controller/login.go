package controller

import (
	"net/url"
	"github.com/gin-gonic/gin"
	"server-manage/authfunc"
)

// LoginPageController : ログイン画面のコントローラー
func LoginPageController(c *gin.Context) {
	errMsg := c.Query("error")
	c.HTML(200, "login", gin.H{
		"error": errMsg,
	})
}

// LoginPagePostController : ログイン画面(POST)のコントローラー
func LoginPagePostController(c *gin.Context) {
	// 入力確認
	c.Request.ParseForm()
	userID := c.Request.Form["userid"][0]
	password := c.Request.Form["password"][0]
	if userID == "" || password == "" {
		redirectWithError(c, "login", "InvalidInput")
		return
	}

	// ログイン処理
	if authfunc.Login(c, userID, password) {
		c.Redirect(302, "")
	} else {
		redirectWithError(c, "login", "WrongInput")
		return
	}
	c.Next()
}

func redirectWithError(c *gin.Context, location, message string) {
	encodedMessage := url.QueryEscape(message)
	c.Redirect(302, location+"?error="+encodedMessage)
	c.Abort()
}
