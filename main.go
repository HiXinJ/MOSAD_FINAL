package main

import (
	"github.com/hixinj/MOSAD_FINAL/views"

	// "code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users", views.UserRegister)
	router.POST("/users/login", views.UserLogin)
	router.GET("/words", views.GetWords)
	router.GET("/words/new", views.GetNewWords)
	router.POST("/words/learnedword", views.AddLearnedWrod)
	router.GET("/words/reviews", views.GetReviews)
	router.GET("/words/translation", views.GetTranslation)
	router.POST("/users/daka", views.DaKa)

	router.Run(":8081")
}
