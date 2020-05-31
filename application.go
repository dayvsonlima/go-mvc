package main

import (
	"application/config"
	"application/database"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.LoadHTMLGlob("app/views/**/*")
	config.DrawRoutes(engine)

	database.Migrations()
	defer database.Connection.Close()

	engine.Run(":8080")
}
