package model

import "time"

type Post struct {
	Id          int
	Name        string
	Receive     string
	Detail      string
	Comment_num int
	Time        time.Time
}

var PostMap = make(map[int]Post)
