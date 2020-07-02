package app

import (
	"github.com/gin-gonic/gin"
)

// Gin struct include *gin.Context
type Gin struct {
	C *gin.Context
}

// FailResponse return fail response
func (g *Gin) FailResponse(httpCode int, msg error, detail error) {
	g.C.JSON(httpCode, gin.H{
		"msg":    msg.Error(),
		"detail": detail.Error(),
	})
	return
}

// Response return success response
func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"msg":  "ok",
		"data": data,
	})
	return
}

// FailJSONResult struct for Swagger
type FailJSONResult struct {
	Msg    string `json:"msg"`
	Detail string `json:"detail "`
}

// JSONResult struct for Swagger
type JSONResult struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
