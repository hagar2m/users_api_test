package post

import (
	"net/http"

	"com.test.users_api_test/handler"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	post, err := CreatePostService(ctx)
	if err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusBadRequest, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, post)
	}
}

func GetPostCommentsHandler(ctx *gin.Context) {
	post, err := GetPostCommentsService(ctx)
	if err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusBadRequest, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, post)
	}
}