package models

import "github.com/jinzhu/gorm"

// Post .
type Post struct {
	gorm.Model `json:"-"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
}
