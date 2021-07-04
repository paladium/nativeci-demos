package api

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/convert", Convert)
	return r
}
