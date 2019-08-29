package controller

import (
	"server-manage/crontabmanage"

	"github.com/gin-gonic/gin"
)

// CrontabPageController : Crontabページのコントローラー
func CrontabPageController(c *gin.Context) {
	c.HTML(200, "crontab", gin.H{
		"pageTitle": "SMS-Manage-Crontab",
		"crontabArray": crontabmanage.GetLatestCrontabData(),
	})
}
