package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

func Post(post model.Post) error {
	err := dao.Post(post)
	return err
}
