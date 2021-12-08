package dao

import (
	"fmt"
	"message-board-demo/model"
)

func CommentPost(comments model.Comments) (bool, error) {
	var comment_num int
	err := dB.QueryRow("select comment_num from post where id = ?", comments.PostPid).Scan(&comment_num)
	if err != nil {
		return false, err
	}
	comment_num += 1
	_, err = dB.Exec("insert into comments(pid,postId,name,comment,time) values (0,?,?,?,?)", comments.PostPid, comments.UserName, comments.Comment, comments.Time)
	_, err = dB.Exec("update post	 set comment_num = ? where id = ?", comment_num, comments.PostPid)
	return true, err
}

func CommentComment(comments model.Comments) (bool, error) {
	var name string     //父亲的用户名
	var id int          //父亲的id
	var comment_num int //父亲的评论数
	err := dB.QueryRow("select name,id,comment_num from comments where id = ?", comments.Pid).Scan(&name, &id, &comment_num)
	_, err = dB.Exec("insert into comments(pid, name, comment, time) values (?,?,?,?) ", comments.Pid, comments.UserName, comments.Comment, comments.Time)
	comment_num += 1
	fmt.Println(comment_num)
	_, err = dB.Exec("update comments set comment_num = ? where id = ?", comment_num, id)
	fmt.Println("comment", comment_num)
	fmt.Println("error:", err)
	return true, err
}

func CommentModify(comment model.Comments) error {
	_, err := dB.Exec("update comments set comment = ? where id = ?", comment.Comment, comment.Id)
	return err
}

func CommentDelete(comment model.Comments) error {
	_, err := dB.Exec("delete from comments where id = ?and name = ? ", comment.Id, comment.UserName)
	return err
}
