package crontabmanage

import (
	"fmt"
	"strings"
	"os/exec"
)

// AllocCrontabStruct : Crontab構造体を新しく作って返す
func AllocCrontabStruct() *Crontab {
	c := Crontab{}
	return &c
}

// GetLatestCrontabData : Crontab情報を更新する
func GetLatestCrontabData() []*Crontab {
	result, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	return crontabResultParse(result)
}

// GetCommandStr : Crontab構造体のCommandを文字列にして返す
func (c *Crontab) GetCommandStr() string {
	return strings.Join(c.Command, " ")
}
