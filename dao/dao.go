package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitMySql() {
	db, err := sql.Open("mysql", "root:root@/test?parseTime=true")
	if err != nil {
		panic(err)
	}
	dB = db
}
