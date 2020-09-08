package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Credentials .
// must be filled with user credentials
type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Create .
func Create(ctx *gin.Context) {
	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
