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

// SessionInfo : DBに格納されているセッション情報をもつ構造体
type SessionInfo struct {
	ID					string
	Passphrase			string
	ExpirationUnixTime	int
}

// CommandInfo : DBに格納されているコマンド情報をもつ構造体
type CommandInfo struct {
	ID		int
	UserID	string
	Name	string
	Command	string
	useOK	bool
}