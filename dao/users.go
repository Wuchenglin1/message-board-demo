package dao

import (
	"message-board-demo/model"
)

func SelectUsername(name string) (model.User, error) {
	user := model.User{}
	err := dB.QueryRow("select id,password from user where name = ?", name).Scan(&user.Id, &user.UserPassword)
	return user, err
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("insert into user(name,password) values(?,?)", user.Username, user.UserPassword)
	return err
}
