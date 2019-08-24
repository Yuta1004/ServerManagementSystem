package db

import (
	"os"
	"testing"
)

func TestMySQLStruct(t *testing.T) {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	db := os.Getenv("MYSQL_DB")
	password := os.Getenv("MYSQL_PASSWORD")

	os.Setenv("MYSQL_USER", "USER")
	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_PORT", "1111")
	os.Setenv("MYSQL_DB", "DB")
	os.Setenv("MYSQL_PASSWORD", "PASSWORD")

	m := AllocMySQLConnectInfo()

	if m.User != "USER" {
		t.Fatal(m)
	}
	if m.Host != "localhost" {
		t.Fatal(m)
	}
	if m.Port != 1111 {
		t.Fatal(m)
	}
	if m.DBName != "DB" {
		t.Fatal(m)
	}
	if m.Password != "PASSWORD" {
		t.Fatal(m)
	}
	if m.GetConnectionInfo() != "USER:PASSWORD@/DB" {
		t.Fatal(m, m.GetConnectionInfo())
	}

	os.Setenv("MYSQL_USER", user)
	os.Setenv("MYSQL_HOST", host)
	os.Setenv("MYSQL_PORT", port)
	os.Setenv("MYSQL_DB", db)
	os.Setenv("MYSQL_PASSWORD", password)
}