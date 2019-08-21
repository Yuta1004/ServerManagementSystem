package view

import (
	"server-manage/crontabmanage"

	"github.com/gin-gonic/gin"
)

// CrontabView : Crontabページのビューを設定する
func CrontabView(c *gin.Context) {
	c.HTML(200, "crontab", gin.H{
		"crontabArray": crontabmanage.GetLatestCrontabData(),
	})
}
