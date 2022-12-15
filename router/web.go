package router

import (
	"github.com/gin-gonic/gin"
	"microblog/controller"
)

func setWebRouter(router *gin.Engine) {
	router.GET("/public/static/:file", controller.GetStaticFile)
	router.GET("/public/lib/:file", controller.GetLibFile)
	router.GET("/", controller.GetIndex)
}
