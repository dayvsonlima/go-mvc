package posts

import (
	"application/app/models"
)

var (
	PostDB = PostStore{
		Last:       0,
		Collection: make(map[int]models.Post),
	}
)

// PostCollection .
type PostCollection map[int]models.Post

// PostStore .
type PostStore struct {
	Last       int
	Collection PostCollection
}
