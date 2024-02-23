package user

import (
	"fmt"

	"com.test.users_api_test/api/models"
	"com.test.users_api_test/pkg/auth"
	"com.test.users_api_test/pkg/conventer"
	"com.test.users_api_test/pkg/validation"
	"github.com/gin-gonic/gin"
)

func CreateUserService(context *gin.Context) (*models.UserTable, error) {
	user := &models.UserTable{}

	if err := conventer.ParseRequestBody(context, user); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	if isValid, errMessage := validation.IsValidCreateUser(*user); !isValid {
		return nil, fmt.Errorf("validation error: %s", errMessage)
	}

	createdUser, err := CreateUserQuery(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func SignInService(context *gin.Context) (*models.ResponseModel, error) {
	loginUserData := models.UserTable{}
	error := conventer.ParseRequestBody(context, &loginUserData)
	if error != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", error)
	}

	if !validation.IsValidEmail(loginUserData.Email) {
		return nil, fmt.Errorf("enter a valid mail")
	}

	user, err := SingInQuery(loginUserData.Email, loginUserData.Password)
	if err != nil {
		return nil, err
	}

	tokenString, err := auth.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    map[string]interface{}{"token": tokenString},
	}
	return &responseModel, nil
}

func GetAllUsersService(ctx *gin.Context) (*models.ResponseModel, error) {
	queries := ctx.Request.URL.Query()
	searchedName := queries.Get("name")

	var users []models.UserTable
	if searchedName != "" {
		users = SearchByNameQuery(searchedName)
	} else {
		users = GetAllUsersQuery()
	}
	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    users,
	}
	return &responseModel, nil
}

func GetUserByIdService(ctx *gin.Context) (*models.ResponseModel, error) {
	idUint, err := conventer.ConvertStringToUint(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	user, err := GetUserByIdQuery(idUint)
	if err != nil {
		return nil, err
	}
	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    user,
	}
	return &responseModel, nil
}

func EditUserService(ctx *gin.Context) (*models.ResponseModel, error) {
	idUint, err := conventer.ConvertStringToUint(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	updatedModel := models.UserTable{}
	error := conventer.ParseRequestBody(ctx, &updatedModel)
	if error != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", error)
	}

	user, err := GetUserByIdQuery(idUint)
	if err != nil {
		return nil, err
	}

	isValid, errorr := validation.ValidateEditing(user, &updatedModel)
	if !isValid && errorr != nil {
		return nil, errorr
	}

	// check if email isn't found before
	// if updatedModel.Email != "" {
	// 	userslist := SearchByEmailQuery(updatedModel.Email)
	// 	if len(userslist) > 1 {
	// 		return nil, fmt.Errorf("%v This email is stored before ", updatedModel.Email)
	// 	}
	// }

	UpdateUserQuery(user)
	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    user,
	}
	return &responseModel, nil
}

func DeleteUserService(ctx *gin.Context) (*models.ResponseModel, error) {
	idUint, err := conventer.ConvertStringToUint(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	if _, err := DeletUserQuery(idUint); err != nil {
		return nil, err
	}
	responseModel := models.ResponseModel{
		Message: "Successfully deleted user",
		Data:    "",
	}
	return &responseModel, nil
}
