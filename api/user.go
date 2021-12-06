package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username:     username,
		UserPassword: password,
	}

	err := service.Register(user)
	if err != nil {
		fmt.Println("ERROR1", err)
		tool.RespErrorWithDate(c, "注册失败，账号已存在！")
		return
	}
	tool.RespSuccessfullWithDate(c, "注册成功！")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	userPassword := c.PostForm("password")
	user := model.User{
		Username:     username,
		UserPassword: userPassword,
	}

	err := service.IsRepeatUsername(user.Username)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("ERROR:", err)
			tool.RespInternalError(c)
			return
		} else {
			tool.RespErrorWithDate(c, "账号不存在！")
			return
		}
	}

	is := service.IsUserCorrect(user)

	if !is {
		tool.RespErrorWithDate(c, "您的密码错误!")
		return
	}

	c.SetCookie("username", username, 3600, "/", "", false, true)
	tool.RespSuccessful(c)
}
