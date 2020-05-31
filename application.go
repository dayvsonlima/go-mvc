package main

import (
	"application/config"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	application := config.DrawRoutes(engine)
	application.Run(":8080")
}
