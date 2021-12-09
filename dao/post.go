package dao

import (
	"database/sql"
	"fmt"
	"message-board-demo/model"
)

func Post(post model.Post) error {
	_, err := dB.Exec("insert into post(name,receive,detail,time) values(?,?,?,?) ", &post.Name, &post.Receive, &post.Detail, &post.Time)
	return err
}

func PostDelete(post model.Post) error {
	fmt.Println(post)
	_, err := dB.Exec("delete from post where id = ?", post.Id)
	return err
}

func PostModify(post model.Post) error {
	fmt.Println(post)
	_, err := dB.Exec("update post set detail=? where name = ? and id = ?", post.Detail, post.Name, post.Id)
	return err
}

func PostViewById(post *model.Post) error {
	err := dB.QueryRow("select name,receive,detail,comment_num,time from post where id = ?", post.Id).Scan(&post.Name, &post.Receive, &post.Detail, &post.Comment_num, &post.Time)
	return err
}

//func PostViewComments(comment model.Comments, commentMap map[int]model.Comments) (map[int]model.Comments, error) {
//	rows, err := dB.Query("select pId,id,name,comment,comment_num,time from comments where pId = ?", comment.Id)
//	fmt.Println("err3", err)
//	if err != nil {
//		return commentMap, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		sc := model.Comments{}
//		err = rows.Scan(&sc.Pid, &sc.Id, &sc.UserName, &sc.Comment, &sc.Comment_Num, &sc.Time)
//		fmt.Println("err", i, err)
//		commentMap[i] = sc
//		i++
//		if err != nil {
//			return commentMap, err
//		}
//		commentMap, err = CommentView(sc, commentMap)
//		if err != nil {
//			return commentMap, err
//		}
//	}
//	return commentMap, nil
//}

func PostView(post model.Post, num int) (map[int]model.Post, bool, error) {
	//我选择用一个map来存储所有查询到的留言,用数字来判断查询的类型
	var err error
	rows := &sql.Rows{}
	postMap := model.PostMap
	ipost := model.Post{}
	switch num {
	case 0:
		rows, err = dB.Query("select id,name,receive,detail,comment_num,time from post where id > ?", 0)
	case 1:
		rows, err = dB.Query("select id,name,receive,detail,comment_num,time from post where receive = ?", post.Receive)
	case 2:
		rows, err = dB.Query("select id,name,receive,detail,comment_num,time from post where name = ?", post.Name)
	}
	if err != nil {
		fmt.Println(err)
		return postMap, false, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&ipost.Id, &ipost.Name, &ipost.Receive, &ipost.Detail, &ipost.Comment_num, &ipost.Time)
		if err != nil {
			fmt.Println(err)
			return postMap, true, err
		}
		postMap[ipost.Id] = ipost
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return postMap, false, err
	}
	return postMap, false, nil
}
