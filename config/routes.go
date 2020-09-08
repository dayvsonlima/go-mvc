package config

import (
	"application/plugs/jwtauthentication/login"
	"application/plugs/jwtauthentication/users"

	"github.com/gin-gonic/gin"
)

// DrawRoutes .
func DrawRoutes(routes *gin.Engine) {
	routes.POST("/login", login.Create)
	routes.POST("/login/renew", login.Refresh)
	routes.POST("/users", users.Create)

}
