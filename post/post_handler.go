package post

import (
	"net/http"

	"com.test.users_api_test/handler"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	post, err := CreatePostService(ctx)
	if err != nil {
		httpError, ok := err.(*handler.HTTPError)
		if ok {
			ctx.JSON(httpError.Status, gin.H{"error": httpError.Message})
		} else {
			handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
		}
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, post)
	}

	//	 err != nil
	//	 {
	//		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	//	} else {
	//
	//		handler.ResponseWithJson(ctx.Writer, http.StatusOK, newUser)
	//	}
}
