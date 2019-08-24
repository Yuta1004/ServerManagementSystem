package db

import (
	"os"
	"testing"
)

func TestMySQLStruct(t *testing.T) {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DB")

	os.Setenv("MYSQL_USER", "USER")
	os.Setenv("MYSQL_HOST", "HOST")
	os.Setenv("MYSQL_DB", "DB")

	m := AllocMySQLConnectInfo()

	if m.User != "USER" {
		t.Fatal(m)
	}
	if m.Host != "HOST" {
		t.Fatal(m)
	}
	if m.DBName != "DB" {
		t.Fatal(m)
	}
	if m.GetConnectionInfo() != "USER:@/DB" {
		t.Fatal(m, m.GetConnectionInfo())
	}

	os.Setenv("MYSQL_USER", user)
	os.Setenv("MYSQL_HOST", host)
	os.Setenv("MYSQL_DB", db)
}