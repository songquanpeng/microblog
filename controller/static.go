package controller

import (
	"github.com/gin-gonic/gin"
	"microblog/common"
	"net/http"
)

func GetIndex(c *gin.Context) {
	html, _ := common.FS.ReadFile("public/index.html")
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

func GetStaticFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/static/"+path, http.FS(common.FS))
}

func GetLibFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/lib/"+path, http.FS(common.FS))
}
