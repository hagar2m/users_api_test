package post

import (
	"errors"
	"fmt"

	"com.test.users_api_test/api/constants"
	"com.test.users_api_test/api/models"
	"com.test.users_api_test/configs"
	"com.test.users_api_test/handler"
	"com.test.users_api_test/pkg/validation"
	"com.test.users_api_test/pkg/conventer"
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
		UserID: userId,
	}

	if isValid, errMessage := validation.IsValidCreatePost(post); !isValid {
		return nil, fmt.Errorf(errMessage)
	}
	// Create post
	createdPost, errr := CreatePostQuery(&post)
	if errr != nil {
		return nil, errr
	}
	return createdPost, nil
}

func GetPostCommentsService(context *gin.Context) (*models.Post, error) {
	id, err := conventer.ConvertStringToInt(context.Param("id"))
	if err != nil {
		return nil, err
	}
	post, err := GetPostCommentsByIdQuery(id)
	if err != nil {
		return nil, err
	}
	
	return post, nil
}