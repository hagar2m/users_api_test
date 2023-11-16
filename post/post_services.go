package post

import (
	"errors"
	"fmt"

	"com.test.users_api_test/models"
	"com.test.users_api_test/utils"
	"com.test.users_api_test/validation"
	"github.com/gin-gonic/gin"
)

func CreatePostService(context *gin.Context) (*models.Post, error) {

	// user check //
	userId := context.Value("userId").(uint)

	// post check //
	post := models.Post{}
	error := utils.ParseRequestBody(context, &post)
	if error != nil {
		return nil, errors.New(fmt.Sprintf("error parsing JSON: %v", error))
	}

	if isValid, errMessage := validation.IsValidCreatePost(post); !isValid {
		return nil, fmt.Errorf(errMessage)
	}

	createdPost, errr := CreatePostQuery(&post, userId)
	if errr != nil {
		return nil, errr
	}
	return createdPost, nil
}
