package config

import (
	"application/app/controllers/posts"

	"github.com/gin-gonic/gin"
)

// DrawRoutes .
func DrawRoutes(routes *gin.Engine) {
	routes.GET("/posts", posts.Index)
	routes.POST("/posts", posts.Create)
	routes.GET("/posts/:id", posts.Show)
	routes.PUT("/posts/:id", posts.Update)
	routes.PATCH("/posts/:id", posts.Update)
	routes.DELETE("/posts/:id", posts.Delete)
}
