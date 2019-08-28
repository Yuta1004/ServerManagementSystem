package db

import (
	"log"
	"database/sql"
	"server-manage/common"
)

// InsertSessionDataToDB : セッションデータをDBに記録、既にidが存在していたら更新処理
func InsertSessionDataToDB(userID, passphrase string, expirationUnixTime int) bool {
	executable :=
		func (tx *sql.Tx) {
			// SQL実行
			sql := "insert into session values(?, ?, ?)" +
					" on duplicate key update passphrase=?, expiration_unix_time=?"
			_, err := tx.Exec(sql, userID, passphrase, expirationUnixTime, passphrase, expirationUnixTime)
			if err != nil {
				log.Println(err.Error())
				panic("An error happened in \"InsertSessionData\" function.")
			}
		}

	result := ControlDBWithTransaction("InsertSessionData", executable)
	return result
}

// GetSessionDataFromDB : セッションデータを取り出す
func GetSessionDataFromDB(request ...string) *[]SessionInfo {
	sessionInfoList := make([]SessionInfo, 0)
	executable :=
		func (conn *sql.DB) interface{} {
			// SQL実行
			sql := "select * from session where id " + common.MakeSQLINOperator(request)
			result, err := conn.Query(sql, common.ConvToInterfaceSlice(request)...)
			if err != nil {
				log.Println(err.Error())
				return &sessionInfoList
			}
			defer result.Close()

			// データ取り出し
			for result.Next() {
				var sessionInfo SessionInfo
				err := result.Scan(
					&sessionInfo.ID,
					&sessionInfo.Passphrase,
					&sessionInfo.ExpirationUnixTime,
				)
				if err != nil {
					log.Println(err.Error())
					return &sessionInfoList
				}
				sessionInfoList = append(sessionInfoList, sessionInfo)
			}
			return &sessionInfoList
		}

	result := ControlDB("GetSessionData", executable)
	return result.(*[]SessionInfo)
}