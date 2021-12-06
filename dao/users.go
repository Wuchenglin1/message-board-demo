package dao

import (
	"fmt"
	"message-board-demo/model"
)

// SelectUsername 返回一个user结构体和一个error
func SelectUsername(name string) (model.User, error) {
	user := model.User{}
	err := dB.QueryRow("select id,name,password from user where name = ?", name).Scan(&user.Id, &user.Username, &user.UserPassword)
	return user, err
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("insert into user(name,password) values(?,?)", user.Username, user.UserPassword)
	return err
}

func AddMB(mb model.MiBao) error {
	_, err := dB.Exec("insert into mibao values(?,?,?,?,?)", mb.Id, mb.Mb1, mb.Mb1pwd, mb.Mb2, mb.Mb2pwd)
	return err
}

func SelectMB(mb model.MiBao) (model.MiBao, error) {
	imb := model.MiBao{}
	err := dB.QueryRow("select id,mb1,mb1pwd,mb2,mb2pwd from mibao where id = ?", mb.Id).Scan(&imb.Id, &imb.Mb1, &imb.Mb1pwd, &imb.Mb2, &imb.Mb2pwd)
	return imb, err
}

func ChangePassword(id string, pwd string) error {
	fmt.Println(id)
	_, err := dB.Exec("update user set password = ? where id = ?", pwd, id)
	return err
}
