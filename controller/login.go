package controller

import (
	"net/url"
	"github.com/gin-gonic/gin"
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
	c.Request.ParseForm()
	userID := c.Request.Form["userid"][0]
	password := c.Request.Form["password"][0]
	if userID == "" || password == "" {
		redirectWithError(c, "login", "入力に不備があります")
		return
	}

	c.Redirect(302, "")
}


func redirectWithError(c *gin.Context, location, message string) {
	encodedMessage := url.QueryEscape(message)
	c.Redirect(302, location+"?error="+encodedMessage)
}