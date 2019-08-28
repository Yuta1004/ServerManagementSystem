package db

import (
	"testing"
)

func TestControllSessionDB(t *testing.T) {
	userID := "testuser"
	passphrase := "testphrase"
	expirationUnixTime := 0
	InsertSessionDataToDB(userID, passphrase, expirationUnixTime)

	sessionInfo := (*GetSessionDataFromDB(userID))[0]
	if sessionInfo.ID != userID {
		t.Fail()
	}
	if sessionInfo.Passphrase != passphrase {
		t.Fail()
	}
	if sessionInfo.ExpirationUnixTime != expirationUnixTime {
		t.Fail()
	}
}