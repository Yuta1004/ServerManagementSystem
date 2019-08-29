package controller

import (
	"github.com/gin-gonic/gin"
    "server-manage/manage/crontab"
)

// CrontabPageController : Crontabページのコントローラー
func CrontabPageController(c *gin.Context) {
	c.HTML(200, "crontab", gin.H{
		"pageTitle": "SMS-Manage-Crontab",
		"crontabArray": crontab.GetLatestCrontabData(),
	})
}
