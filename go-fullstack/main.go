package main

import (
	"net/http"

	"go-fullstack/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		stories, err := models.LoadTopStories()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"stories": stories,
		})
	})
	r.Run("0.0.0.0:8000")
}
