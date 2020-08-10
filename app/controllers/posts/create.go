package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create .
func Create(ctx *gin.Context) {
	var post Input
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
