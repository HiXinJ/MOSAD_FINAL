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
	router.POST("/users/{user_name}/learnedword", views.AddLearnedWrod)
	router.GET("/words", views.GetWords)
	router.GET("/words/new", views.GetNewWords)
	router.GET("/words/translation", views.GetTranslation)
	router.POST("/users/{user_name}/daka", views.DaKa)

	router.Run(":8081")
}
