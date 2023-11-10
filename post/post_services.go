package post

import (
	"fmt"
	"net/http"

	"com.test.users_api_test/handler"
	"com.test.users_api_test/models"
	"com.test.users_api_test/user"
	"com.test.users_api_test/utils"
	"com.test.users_api_test/validation"
	"github.com/gin-gonic/gin"
)

func CreatePostService(context *gin.Context) (*models.Post, error) {
	// ctx, err := utils.ValidateToken(context.Writer, context.Request)
	// if err != nil {
	// 	return nil, &handler.HTTPError{Status: http.StatusBadRequest, Message: fmt.Sprintf("error parsing JSON: %v", err)}
	// }
	// user check //
	userId := context.Value("userId").(uint)
	_, err := user.GetUserByIdQuery(userId)
	if err != nil {
		return nil, err
	}

	// post check //
	post := models.Post{}
	error := utils.ParseRequestBody(context, &post)
	if error != nil {
		return nil, &handler.HTTPError{Status: http.StatusBadRequest, Message: fmt.Sprintf("error parsing JSON: %v", error)}
	}

	if isValid, errMessage := validation.IsValidCreatePost(post); !isValid {
		return nil, &handler.HTTPError{Status: http.StatusBadRequest, Message: fmt.Sprintf(errMessage)}
	}

	createdPost, errr := CreatePostQuery(&post, userId)
	if errr != nil {
		return nil, errr
	}
	return createdPost, nil
}
