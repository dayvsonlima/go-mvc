package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Show .
func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post := *PostDB.Collection[id]
	ctx.JSON(http.StatusOK, post)
}
