package service

import (
	"message-board-demo/dao"
	"message-board-demo/model"
)

func CommentPost(comment model.Comments) (bool, error) {
	is, err := dao.CommentPost(comment)
	return is, err
}

func CommentComment(comment model.Comments) (bool, error) {
	is, err := dao.CommentComment(comment)
	return is, err
}

func CommentModify(comment model.Comments) error {
	err := dao.CommentModify(comment)
	return err
}

func CommentDelete(comments model.Comments) error {
	err := dao.CommentDelete(comments)
	return err
}
