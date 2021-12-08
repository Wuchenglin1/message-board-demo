package model

import "time"

type Comments struct {
	UserName    string
	Id          int
	PostPid     int
	Pid         int
	Comment     string
	Comment_Num int
	Time        time.Time
}
