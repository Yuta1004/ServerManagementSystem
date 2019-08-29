package db

import (
	"testing"
)

func TestControllCommandDB(t *testing.T) {
	// !!注意!!
	// 事前に以下のデータをcommandテーブルに挿入しておくこと
	// (1, "testuser", "test1", "command1", 0)

	// Update Test1
	request := make(map[string]interface{})
	request["command"] = "commandA"
	request["use_ok"] = 1
	result := UpdateCommandDataOfDB(1, "testuser", request)
	if !result {
		t.Fail()
	}

	// Get Test
	commandInfoList := *GetCommandDataFromDB("testuser")
	if len(commandInfoList) != 1 {
		t.Fail()
		return
	}
	if commandInfoList[0].Command != "commandA" {
		t.Fail()
	}
	if !commandInfoList[0].useOK {
		t.Fail()
	}

	// Update Test2
	request["command"] = "command1"
	request["use_ok"] = 0
	result = UpdateCommandDataOfDB(1, "testuser", request)
	if !result {
		t.Fail()
	}
}