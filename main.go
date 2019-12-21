package main

import (
	"github.com/hixinj/MOSAD_FINAL/views"

	// "code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/words", views.GetWords)
	r.Run(":8081")
}
