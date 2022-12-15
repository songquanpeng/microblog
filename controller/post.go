package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"microblog/model"
	"net/http"
	"strconv"
	"time"
)

func GetAllPosts(c *gin.Context) {
	pStr := c.DefaultQuery("p", "0")
	p, _ := strconv.Atoi(pStr)
	if p < 0 {
		p = 0
	}
	pageSize := 10
	var posts []*model.Post
	authed := c.GetBool("authed")
	var err error
	if authed {
		err = model.DB.Limit(pageSize).Offset(p).Order("id desc").Find(&posts).Error
	} else {
		err = model.DB.Limit(pageSize).Offset(p).Order("id desc").Where("status = ?", model.PostStatusPublic).Find(&posts).Error
	}
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
		"data":    posts,
	})
	return
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var post model.Post
	err := model.DB.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if post.Status != model.PostStatusPublic {
		authed := c.GetBool("authed")
		if !authed {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无权限获取此微博",
				"data":    nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    post,
	})
	return
}

func CreatePost(c *gin.Context) {
	var post model.Post
	err := json.NewDecoder(c.Request.Body).Decode(&post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if post.Content == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "内容不能为空",
		})
		return
	}
	post.Timestamp = time.Now().Unix()
	err = post.Insert()
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
		"data":    post.Id,
	})
	return
}

func UpdatePost(c *gin.Context) {
	var post model.Post
	err := json.NewDecoder(c.Request.Body).Decode(&post)
	if err != nil || post.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	err = post.Update()
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

func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	post := model.Post{Id: id}
	// TODO: What will happen if delete a non-existed record?
	err := post.Delete()
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
