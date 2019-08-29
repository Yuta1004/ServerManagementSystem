package db

import (
	"log"
	"testing"
	"golang.org/x/crypto/bcrypt"
)

func TestControllUserDB(t *testing.T) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("servermanage"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Faild to hash the password.")
		return
	}

	result := InsertUserDataToDB("testuser", string(hashPassword))
	if result {
		t.Fail()	// すでにtestuserが挿入された状態でテストを行うこと!
	}

	userInfoList := *GetUserDataFromDB("testuser")
	if len(userInfoList) != 1 {
		t.Fail()
	}
	userInfoList = *GetUserDataFromDB("sjklfjsdfalfjsladkfj")
	if len(userInfoList) != 0 {
		t.Fail()
	}
}