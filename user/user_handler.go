package user

import (
	"net/http"

	"com.test.users_api_test/handler"
	"github.com/gin-gonic/gin"
)

func CreateUserAPiHandler(ctx *gin.Context) {
	if newUser, err := CreateUserService(ctx); err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, newUser)
	}
}

func SingInAPiHandler(ctx *gin.Context) {
	if newUser, err := SignInService(ctx); err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, newUser)
	}
}

func GetAllUsersHandler(ctx *gin.Context) {
	if users, err := GetAllUsersService(ctx); err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, users)
	}
}

func GetUserByIdhandler(ctx *gin.Context) {
	if users, err := GetUserByIdService(ctx); err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, users)
	}
}

func EditUserhandler(ctx *gin.Context) {
	if users, err := EditUserService(ctx); err != nil {
		handler.ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		handler.ResponseWithJson(ctx.Writer, http.StatusOK, users)
	}
}
