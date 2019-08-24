package db

import (
	"os"
	"fmt"
	"strconv"
)

// AllocMySQLConnectInfo : MyS...Info構造体を確保して返す
func AllocMySQLConnectInfo() *MySQLConnectInfo {
	m := MySQLConnectInfo{}
	m.Update()
	return &m
}

// Update : 情報更新
func (m *MySQLConnectInfo) Update() {
	m.User = os.Getenv("MYSQL_USER")
	m.Host = os.Getenv("MYSQL_HOST")
	m.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	m.DBName = os.Getenv("MYSQL_DB")
	m.Password = os.Getenv("MYSQL_PASSWORD")
}

// GetConnectionInfo : MySQLに接続する際に必要な情報を返す
func (m *MySQLConnectInfo) GetConnectionInfo() string {
	var addr string
	if m.Host == "localhost" {
		addr = ""
	} else {
		addr = fmt.Sprintf("tcp(%s:%d)", m.Host, m.Port)
	}
	return fmt.Sprintf("%s:%s@%s/%s", m.User, m.Password, addr, m.DBName)
}