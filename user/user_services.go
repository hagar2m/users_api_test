package user

import (
	"encoding/json"
	"fmt"
	"strconv"

	"com.test.users_api_test/models"
	"com.test.users_api_test/utils"
	"github.com/gin-gonic/gin"

	"com.test.users_api_test/validation"
)

func CreateUserService(context *gin.Context) (*models.UserTable, error) {
	user := &models.UserTable{}

	if err := json.NewDecoder(context.Request.Body).Decode(user); err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}

	if isValid, errMessage := validation.IsValidCreateUser(*user); !isValid {
		return nil, fmt.Errorf("Validation error: %s", errMessage)
	}

	user, err := CreateUserQuery(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SignInService(context *gin.Context) (*models.ResponseModel, error) {
	loginUserData := models.UserTable{}
	error := json.NewDecoder(context.Request.Body).Decode(&loginUserData)
	if error != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", error)
	}

	if !validation.IsValidEmail(loginUserData.Email) {
		return nil, fmt.Errorf("Enter a valid mail")
	}

	user, err := SingInQuery(loginUserData.Email, loginUserData.Password)
	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateToken(*user)
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
	users := []models.UserTable{}
	queries := ctx.Request.URL.Query()
	searchedName := queries.Get("name")

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
	idUint, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
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
	idUint, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	updatedModel := models.UserTable{}
	error := json.NewDecoder(ctx.Request.Body).Decode(&updatedModel)
	if error != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", error)
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
	if updatedModel.Email != "" {
		userslist := SearchByEmailQuery(updatedModel.Email)
		if len(userslist) > 0 {
			return nil, fmt.Errorf("%v This email is stored before ", updatedModel.Email)
		}
	}

	UpdateUserQuery(user)
	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    user,
	}
	return &responseModel, nil
}

// func (gormDb *GormDB) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
// 	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	user := models.UserTable{
// 		ID: uint(idUint),
// 	}
// 	result := gormDb.DB.Where("id = ?", idUint).Delete(&user)
// 	if result.Error != nil {
// 		responseWithError(w, http.StatusNotFound, "failed to delete user: "+result.Error.Error())
// 		return
// 	} else if result.RowsAffected == 0 {
// 		responseWithError(w, http.StatusNotFound, "User not found")
// 		return
// 	}
// 	responseWithJson(w, http.StatusOK, models.ResponseModel{
// 		Message: "Successfully deleted user",
// 	})
// }

// func CreateUserFromBrowser(w http.ResponseWriter, r *http.Request) {
// 	email := r.FormValue("exampleInputEmail1")
// 	password := r.FormValue("exampleInputPassword1")
// 	name := r.FormValue("exampleInputName1")
// 	user := models.UserTable{
// 		Name:      name,
// 		Email:     email,
// 		Password:  password,
// 		CreatedAt: time.Now(),
// 	}
// 	result := config.DB.Create(&user)
// 	if result != nil {
// 		handler.ResponseWithJson(w, http.StatusOK, user)
// 	}
// }
