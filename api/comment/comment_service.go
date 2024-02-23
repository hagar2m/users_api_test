package comment

import (
	"errors"
	"fmt"
	"net/http"

	"com.test.users_api_test/api/constants"
	"com.test.users_api_test/api/models"
	"com.test.users_api_test/configs"
	"com.test.users_api_test/handler"
	"com.test.users_api_test/pkg/conventer"
	"com.test.users_api_test/pkg/validation"
	"github.com/gin-gonic/gin"
)

func CreateCommentService(context *gin.Context) (*models.Comment, error) {
	// user check //
	userId := context.Value("userId").(uint)
	userName := context.Value("name").(string)

	// Try parsing as JSON
	var input models.Comment
	if err := context.ShouldBindJSON(&input); err == nil {
		return createCommentFromInput(userId, userName, input)
	}

	err := context.Request.ParseMultipartForm(5 << 20) // 5 MB max file size
	if err != nil {
		return nil, err
	}

	// Validate input
	if err := context.ShouldBind(&input); err != nil {
		return nil, errors.New(fmt.Sprintf("%s", err))
	}

	// Check if an image file is uploaded
	file, fileHeader, err := context.Request.FormFile("image")
	if err != nil && err != http.ErrMissingFile { // Ignore ErrMissingFile error
		return nil, err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	var imgUrl string
	if file != nil {
		imgName, _ := handler.UploadFileHandler(file, fileHeader)
		imgUrl = configs.GetServerHost() + ":" + configs.GetServerPort() + constants.UploadRoot + "/" + imgName
	}
	input.Image = imgUrl

	return createCommentFromInput(userId, userName, input)
}

func createCommentFromInput(userId uint, userName string, input models.Comment) (*models.Comment, error) {

	comment := models.Comment{
		Body:     input.Body,
		PostID:   input.PostID,
		Image:    input.Image,
		UserName: userName,
		UserID:   userId,
		ParentID: input.ParentID,
		Comments: input.Comments,
	}

	if isValid, errMessage := validation.IsValidCreateComment(comment); !isValid {
		return nil, fmt.Errorf(errMessage)
	}

	// Create the comment in the database
	createdComment, err := CreateCommentQuery(&comment)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func GetCommentByIdService(ctx *gin.Context) (*models.Comment, error) {
	id, err := conventer.ConvertStringToInt(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	comment, err := GetCommentByIdQuery(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
