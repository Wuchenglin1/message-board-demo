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

// Post 需要两个key:to(接收者)和detail(内容)
func Post(c *gin.Context) {
	name, _ := c.Cookie("username")
	post := model.Post{
		Name:    name,
		Receive: c.PostForm("to"),
		Detail:  c.PostForm("detail"),
		Time:    time.Now().Add(time.Hour * 8),
	}
	if len(post.Detail) >= 20 {
		tool.RespErrorWithDate(c, "留言长度过长！")
		return
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

// Modify 需要两个key:id(想要修改留言的id)和detail(新的修改的内容)
func Modify(c *gin.Context) {
	name, _ := c.Cookie("username")
	id, _ := strconv.Atoi(c.PostForm("id"))
	post := model.Post{
		Name:   name,
		Detail: c.PostForm("detail"),
		Id:     id,
	}
	if len(post.Detail) >= 20 {
		tool.RespErrorWithDate(c, "留言长度过长！")
		return
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

// Delete 需要一个key:id(要被删除的留言id)
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

// View 需要一个key:留言的id
func View(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	post := model.Post{Id: id}
	//查询一条留言的所有信息并赋值
	err := service.PostViewById(&post)
	if err != nil {
		tool.RespErrorWithDate(c, "该留言不存在！")
		return
	}
	comment := model.Comments{PostPid: id}
	commentMap := map[int]model.Comments{}
	commentSlice, _ := service.PostView2(comment, commentMap)
	fmt.Println(commentSlice)
	for _, v1 := range commentSlice {
		for _, v2 := range v1 {
			tool.RespSuccessfullWithDate(c, v2)
		}
	}
}

// ViewAll 直接查看所有留言(没有评论)
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
