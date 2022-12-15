package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PostRequest struct {
	Token   string `json:"token"`
	Content string `json:"content"`
}

type DeleteRequest struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func GetIndex(c *gin.Context) {
	html, _ := fs.ReadFile("public/index.html")
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

func GetStaticFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/static/"+path, http.FS(fs))
}

func GetLibFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/lib/"+path, http.FS(fs))
}

func GetNonsenseById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var nonsense Nonsense
	if err := DB.First(&nonsense, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    nonsense,
		})
	}

}

func GetNonsense(c *gin.Context) {
	pStr := c.DefaultQuery("p", "0")
	p, _ := strconv.Atoi(pStr)
	if p < 0 {
		p = 0
	}
	pageSize := 10
	var nonsenses []*Nonsense
	if err := DB.Limit(pageSize).Offset(p).Order("id desc").Find(&nonsenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "",
			"data":    nonsenses,
		})
		return
	}

}

func PostNonsense(c *gin.Context) {
	var postRequest PostRequest
	err := json.NewDecoder(c.Request.Body).Decode(&postRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	if *Token != postRequest.Token {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的凭证",
		})
		return
	}
	// Enable user to delete nonsense by type `delete id`
	if strings.HasPrefix(postRequest.Content, "delete ") {
		t := strings.Split(postRequest.Content, " ")
		idStr := t[len(t)-1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无效的参数",
			})
			return
		}
		var nonsenseObj Nonsense
		DB.Where("id = ?", id).First(&nonsenseObj)
		if nonsenseObj.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": fmt.Sprintf("无效的微博 ID: %s", idStr),
			})
			return
		}
		err = nonsenseObj.Delete()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
		})
		return
	}
	nonsenseObj := &Nonsense{
		Content: postRequest.Content,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err = nonsenseObj.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    nonsenseObj.Id,
	})
	return
}

func DeleteNonsense(c *gin.Context) {
	var deleteRequest DeleteRequest
	err := json.NewDecoder(c.Request.Body).Decode(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	if *Token != deleteRequest.Token {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的凭证",
		})
		return
	}
	nonsenseObj := &Nonsense{
		Id: deleteRequest.Id,
	}
	DB.Where("id = ?", deleteRequest.Id).First(&nonsenseObj)
	err = nonsenseObj.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
