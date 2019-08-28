package db

import (
	"log"
	"database/sql"
	"server-manage/common"
)


// GetSessionDataFromDB : セッションデータを取り出す
func GetSessionDataFromDB(request ...string) *[]SessionInfo {
	sessionInfoList := make([]SessionInfo, 0)
	executable :=
		func (conn *sql.DB) interface{} {
			// SQL実行
			sql := "select * from user where id " + common.MakeSQLINOperator(request)
			result, err := conn.Query(sql, common.ConvToInterfaceSlice(request)...)
			if err != nil {
				log.Println(err.Error())
				return &sessionInfoList
			}
			defer result.Close()

			// データ取り出し
			for result.Next() {
				var sessionInfo SessionInfo
				err := result.Scan(&sessionInfo.ID, &sessionInfo.Passphrase, &sessionInfo.ExpirationUnixTime)
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