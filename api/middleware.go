package api

import (
	"github.com/gin-gonic/gin"
	"message-board-demo/tool"
)

func auth(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		tool.RespErrorWithDate(c, "请登录后进行操作")
		c.Abort()
	}

	c.Set("username", username)
	c.Next()
}
