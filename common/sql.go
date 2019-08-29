package common

import (
	"strings"
)


// MakeSQLINOperator : IN演算子生成
func MakeSQLINOperator(keywords []string) string {
	if len(keywords) == 0 {
		return "LIKE \"%\""
	}
	return "in (?" + strings.Repeat(",?", len(keywords)-1) + ")"
}
