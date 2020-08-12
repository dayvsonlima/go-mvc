package posts

import (
	"net/http"

	"application/app/models"

	"github.com/gin-gonic/gin"
)

type UpdateInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Update .
func Update(ctx *gin.Context) {
	var post models.Post
	if err := db.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&post).Updates(input)

	ctx.JSON(http.StatusOK, post)
}
