package api

import (
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
)

func Post(c *gin.Context) {
	name, _ := c.Cookie("username")
	post := model.Post{
		Name:    name,
		Receive: c.PostForm("to"),
		Detail:  c.PostForm("post"),
	}
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
