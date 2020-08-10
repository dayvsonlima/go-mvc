package config

import (
	c "application/app/controllers"

	"github.com/gin-gonic/gin"
)

// DrawRoutes .
func DrawRoutes(routes *gin.Engine) {
	routes.GET("/posts", c.PostsIndex)
	routes.POST("/posts", c.PostsCreate)
	routes.GET("/posts/:id", c.PostsShow)
	routes.PUT("/posts/:id", c.PostsUpdate)
	routes.PATCH("/posts/:id", c.PostsUpdate)
	routes.DELETE("/posts/:id", c.PostsDelete)
}
