package db

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	/* => database/sqlで参照する */
)

// InsertUserDataToDB : ユーザ情報をDBに追加する
func InsertUserDataToDB(userID, hashPassword string) (insertResult bool) {
	insertResult = false

	// DB接続
	connect, err := getDBConnect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer connect.Close()

	// トランザクション関連処理
	tx, err := connect.Begin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func() {
		result := recover()
		if result != nil {
			tx.Rollback()
			log.Println("[INFO] Executed the rollback of DB in \"InsertUserDataToDB\" function.")
			insertResult = false
		} else {
			tx.Commit()
		}
	}()

	// SQL実行
	_, err = tx.Exec("insert into user values(?, ?)", userID, hashPassword)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

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