package controller

import (
	"github.com/gin-gonic/gin"
	"microblog/common"
	"net/http"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data": gin.H{
			"version": common.Version,
			"theme":   common.Theme,
			"author":  common.Username,
			"authed":  c.GetBool("authed"),
		},
	})
}
