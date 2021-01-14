package base

import "github.com/gin-gonic/gin"

var commonOk = 1007
var commonNotFound = 1009
var commonForbidden = 1009

func ok(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": commonOk,
		"msg":  msg,
		"data": data,
	})
}

func notFound(c *gin.Context, msg string, data interface{}) {
	c.JSON(404, gin.H{
		"code": commonNotFound,
		"msg":  msg,
		"data": data,
	})
}

func forbidden(c *gin.Context, msg string, data interface{}) {
	c.JSON(403, gin.H{
		"code": commonForbidden,
		"msg":  msg,
		"data": data,
	})
}

func GetOk(c *gin.Context, msg string, data interface{}) {
	ok(c, msg, data)
}

func PostOk(c *gin.Context, msg string, data interface{}) {
	ok(c, msg, data)
}

func NotFound(c *gin.Context, msg string, data interface{}) {
	notFound(c, msg, data)
}

func Forbidden(c *gin.Context, msg string, data interface{}) {
	forbidden(c, msg, data)
}
