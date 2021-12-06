package main

import (
	"message-board-demo/api"
	"message-board-demo/dao"
)

func main() {
	dao.InitMySql()
	api.InitEngine()

}
