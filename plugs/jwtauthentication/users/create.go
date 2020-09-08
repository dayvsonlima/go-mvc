package users

import (
	"application/app/models"
	"net/http"

	"application/database"

	"github.com/gin-gonic/gin"
)

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

	user := &models.User{}
	user.SetEmail(credentials.Email)
	user.SetPassword(credentials.Email)
	result := database.ORM().Create(user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	ctx.JSON(http.StatusOK, user)
}
