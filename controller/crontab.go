package controller

import (
	"server-manage/crontabmanage"

	"github.com/gin-gonic/gin"
)

// CrontabPageController : Crontabページのコントローラー
func CrontabPageController(c *gin.Context) {
	c.HTML(200, "crontab", gin.H{
		"crontabArray": crontabmanage.GetLatestCrontabData(),
	})
}
