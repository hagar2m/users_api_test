package post

import (
	"errors"
	"fmt"

	"com.test.users_api_test/api/constants"
	"com.test.users_api_test/api/models"
	"com.test.users_api_test/configs"
	"com.test.users_api_test/handler"
	"com.test.users_api_test/pkg/validation"
	"github.com/gin-gonic/gin"
)

func CreatePostService(context *gin.Context) (*models.Post, error) {

	// user check //
	userId := context.Value("userId").(uint)

	err := context.Request.ParseMultipartForm(5 << 20) // 5 MB max file size
	if err != nil {
		return nil, err
	}

	// Validate input
	input := models.Post{}
	if err := context.ShouldBind(&input); err != nil {
		return nil, errors.New(fmt.Sprintf("%s", err))
	}

	// Get the uploaded file
	file, fileHeader, err := context.Request.FormFile("image")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imgName, _ := handler.UploadFileHandler(file, fileHeader)
	imgUrl := configs.GetServerHost() + ":" + configs.GetServerPort() + constants.UploadRoot + "/" + imgName
	post := models.Post{
		Title: input.Title,
		Body:  input.Body,
		Image: imgUrl,
	}

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
