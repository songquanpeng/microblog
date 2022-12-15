package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"microblog/common"
	"net/http"
)

func ApiAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		password := session.Get("password")
		if password == nil || password != common.Password {
			c.Set("authed", false)
		} else {
			c.Set("authed", true)
		}
		c.Next()
	}
}

func AuthRequired() func(c *gin.Context) {
	return func(c *gin.Context) {
		if !c.GetBool("authed") {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，未登录或登录已过期",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
