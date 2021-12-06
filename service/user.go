package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

// IsRepeatUsername 如果账户不存在则返回一个err，若存在则返回一个nil
func IsRepeatUsername(username string) (model.User, error) {
	user, err := dao.SelectUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

// IsUserCorrect 检查账号密码是否正确，如果正确返回一个true，若错误则返回一个false
func IsUserCorrect(user *model.User) bool {
	iUser, _ := dao.SelectUsername(user.Username)
	if user.UserPassword == iUser.UserPassword {
		user.Id = iUser.Id
		return true
	}
	return false
}

func AddMB(mb model.MiBao) error {
	err := dao.AddMB(mb)
	return err
}

func CheckMB(mb model.MiBao) (model.MiBao, error) {
	imb, err := dao.SelectMB(mb)
	return imb, err
}

func ChangePassword(mb model.MiBao, pwd string) (bool, error) {
	imb, err := dao.SelectMB(mb)
	if err != nil {
		return false, err
	}
	if mb.Mb1pwd == imb.Mb1pwd || mb.Mb2pwd == imb.Mb2pwd {
		err = dao.ChangePassword(mb.Id, pwd)
		return true, nil
	} else {
		return false, nil
	}
}
