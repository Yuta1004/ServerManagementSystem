package db

import (
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	userInfoList := *GetUserDataFromDB()
	if len(userInfoList) != 0 {
		t.Fail()
	}
}