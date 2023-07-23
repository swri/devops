package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"status":  "success",
			"message": "Hello, This is Go Lang Application",
		})

	})

	g.Run(":5000")
}
