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
		engine.PUT("/changePassword", ChangePassword)
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
		post.GET("/", ViewAll)
		post.POST("/", Post)
		post.PUT("/", Modify)
		post.DELETE("/", Delete)
		post.POST("/view", View)

	}

	//评论系统
	comment := engine.Group("/comment", auth)
	{
		comment.POST("/")
		comment.PUT("/")
		comment.GET("/")
		comment.DELETE("/")
	}
	_ = engine.Run()
}
