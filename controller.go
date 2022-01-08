package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "okay",
			"data":    nonsenses,
		})
	}

}

func PostNonsense(c *gin.Context) {
	var postRequest PostRequest
	err := json.NewDecoder(c.Request.Body).Decode(&postRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid parameter",
		})
		return
	}
	if *Token == postRequest.Token {
		nonsenseObj := &Nonsense{
			Content: postRequest.Content,
			Time:    time.Now().Format("2006-01-02 15:04:05"),
		}
		err := nonsenseObj.Insert()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
			})
		} else {
			message := "Nonsense posted successfully."
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": message,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Token is invalid.",
		})
	}
}

func DeleteNonsense(c *gin.Context) {
	var deleteRequest DeleteRequest
	err := json.NewDecoder(c.Request.Body).Decode(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid parameter",
		})
		return
	}
	if *Token == deleteRequest.Token {
		fileObj := &Nonsense{
			Id: deleteRequest.Id,
		}
		DB.Where("id = ?", deleteRequest.Id).First(&fileObj)
		err := fileObj.Delete()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
			})
		} else {
			message := "Nonsense deleted successfully."
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": message,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Token is invalid.",
		})
	}
}
