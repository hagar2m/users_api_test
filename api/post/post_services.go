package post

import (
	"errors"
	"fmt"

	"com.test.users_api_test/api/models"
	conventer "com.test.users_api_test/pkg/converter"
	"com.test.users_api_test/pkg/validation"
	"github.com/gin-gonic/gin"
)

func CreatePostService(context *gin.Context) (*models.Post, error) {

	// user check //
	userId := context.Value("userId").(uint)

	// post check //
	post := models.Post{}
	error := conventer.ParseRequestBody(context, &post)
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
