package controller

import (
	"github.com/gin-gonic/gin"
)


func RegisterPageController(c *gin.Context) {
	c.HTML(200, "register", gin.H{})
}

func RegisterPagePostController(c *gin.Context) {
	c.Redirect(302, "/")
}