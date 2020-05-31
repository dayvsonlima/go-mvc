package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
