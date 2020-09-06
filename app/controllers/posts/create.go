package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Create .
func Create(ctx *gin.Context) {
	var input CreateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PostDB.CollectionMutex.Lock()
	PostDB.Count++

	id := PostDB.Count

	post := &models.Post{
		ID:      id,
		Title:   input.Title,
		Content: input.Content,
	}

	PostDB.Collection[id] = post
	PostDB.CollectionMutex.Unlock()

	ctx.JSON(http.StatusOK, post)
}
