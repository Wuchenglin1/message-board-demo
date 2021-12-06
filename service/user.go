package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

// IsRepeatUsername 先检测error,
func IsRepeatUsername(username string) error {
	_, err := dao.SelectUsername(username)
	if err != nil {
		return err
	}
	return nil
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

// IsUserCorrect www
func IsUserCorrect(user model.User) bool {
	iUser, _ := dao.SelectUsername(user.Username)
	if user.UserPassword == iUser.UserPassword {
		return true
	}
	return false
}
