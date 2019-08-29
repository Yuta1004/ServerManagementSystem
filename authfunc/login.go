package authfunc

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"server-manage/db"
	"server-manage/common"
)

// Login : ログイン処理
func Login(c *gin.Context, userID, password string) bool {
	// ユーザ存在チェック
	fetchResult := db.GetUserDataFromDB(userID)
	if len(*fetchResult) == 0 {
		return false
	}

	// パスワードチェック
	dbPassword := (*fetchResult)[0].HashPassword
	checkResult := AuthPassword(password, dbPassword)

	// セッション管理
	if checkResult {
		passphrase := common.GenPassphrase(30)
		db.InsertSessionDataToDB(userID, passphrase, int(time.Now().Unix()+int64(3600)))
		session := sessions.Default(c)
		session.Set("UserID", userID)
		session.Set("Passphrase", passphrase)
		session.Save()
	}
	return checkResult
}