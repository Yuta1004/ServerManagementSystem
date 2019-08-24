package db

import (
	"os"
	"testing"
)

func TestMySQLStruct(t *testing.T) {
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
	if m.GetConnectionInfo() != "USER@HOST/DB" {
		t.Fatal(m, m.GetConnectionInfo())
	}
}