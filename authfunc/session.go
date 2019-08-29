package authfunc

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"server-manage/db"
	"server-manage/common"
)

// InitSession : 新規セッションを作成して保存する
func InitSession(c *gin.Context, userID string) {
	passphrase := common.GenPassphrase(30)
	db.InsertSessionDataToDB(userID, passphrase, int(time.Now().Unix()+int64(3600)))
	session := sessions.Default(c)
	session.Set("UserID", userID)
	session.Set("Passphrase", passphrase)
	session.Save()
}
