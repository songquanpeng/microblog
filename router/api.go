package router

import (
	"github.com/gin-gonic/gin"
	"microblog/controller"
	"microblog/middleware"
)

func setApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.ApiAuth())
	{
		apiRouter.GET("/status", controller.Status)
		apiRouter.POST("/login", controller.Login)
		apiRouter.GET("/logout", controller.Logout)
		postRouter := apiRouter.Group("/post")
		{
			postRouter.GET("/", controller.GetAllPosts)
			postRouter.GET("/:id", controller.GetPost)
			postRouter.POST("/", middleware.AuthRequired(), controller.CreatePost)
			postRouter.PUT("/:id", middleware.AuthRequired(), controller.UpdatePost)
			postRouter.DELETE("/:id", middleware.AuthRequired(), controller.DeletePost)
		}
	}
}
