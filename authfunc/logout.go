package authfunc

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}