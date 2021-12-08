package api

import (
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
func CommentDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name, _ := c.Cookie("username")
	comments := model.Comments{
		Id:       id,
		UserName: name,
	}
	err := service.CommentDelete(comments)
	if err != nil {
		tool.RespErrorWithDate(c, "该评论不存在！")
		return
	}
	tool.RespSuccessfullWithDate(c, "删除成功！")

}
