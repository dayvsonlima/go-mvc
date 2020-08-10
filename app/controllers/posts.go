package controllers

import (
	"application/app/models"
	"application/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db = database.ORM()
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// PostsIndex .
func PostsIndex(ctx *gin.Context) {
	var posts []models.Post
	db.Find(&posts)
	ctx.JSON(http.StatusOK, posts)
}

// PostsShow .
func PostsShow(ctx *gin.Context) {
	var post models.Post
	db.First(&post, ctx.Param("id"))

	if post.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "post not found!",
		})

		return
	}

	ctx.JSON(http.StatusOK, post)
}

// PostsCreate .
func PostsCreate(ctx *gin.Context) {

	var post CreatePostInput
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost := &models.Post{
		Title:   post.Title,
		Content: post.Content,
	}

	db.Create(newPost)
	ctx.JSON(http.StatusOK, newPost)
}

// PostsUpdate .
func PostsUpdate(ctx *gin.Context) {

}

// PostsDelete .
func PostsDelete(ctx *gin.Context) {

}
