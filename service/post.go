package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

func Post(post model.Post) error {
	err := dao.Post(post)
	return err
}

func PostDelete(post model.Post) error {
	err := dao.PostDelete(post)
	return err
}

func PostModify(post model.Post) error {
	err := dao.PostModify(post)
	return err
}

func PostView(post model.Post, num int) (map[int]model.Post, bool, error) {
	postMap, is, err := dao.PostView(post, num)
	return postMap, is, err
}
