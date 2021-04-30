package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": []string{"Computer", "Iphone"},
		})
	})
	r.Run("0.0.0.0:8000")
}
