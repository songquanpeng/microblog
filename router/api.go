package router

import (
	"github.com/gin-gonic/gin"
	"microblog/controller"
	"microblog/middleware"
)

func setApiRouter(router *gin.Engine) {
	router.Use(middleware.ApiAuth())
	router.GET("/status", controller.Status)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/api/post", controller.GetAllPosts)
	router.GET("/api/post/:id", controller.GetPost)
	router.POST("/api/post", middleware.AuthRequired(), controller.CreatePost)
	router.PUT("/api/post/:id", middleware.AuthRequired(), controller.UpdatePost)
	router.DELETE("/api/post/:id", middleware.AuthRequired(), controller.DeletePost)
}
