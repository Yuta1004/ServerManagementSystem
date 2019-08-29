package resource

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// ErrorWrongOS : vmstatParseのエラー
type ErrorWrongOS string

func (e ErrorWrongOS) Error() string {
	return fmt.Sprintf("[ERROR] OS must be Linux. %s", string(e))
}

// vmstatParse : vmstat実行結果から値を読みとる
func vmstatParse(vmstatResult []byte) (map[string]int, error) {
	// OSチェック
	if runtime.GOOS != "linux" {
		return make(map[string]int), ErrorWrongOS(runtime.GOOS)
	}

	// 読み取り準備
	resultStr := string(vmstatResult)
	infoArray := strings.Split(strings.Split(resultStr, "\n")[2], " ")
	keys := []string{
		"ProcessR", "ProcessB", "MemSwap", "MemFree", "MemBuff", "MemCache", "SwapIn", "SwapOut",
		"IOIn", "IOOut", "SystemIn", "SystemCs", "CPUUser", "CPUSystem", "CPUIdol", "CPUWait", "CPUSt",
	}

	// 値読み取り
	var err error
	keyIdx := 0
	vmstatInfo := make(map[string]int)

	for _, elem := range infoArray {
		if elem == "" {
			continue
		}
		vmstatInfo[keys[keyIdx]], err = strconv.Atoi(elem)
		if err != nil {
			fmt.Println(err.Error())
		}
		keyIdx++
	}
	return vmstatInfo, nil
}
