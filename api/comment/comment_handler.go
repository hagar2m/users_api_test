package comment

import (
	"net/http"

	"com.test.users_api_test/handler"
	"github.com/gin-gonic/gin"
)

func CreateCommentHandler(ctx *gin.Context) {
	post, err := CreateCommentService(ctx)
	if err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusBadRequest, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, post)
	}
}
func GetCommentByIDHandler(ctx *gin.Context) {
	post, err := GetCommentByIdService(ctx)
	if err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusBadRequest, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, post)
	}
}
