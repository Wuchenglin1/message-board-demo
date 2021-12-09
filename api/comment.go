package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
	"strconv"
	"time"
)

// CommentPost 评论留言需要两个key:pid(选择评论的留言id)和comment
func CommentPost(c *gin.Context) {
	postPid, _ := strconv.Atoi(c.PostForm("pid"))
	name, _ := c.Cookie("username")
	comment := model.Comments{
		PostPid:  postPid,
		UserName: name,
		Comment:  c.PostForm("comment"),
		Time:     time.Now().Add(time.Hour * 8),
	}
	if len(comment.Comment) >= 20 {
		tool.RespErrorWithDate(c, "评论长度过长！")
		return
	}
	is, err := service.CommentPost(comment)
	if is == false {
		tool.RespErrorWithDate(c, "该留言不存在！")
		return
	}
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfullWithDate(c, "评论成功！")
}

// CommentComment 评论套娃需要两个key:pid(选择评论的评论)和comment
func CommentComment(c *gin.Context) {
	username, _ := c.Cookie("username")
	pid, _ := strconv.Atoi(c.PostForm("pid"))

	comments := model.Comments{
		UserName: username,
		Pid:      pid,
		Comment:  c.PostForm("comment"),
		Time:     time.Now().Add(time.Hour * 8),
	}
	if len(comments.Comment) >= 20 {
		tool.RespErrorWithDate(c, "评论长度过长！")
		return
	}

	is, err := service.CommentComment(comments)
	if is == false {
		tool.RespErrorWithDate(c, "评论不存在！")
		return
	}
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfullWithDate(c, "评论成功！")
}

// CommentModify 评论修改需要两个key:id(需要修改评论的id)和comment(修改之后的内容)
func CommentModify(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comments := model.Comments{
		Comment: c.PostForm("comment"),
		Id:      id,
	}
	err := service.CommentModify(comments)
	if err != nil {
		tool.RespErrorWithDate(c, "该评论不存在！")
		return
	}
	tool.RespSuccessfullWithDate(c, "修改成功！")
}

// CommentDelete 删除评论需要1个key:id(删除评论的id)
//这里我采用的是覆盖评论，就将内容覆盖成“该内容已删除”,其他的东西就不动了
func CommentDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name, _ := c.Cookie("username")
	comments := model.Comments{
		Id:       id,
		UserName: name,
		Comment:  "已删除",
	}
	err := service.CommentDelete(comments)
	if err != nil {
		tool.RespErrorWithDate(c, "该评论不存在！")
		return
	}
	tool.RespSuccessfullWithDate(c, "删除成功！")

}

// CommentView 需要传入一个key:id(该评论的id)
func CommentView(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	comment := model.Comments{Id: id}

	//先通过id把该评论的详情给赋值出来
	err := service.CommentViewOne(&comment)
	tool.RespSuccessfullWithDate(c, comment)
	//检测该评论是否存在
	if err != nil {
		tool.RespErrorWithDate(c, "评论不存在！")
		return
	}
	//命名一个存comment的map
	commentMap := map[int]model.Comments{}
	//通过 子评论的pid=该评论的id 来查询他的所有子评论
	icommentMap, err1 := service.CommentView(comment, commentMap)
	if err != nil {
		fmt.Println(err1)
		tool.RespErrorWithDate(c, "啊嘞嘞，出现错误了诶~")
		return
	}
	for _, v := range icommentMap {
		tool.RespSuccessfullWithDate(c, v)
	}
}
