package main

import (
	"application/config"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.LoadHTMLGlob("app/views/**/*")
	config.DrawRoutes(engine)

	engine.Run(":8080")
}
