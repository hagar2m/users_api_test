package comment

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

func CreateCommentService(context *gin.Context) (*models.Comment, error) {
	// user check //
	userId := context.Value("userId").(uint)
	userName := context.Value("name").(string)

	err := context.Request.ParseMultipartForm(5 << 20) // 5 MB max file size
	if err != nil {
		return nil, err
	}

	// Validate input
	input := models.Comment{}
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
	comment := models.Comment{
		Body:     input.Body,
		PostID:   input.PostID,
		Image:    imgUrl,
		UserName: userName,
		UserID:   userId,
	}

	if isValid, errMessage := validation.IsValidCreateComment(comment); !isValid {
		return nil, fmt.Errorf(errMessage)
	}
	// Create comment
	createdComment, errr := CreateCommentQuery(&comment)
	if errr != nil {
		return nil, errr
	}
	return createdComment, nil
}
