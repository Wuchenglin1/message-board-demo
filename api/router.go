package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register)
	engine.POST("/login", Login)
	engine.POST("/changePassword", ChangePassword)
	engine.PUT("/mibaoAdd", auth, AddMB)
	engine.GET("/mibaoCheck", auth, CheckMB)

	_ = engine.Run()
}
