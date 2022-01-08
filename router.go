package main

import (
	"github.com/gin-gonic/gin"
)

func SetIndexRouter(router *gin.Engine) {
	router.GET("/public/static/:file", GetStaticFile)
	router.GET("/public/lib/:file", GetLibFile)
	router.GET("/", GetIndex)
}

func SetApiRouter(router *gin.Engine) {
	router.GET("/api/nonsense/:id", GetNonsenseById)
	router.GET("/api/nonsense", GetNonsense)
	router.POST("/api/nonsense", PostNonsense)
	router.DELETE("/api/nonsense", DeleteNonsense)
}
