package dao

import "message-board-demo/model"

func Post(post model.Post) error {
	_, err := dB.Exec("insert into post values(?,?,?) ", &post.Name, &post.Receive, &post.Detail)
	return err
}
