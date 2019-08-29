package db

import (
	"log"
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