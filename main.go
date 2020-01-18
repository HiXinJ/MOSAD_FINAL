package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hixinj/MOSAD_FINAL_Group05/views"
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
	router.GET("/users/head", views.GetHead)
	router.POST("/users/head", views.PostHead)
	router.Run(":8081")
}
