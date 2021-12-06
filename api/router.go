package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register)
	engine.POST("/login", Login)

	engine.Run()
}
