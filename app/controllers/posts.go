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
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "post not found!"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// PostsNew .
func PostsNew(ctx *gin.Context) {

}

// PostsCreate .
func PostsCreate(ctx *gin.Context) {

}

// PostsUpdate .
func PostsUpdate(ctx *gin.Context) {

}

// PostsDelete .
func PostsDelete(ctx *gin.Context) {

}
