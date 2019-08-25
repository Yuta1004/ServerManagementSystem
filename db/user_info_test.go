package db

import (
	"testing"
	"server-manage/authfunc"
)

func TestControllUserDB(t *testing.T) {
	hashPassword := authfunc.GenPasswordHash("servermanage")
	result := InsertUserDataToDB("testuser", hashPassword)
	if result {
		t.Fail()	// すでにtestuserが挿入された状態でテストを行うこと!
	}

	userInfoList := *GetUserDataFromDB()
	if len(userInfoList) != 1 {
		t.Fail()
	}
	userInfoList = *GetUserDataFromDB("testuser")
	if len(userInfoList) != 1 {
		t.Fail()
	}
	userInfoList = *GetUserDataFromDB("sjklfjsdfalfjsladkfj")
	if len(userInfoList) != 0 {
		t.Fail()
	}
}