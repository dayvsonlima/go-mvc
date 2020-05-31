package controllers

import (
	"application/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db = database.Connection
)

// PostsIndex .
func PostsIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "posts/index", gin.H{})
}

// PostsShow .
func PostsShow(ctx *gin.Context) {

}

// PostsNew .
func PostsNew(ctx *gin.Context) {

}

// PostsCreate .
func PostsCreate(ctx *gin.Context) {

}

// PostsUpdate .
func PostsUpdate(ctx *gin.Context) {

}

// PostsDelete .
func PostsDelete(ctx *gin.Context) {

}
