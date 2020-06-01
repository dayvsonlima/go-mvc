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
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	db.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// PostsShow .
func PostsShow(ctx *gin.Context) {

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
