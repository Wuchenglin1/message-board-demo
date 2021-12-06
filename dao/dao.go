package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitMySql() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	dB = db
}
