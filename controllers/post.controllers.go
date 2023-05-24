package controllers

import (
	"backend-demo/initialzers"
	"backend-demo/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initialzers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostFindAll(c *gin.Context) {
	var posts []models.Post

	initialzers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostFindOne(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initialzers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initialzers.DB.First(&post, id)

	initialzers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initialzers.DB.Delete(&post, id)

	c.Status(200)
}
