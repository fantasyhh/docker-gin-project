package api

import (
	"github.com/gin-gonic/gin"
)

// PING godoc
// @Summary PING
// @Description This is a test function
// @Tags PING
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PING(c *gin.Context) {
	//log.Println("this is a log test")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
