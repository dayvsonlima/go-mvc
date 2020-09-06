package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Count .
func Count(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Quantity": PostDB.Count,
	})
}
