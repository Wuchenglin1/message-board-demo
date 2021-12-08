package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"strconv"
	"time"
)

// Post 没有对字数的检测
func Post(c *gin.Context) {
	name, _ := c.Cookie("username")
	post := model.Post{
		Name:    name,
		Receive: c.PostForm("to"),
		Detail:  c.PostForm("post"),
		Time:    time.Now().Add(time.Hour * 8),
	}
	fmt.Println(post.Time)
	_, err := service.IsRepeatUsername(post.Receive)
	if err != nil {
		tool.RespErrorWithDate(c, "您想发送的人不存在！")
		return
	}

	err = service.Post(post)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfullWithDate(c, "发送成功！")
}

func Modify(c *gin.Context) {
	name, _ := c.Cookie("username")
	id, _ := strconv.Atoi(c.PostForm("id"))
	post := model.Post{
		Name:   name,
		Detail: c.PostForm("detail"),
		Id:     id,
	}
	err := service.PostModify(post)
	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithDate(c, "您想修改的内容不存在！")
			return
		}
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfullWithDate(c, "修改成功！")
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	post := model.Post{
		Id: id,
	}
	err := service.PostDelete(post)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfullWithDate(c, "删除成功！")
}

func View(c *gin.Context) {
	var num int
	post := model.Post{
		Name:    c.PostForm("username"),
		Receive: c.PostForm("receiveName"),
	}
	if post.Name == "" && post.Receive == "" {
		tool.RespErrorWithDate(c, "请输入正确的名字username或receiveName")
		return
	}
	if post.Name == "" {
		num = 1
	} else {
		num = 2
	}
	postMap, is, err := service.PostView(post, num)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	if is == true {
		tool.RespErrorWithDate(c, "该用户还没有（收到）留言！")
		return
	}
	for _, v := range postMap {
		tool.RespSuccessfullWithDate(c, v)
	}
}

func ViewAll(c *gin.Context) {
	post := model.Post{}
	postMap, _, err := service.PostView(post, 0)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	for _, v := range postMap {
		tool.RespSuccessfullWithDate(c, v)
	}
}
