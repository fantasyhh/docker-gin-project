package app

import (
	"github.com/fantasyhh/drizzle/pkg/e"
	"github.com/gin-gonic/gin"
)

// Gin struct include *gin.Context
type Gin struct {
	C *gin.Context
}

// FailResponse return fail response
func (g *Gin) FailResponse(httpCode int, flag string, err error) {
	g.C.JSON(httpCode, gin.H{
		"error":  flag,
		"msg":    e.GetMsg(flag),
		"detail": err.Error(),
	})
	return
}

// Response return success response
func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"data": data,
		"msg":  e.GetMsg(e.CommonSuccess),
	})
	return
}

// FailJSONResult struct for Swagger
type FailJSONResult struct {
	Error  string `json:"error"`
	Msg    string `json:"msg"`
	Detail string `json:"detail "`
}

// JSONResult struct for Swagger
type JSONResult struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
