package db

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	/* => database/sqlで参照する */
	"server-manage/common"
)

// InsertUserDataToDB : ユーザ情報をDBに追加する
func InsertUserDataToDB(userID, hashPassword string) bool {
	executable :=
		func (tx *sql.Tx) {
			// SQL実行
			_, err := tx.Exec("insert into user values(?, ?)", userID, hashPassword)
			if err != nil {
				log.Println(err.Error())
				panic("An error happened in \"InsertUser\" function.")
			}
			return
		}

	result := ControlDBWithTransaction("InsertUser", executable)
	return result
}

// GetUserDataFromDB : DBからユーザ情報を引っ張ってくる
func GetUserDataFromDB(request ...string) *[]UserInfo {
	userInfoList := make([]UserInfo, 0)
	executable :=
		func (conn *sql.DB) (interface{}) {
			// SQL実行
			sql := "select * from user where id " + common.MakeSQLINOperator(request)
			rows, err := conn.Query(sql, common.ConvToInterfaceSlice(request)...)
			if err != nil {
				log.Println(err.Error())
				return &userInfoList
			}
			defer rows.Close()

			// データ取り出し
			for rows.Next() {
				var userInfo UserInfo
				err := rows.Scan(&userInfo.ID, &userInfo.HashPassword)
				if err != nil {
					log.Println(err.Error())
				}
				userInfoList = append(userInfoList, userInfo)
			}
			return &userInfoList
		}

	result := ControlDB("GetUserData", executable)
	return result.(*[]UserInfo)
}
