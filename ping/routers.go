package ping

import (
	"github.com/gin-gonic/gin"
)

func PingRegister(router *gin.RouterGroup) {
	router.GET("/ping", PingAnswer)
}

func PingAnswer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
