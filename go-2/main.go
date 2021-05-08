package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/tamerlan", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world2",
		})
	})
	r.Run("0.0.0.0:8000")
}
