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

// CheckOKSession : 有効なセッションかチェックして結果を返す
func CheckOKSession(c *gin.Context) (bool, string) {
	session := sessions.Default(c)
	userID := session.Get("UserID")
	passphrase := session.Get("Passphrase")

	// nilチェック
	if userID == nil || passphrase == nil {
		return false, "NoLogin"
	}

	// 有効期限チェック
	userIDStr := userID.(string)
	sessionDBData := (*db.GetSessionDataFromDB(userIDStr))[0]
	if time.Now().Unix() > int64(sessionDBData.ExpirationUnixTime) {
		return false, "OverExpirationDate"
	}

	// パスフレーズチェック
	passphraseStr := passphrase.(string)
	if passphraseStr != sessionDBData.Passphrase {
		return false, "FraudSession"
	}

	return true, "OK"
}