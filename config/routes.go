package config

import (
	"application/plugs/jwtauthentication"

	"github.com/gin-gonic/gin"
)

// DrawRoutes .
func DrawRoutes(routes *gin.Engine) {
	routes.POST("/login", jwtauthentication.Create)
	routes.POST("/login/renew", jwtauthentication.Refresh)
}
