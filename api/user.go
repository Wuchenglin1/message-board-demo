package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board-demo/model"
	"message-board-demo/service"
	"message-board-demo/tool"
)

//需要传入两个key:username和password
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) <= 6 || len(password) <= 6 {
		tool.RespErrorWithDate(c, "非法输入：账号和密码长度需要大于六位！")
		return
	}
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

// Login 需要传入两个key:username和password
func Login(c *gin.Context) {
	user := model.User{
		Username:     c.PostForm("username"),
		UserPassword: c.PostForm("password"),
	}
	_, err := service.IsRepeatUsername(user.Username)
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

	is := service.IsUserCorrect(&user)

	if is == false {
		tool.RespErrorWithDate(c, "您的密码错误!")
		return
	}
	fmt.Println(user)
	c.SetCookie("username", user.Username, 3600, "/", "", false, true)
	c.SetCookie("id", user.Id, 3600, "/", "", false, true)
	tool.RespSuccessfullWithDate(c, "登录成功！")

}

// AddMB 需要传入4个key：mb1,mb1pwd和mb2,mb2pwd
func AddMB(c *gin.Context) {
	id, _ := c.Cookie("id")
	mb := model.MiBao{
		Id:     id,
		Mb1:    c.PostForm("mb1"),
		Mb1pwd: c.PostForm("mb1pwd"),
		Mb2:    c.PostForm("mb2"),
		Mb2pwd: c.PostForm("mb2pwd"),
	}
	if mb.Mb1 == "" || mb.Mb2 == "" || mb.Mb1pwd == "" || mb.Mb2pwd == "" {
		tool.RespErrorWithDate(c, "值不能为空！")
		return
	}
	err := service.AddMB(mb)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(c, "您已经存在密保啦！")
		return
	}
	tool.RespSuccessfullWithDate(c, "添加成功！")
}

func CheckMB(c *gin.Context) {
	id, _ := c.Cookie("id")
	mb := model.MiBao{Id: id}
	imb, err := service.CheckMB(mb)
	if err != nil {
		tool.RespErrorWithDate(c, "您还没有密保哟~")
		return
	}
	tool.RespSuccessfullWithDate(c, imb)
}

// ChangePassword 需要传入3个key:username(想要修改名字的username)和mb1pwd,mb2pwd(两个密保的密码)
func ChangePassword(c *gin.Context) {
	user := model.User{Username: c.PostForm("username")}
	iuser, err := service.IsRepeatUsername(user.Username)
	if err != nil {
		tool.RespErrorWithDate(c, "账号不存在！")
	}
	pwd := c.PostForm("password")
	mb := model.MiBao{
		Id:     iuser.Id,
		Mb1pwd: c.PostForm("mb1pwd"),
		Mb2pwd: c.PostForm("mb2pwd"),
	}
	is, err1 := service.ChangePassword(mb, pwd)
	if err1 != nil {
		tool.RespInternalError(c)
		return
	}
	if is != true {
		tool.RespErrorWithDate(c, "密保答案错误！")
		return
	} else {
		tool.RespSuccessfullWithDate(c, "修改成功！")
	}
}
