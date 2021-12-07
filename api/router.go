package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	user := engine.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", Login)
		engine.PUT("/changePassword", ChangePassword)
	}

	mibaoSystem := engine.Group("/mibao")
	{
		mibaoSystem.POST("/add", auth, AddMB)
		mibaoSystem.GET("/check", auth, CheckMB)
	}

	post := engine.Group("/post", auth)
	{
		post.POST("/", Post)
		//post.PUT("/modify", Modify)
		//post.GET("/view", View)
	}

	comment := engine.Group("/comment", auth)
	{
		comment.POST("/")
		comment.PUT("/modify")
		comment.GET("/view")
	}
	_ = engine.Run()
}
