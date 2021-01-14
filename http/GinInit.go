package http

import (
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	engine := gin.New()
	RouteInit(engine)
	return engine
}
