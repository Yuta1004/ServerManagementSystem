package db

import (
	"log"
	"strings"
	"database/sql"
)

// InsertCommandDataToDB : コマンド情報をDBに挿入する
func InsertCommandDataToDB(userID, name, command string) bool {
	executable :=
		func(tx *sql.Tx) {
			// SQL実行
			sql := "insert into command (user_id, name, command, use_ok) values(?, ?, ?, 0)"
			_, err := tx.Exec(sql, userID, name, command)
			if err != nil {
				log.Println(err.Error())
				panic("An error happened in \"InsertCommandData\" function.")
			}
		}

	result := ControlDBWithTransaction("InsertCommandData", executable)
	return result
}

// UpdateCommandDataDB : DBに格納されているコマンド情報を更新する
func UpdateCommandDataOfDB(id int, userID string, request map[string]interface{}) bool {
	// 更新するキー, 値を取り出す
	reqKeySqls := make([]string, 0)
	reqValues := make([]interface{}, 0)
	for key, value := range request {
		reqKeySqls = append(reqKeySqls, key+"=?")
		reqValues = append(reqValues, value)
	}

	// SQL組み立て
	sqlBody := "update command set " + strings.Join(reqKeySqls, ", ") + " where id=? and user_id=?"
	reqValues = append(reqValues, id)
	reqValues = append(reqValues, userID)

	executable :=
		func (tx *sql.Tx) {
			// SQL実行
			_, err := tx.Exec(sqlBody, reqValues...)
			if err != nil {
				log.Println(err.Error())
				panic("An error happened in \"UpdateCommandData\" function.")
			}
		}

	result := ControlDBWithTransaction("UpdateCommandData", executable)
	return result
}

// GetCommandDataFromDB : コマンド情報をDBから引っ張ってくる
func GetCommandDataFromDB(userID string) *[]CommandInfo {
	commandInfoList := make([]CommandInfo, 0)

	executable :=
	 	func (conn *sql.DB) interface{} {
			// SQL実行
			sql := "select * from command where user_id=?"
			fetchResult, err := conn.Query(sql, userID)
			if err != nil {
				log.Println(err.Error())
				return &commandInfoList
			}

			// データ取り出し
			var tmp, useOk int
			for fetchResult.Next() {
				var commandInfo CommandInfo
				fetchResult.Scan(&tmp, &commandInfo.UserID, &commandInfo.Name,
								 &commandInfo.Command, &useOk)
				commandInfo.useOK = useOk == 1
				commandInfoList = append(commandInfoList, commandInfo)
			}
			return &commandInfoList
		}

	result := ControlDB("GetCommandData", executable)
	return result.(*[]CommandInfo)
}