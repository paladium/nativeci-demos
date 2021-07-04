package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func celciusToF(value float64) float64 {
	return (value * 1.8) + 32
}

type ConvertTemperature struct {
	Value float64 `json:"value"`
}

func Convert(c *gin.Context) {
	convert := new(ConvertTemperature)
	err := c.BindJSON(convert)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"value": celciusToF(convert.Value),
	})
}
