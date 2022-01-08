package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteRequest struct {
	Id    int
	Link  string
	Token string
}

func GetIndex(c *gin.Context) {
	//query := c.Query("query")
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

func PostNonsense(c *gin.Context) {
	description := c.PostForm("description")
	if description == "" {
		description = "No description."
	}
	uploader := c.PostForm("uploader")
	if uploader == "" {
		uploader = "Anonymous User"
	}
	//currentTime := time.Now().Format("2006-01-02 15:04:05")
	c.Redirect(http.StatusSeeOther, "./")
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
