package config

import (
	c "application/app/controllers"

	"github.com/gin-gonic/gin"
)

// DrawRoutes .
func DrawRoutes(routes *gin.Engine) {
	routes.GET("/posts", c.PostsIndex)
	routes.GET("/posts/new", c.PostsNew)
	routes.POST("/posts", c.PostsCreate)
	routes.GET("/post/:id", c.PostsShow)
	routes.PUT("/post/:id", c.PostsUpdate)
	routes.DELETE("/post/:id", c.PostsDelete)
}
