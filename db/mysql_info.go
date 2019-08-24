package db

import (
	"os"
	"fmt"
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
	m.DBName = os.Getenv("MYSQL_DB")
}

// GetConnectionInfo : MySQLに接続する際に必要な情報を返す
func (m *MySQLConnectInfo) GetConnectionInfo() string {
	return fmt.Sprintf("%s:@/%s", m.User, m.DBName)
}