package router

import "github.com/gin-gonic/gin"

func NewEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	return engine
}
