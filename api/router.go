package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	user := engine.Group("/user")
	{
		user.POST("/register", register)
		user.PUT("/login", Login)
		engine.PUT("/changePassword", ChangePassword)
	}

	mibaoSystem := engine.Group("/mibao")
	{
		mibaoSystem.POST("/add", auth, AddMB)
		mibaoSystem.GET("/check", auth, CheckMB)
	}

	_ = engine.Run()
}
