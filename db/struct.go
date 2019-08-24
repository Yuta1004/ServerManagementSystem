package db

// MySQLConnectInfo : Mysqlに接続するための情報を持つ構造体
type MySQLConnectInfo struct {
	User	string
	Host	string
	DBName	string
}