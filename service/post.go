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

// PostViewById PostCheck 通过留言的id查询一条留言信息
func PostViewById(post *model.Post) error {
	err := dao.PostViewById(post)
	return err
}

func PostView2(comment model.Comments, commentMap map[int]model.Comments) ([]map[int]model.Comments, error) {
	slice, err := dao.CommentViewByPostId(comment, commentMap)
	return slice, err
}

func PostView(post model.Post, num int) (map[int]model.Post, bool, error) {
	postMap, is, err := dao.PostView(post, num)
	return postMap, is, err
}
