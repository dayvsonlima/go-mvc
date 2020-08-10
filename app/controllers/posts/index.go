package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostsIndex .
func Index(ctx *gin.Context) {
	var posts []models.Post
	db.Find(&posts)
	ctx.JSON(http.StatusOK, posts)
}
