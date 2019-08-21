package crontabmanage

import (
	"server-manage/common"
	"strconv"
	"strings"
)

func crontabResultParse(result []byte) []*Crontab {
	resultStr := string(result)
	resultArray := strings.Split(resultStr, "\n")
	crontabInfoArray := make([]*Crontab, 0)

	for _, line := range resultArray {
		crontabInfo := strings.Split(line, " ")
		if len(crontabInfo) < 2 {
			continue;
		}
		crontab := AllocCrontabStruct()
		crontab.Minute = crontabDateParse(crontabInfo[0], 0, 59)
		crontab.Hour = crontabDateParse(crontabInfo[1], 0, 23)
		crontab.DayOfMonth = crontabDateParse(crontabInfo[2], 1, 31)
		crontab.Month = crontabDateParse(crontabInfo[3], 1, 12)
		crontab.Day = crontabDateParse(crontabInfo[4], 0, 6)
		crontab.Command = crontabInfo[5:]
		crontabInfoArray = append(crontabInfoArray, crontab)
	}

	return crontabInfoArray
}

func crontabDateParse(date string, lower, upper int) CrontabDateNum {
	dateArray := []int{}
	wildcardFlag := false

	for _, elem := range strings.Split(date, ",") {
		// ワイルドカード
		if strings.Contains(elem, "*") {
			wildcardFlag = true
			interval := 1
			if strings.Contains(elem, "/") {
				interval, _ = strconv.Atoi(strings.Split(elem, "/")[1])
			}
			dateArray = append(
				dateArray,
				makeArrayWithInterval(lower, upper, interval)...,
			)
		}

		// その他
		convertedI, _ := strconv.Atoi(elem)
		dateArray = append(dateArray, convertedI)
	}

	return CrontabDateNum{wildcardFlag, common.DeduplicationArrayInt(dateArray), date}
}

func makeArrayWithInterval(lower, upper, interval int) []int {
	retArray := make([]int, 0)
	for num := lower; num <= upper; num += interval {
		retArray = append(retArray, num)
	}
	return retArray
}
