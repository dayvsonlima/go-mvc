package posts

import (
	"application/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostsDelete .
func Delete(ctx *gin.Context) {
	var post models.Post
	if err := db.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&post)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
