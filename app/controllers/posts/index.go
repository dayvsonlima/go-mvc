package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index .
func Index(ctx *gin.Context) {
	var posts = make([]models.Post, 0, len(PostDB.Collection))

	for _, value := range PostDB.Collection {
		posts = append(posts, value)
	}

	ctx.JSON(http.StatusOK, posts)
}
