package db

// MySQLConnectInfo : Mysqlに接続するための情報を持つ構造体
type MySQLConnectInfo struct {
	User		string
	Host		string
	Port 		int
	DBName		string
	Password	string
}

// UserInfo : DBに格納されているユーザ情報をもつ構造体
type UserInfo struct {
	ID				string
	HashPassword 	string
}