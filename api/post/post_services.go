package post

import (
	"errors"
	"fmt"

	"com.test.users_api_test/api/models"
	"com.test.users_api_test/pkg/validation"
	"github.com/gin-gonic/gin"
)

func CreatePostService(context *gin.Context) (*models.Post, error) {

	// user check //
	userId := context.Value("userId").(uint)

	// Validate input
	input := models.Post{}
	if err := context.ShouldBind(&input); err != nil {
		return nil, errors.New(fmt.Sprintf("%s", err))
	}

	post := models.Post{Title: input.Title, Body: input.Body}

	if isValid, errMessage := validation.IsValidCreatePost(post); !isValid {
		return nil, fmt.Errorf(errMessage)
	}
	// Create post
	createdPost, errr := CreatePostQuery(&post, userId)
	if errr != nil {
		return nil, errr
	}
	return createdPost, nil
}
