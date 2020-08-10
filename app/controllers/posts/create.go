package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create .
func Create(ctx *gin.Context) {
	var input Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := &models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	db.Create(post)
	ctx.JSON(http.StatusOK, post)
}
