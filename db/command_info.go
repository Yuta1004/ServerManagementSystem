package db

import (
	"log"
	"strings"
	"database/sql"
)

func InsertCommandDataToDB(userID, name, command string) bool {
	executable :=
		func(tx *sql.Tx) {
			// SQL実行
			sql := "insert into command values(?, ?, ?, 0)"
			_, err := tx.Exec(sql, userID, name, command)
			if err != nil {
				log.Println(err.Error())
				panic("An error happened in \"InsertCommandData\" function.")
			}
		}

	result := ControlDBWithTransaction("InsertCommandData", executable)
	return result
}

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