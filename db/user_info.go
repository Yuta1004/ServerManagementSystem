package db

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	/* => database/sqlで参照する */
)

// GetUserDataFromDB : DBからユーザ情報を引っ張ってくる
func GetUserDataFromDB() *[]UserInfo {
	userInfoList := make([]UserInfo, 0)

	// DB接続
	connect, err := getDBConnect()
	if err != nil {
		log.Println(err.Error())
		return &userInfoList
	}
	defer connect.Close()

	// SQL実行
	rows, err := connect.Query("select * from user")
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

// FaildToGetConnectionError : getDBConnect() が返すエラー
type FaildToGetConnectionError string

func (f FaildToGetConnectionError) Error() string {
	return fmt.Sprintf("[ERROR] Faild to get connect to db. (%s)", string(f))
}

func getDBConnect() (*sql.DB, error) {
	// DB接続
	mysqlInfo := AllocMySQLConnectInfo()
	connect, err := sql.Open("mysql", mysqlInfo.GetConnectionInfo())
	if err != nil {
		return nil, FaildToGetConnectionError(mysqlInfo.GetConnectionInfo())
	}
	return connect, nil
}