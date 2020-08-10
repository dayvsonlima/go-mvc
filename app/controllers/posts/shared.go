package posts

import "application/database"

var (
	db = database.ORM()
)

type Input struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
