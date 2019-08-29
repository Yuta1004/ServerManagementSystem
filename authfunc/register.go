package authfunc

import (
	"server-manage/db"
)

func Register(userID, password, passwordConf string) bool {
	// 入力チェック
	if password != passwordConf {
		return false
	}

	// ハッシュ化 → 登録
	hashPassword := GenPasswordHash(password)
	result := db.InsertUserDataToDB(userID, hashPassword)
	return result
}