package authfunc

import (
	"server-manage/db"
)

// Login : ログイン処理
func Login(userID, password string) bool {
	// ユーザ存在チェック
	fetchResult := db.GetUserDataFromDB(userID)
	if len(*fetchResult) == 0 {
		return false
	}

	// パスワードチェック
	dbPassword := (*fetchResult)[0].HashPassword
	return AuthPassword(password, dbPassword)
}