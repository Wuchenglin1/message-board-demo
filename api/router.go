package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	//用户系统
	user := engine.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", Login)
		user.PUT("/changePassword", ChangePassword)
	}

	//密保系统
	mibaoSystem := engine.Group("/mibao")
	{
		mibaoSystem.POST("/add", auth, AddMB)
		mibaoSystem.GET("/check", auth, CheckMB)
	}

	//留言系统
	post := engine.Group("/post", auth)
	{
		post.GET("/", ViewAll)   //查看所有留言
		post.POST("/", Post)     //发送留言
		post.PUT("/", Modify)    //修改留言
		post.DELETE("/", Delete) //删除留言
		post.POST("/view", View) //按username或者receiveName来查找留言

	}

	//评论系统
	comment := engine.Group("/comment", auth)
	{
		comment.POST("/p", CommentPost)    //评论留言
		comment.POST("/c", CommentComment) //评论套娃
		comment.PUT("/", CommentModify)    //修改评论
		comment.POST("/", CommentView)     //查找评论(以及他的子评论)
		comment.DELETE("/", CommentDelete) //删除评论
	}
	_ = engine.Run()
}
