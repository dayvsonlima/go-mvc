package database

import (
	"application/app/models"

	"github.com/jinzhu/gorm"

	// stupid vscode rule
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// Connection .
	Connection *gorm.DB
)

// Setup .
func init() {
	Connection, _ = gorm.Open("sqlite3", "development.db")
}

// Migrations .
func Migrations() {
	Connection.AutoMigrate(&models.Post{})
}
