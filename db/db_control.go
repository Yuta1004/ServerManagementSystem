package db

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	/* => database/sqlで参照する */
)

// ControlDB : DB操作
func ControlDB(cName string, executable func(conn *sql.DB)(interface{})) (interface{}) {
	// DB接続
	connect, err := getDBConnect()
	if err != nil {
		return nil
	}
	defer connect.Close()

	result := executable(connect)
	return result
}

// ControlDBWithTransaction : トランザクション付きでDBを操作する
func ControlDBWithTransaction(txName string, executable func(tx *sql.Tx)) (result bool) {
	result = false

	// DB接続
	connect, err := getDBConnect()
	if err != nil {
		return
	}
	defer connect.Close()

	// トランザクション開始 & 後処理
	tx, err := connect.Begin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			log.Printf("[ERROR] Executed the rollback of DB on \"%s\" process.\n", txName)
		} else {
			tx.Commit()
			result = true
		}
	}()

	executable(tx)
	return
}


func getDBConnect() (*sql.DB, error) {
	mysqlInfo := AllocMySQLConnectInfo()
	connect, err := sql.Open("mysql", mysqlInfo.GetConnectionInfo())
	if err != nil {
		log.Println("[ERROR] Faild to get connect to db.")
		return nil, err
	}
	return connect, nil
}