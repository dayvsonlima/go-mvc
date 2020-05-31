package config

import (
	c "application/app/controllers"

	"github.com/gin-gonic/gin"
)

// DrawRoutes ...
func DrawRoutes(routes *gin.Engine) *gin.Engine {

	// Posts resources
	routes.GET("/posts", c.PostIndex)
	routes.GET("/posts/:id", c.PostShow)
	routes.POST("/posts", c.PostCreate)
	routes.PUT("/posts", c.PostUpdate)
	routes.DELETE("/posts", c.PostDelete)

	return routes
}
