package posts

import (
	"application/app/models"
	"sync"
)

var (
	PostDB = PostStore{
		Count:           0,
		Collection:      make(map[int]*models.Post),
		CollectionMutex: sync.RWMutex{},
	}
)

// PostCollection .
type PostCollection map[int]*models.Post

// PostStore .
type PostStore struct {
	Count           int
	Collection      PostCollection
	CollectionMutex sync.RWMutex
}
