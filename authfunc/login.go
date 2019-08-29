package authfunc

import (
	"github.com/gin-gonic/gin"
	"server-manage/db"
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
		InitSession(c, userID)
	}
	return checkResult
}