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
	//将被评论的评论数+1
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
	_, err := dB.Exec("update comments set comment=? where id = ?and name = ? ", comment.Comment, comment.Id, comment.UserName)
	return err
}

func CommentViewOne(comment *model.Comments) error {
	err := dB.QueryRow("select pId,id,name,comment,comment_num,time from comments where id = ?", comment.Id).Scan(&comment.Pid, &comment.Id, &comment.UserName, &comment.Comment, &comment.Comment_Num, &comment.Time)
	return err
}

//为了方便存储，我命名了个全局变量i
var i = 1

func CommentView(comment model.Comments, commentMap map[int]model.Comments) (map[int]model.Comments, error) {
	rows, err := dB.Query("select pId,id,name,comment,comment_num,time from comments where pId = ?", comment.Id)

	if err != nil {
		return commentMap, err
	}
	defer rows.Close()
	for rows.Next() {

		sc := model.Comments{} //新建一个子评论的模板

		//将子评论的信息赋值给sc(sonComment)
		err = rows.Scan(&sc.Pid, &sc.Id, &sc.UserName, &sc.Comment, &sc.Comment_Num, &sc.Time)

		//因为父评论已经被打印，这里只需要存子评论就行
		//为了好存储，在api里遍历，我就通过自增全局变量i来按1234...存储子评论
		commentMap[i] = sc
		i++

		if err != nil {
			return commentMap, err
		}
		//采用递归的方式遍历评论树
		//这里要将返回的子评论map赋值给评论map,最后的commentMap里面才会存有子评论
		commentMap, err = CommentView(sc, commentMap)
		if err != nil {
			return commentMap, err
		}
	}
	return commentMap, nil
}
