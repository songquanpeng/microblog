package router

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	setApiRouter(router)
	setWebRouter(router)
}
