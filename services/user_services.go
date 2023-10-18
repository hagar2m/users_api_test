package services

import (
	"encoding/json"
	"fmt"
	"time"

	"com.test.users_api_test/app"
	"com.test.users_api_test/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"com.test.users_api_test/validation"
)

type UserService struct {
	DB      *gorm.DB
	Context *gin.Context
}

func NewUserService(appConfig *app.App) *UserService {
	return &UserService{
		DB: appConfig.DB,
	}
}

func (us *UserService) CreateUserAPiService() (*models.UserTable, error) {
	user := &models.UserTable{}

	if err := json.NewDecoder(us.Context.Request.Body).Decode(user); err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}

	if isValid, errMessage := validation.IsValidCreateUser(*user); !isValid {
		return nil, fmt.Errorf("Validation error: %s", errMessage)
	}

	user.CreatedAt = time.Now()
	if err := us.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func (us *UserService) HandlerSignInAPiService(appContext *config.AppContext) (*models.ResponseModel, string) {
// 	loginUserData := models.UserTable{}
// 	error := json.NewDecoder(appContext.Context.Request.Body).Decode(&loginUserData)
// 	if error != nil {
// 		return nil, fmt.Sprintf("Error parsing JSON: %v", error)
// 	}

// 	if !validation.IsValidEmail(loginUserData.Email) {
// 		return nil, fmt.Sprint("Enter a valid mail")
// 	}

// 	result := us.DB.Where("email = ? AND password = ?", loginUserData.Email, loginUserData.Password).Find(&loginUserData)
// 	if result.Error != nil {
// 		return nil, result.Error.Error()
// 	}

// 	if result.RowsAffected == 0 {
// 		return nil, "User not found"
// 	}

// 	tokenString, err := utils.GenerateToken(loginUserData)
// 	if err != nil {
// 		return nil, err.Error()
// 	}

// 	responseModel := models.ResponseModel{
// 		Message: "Success",
// 		Data:    map[string]interface{}{"token": tokenString},
// 	}
// 	return &responseModel, ""
// }

// func (gormDb *GormDB) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	users := []models.UserTable{}
// 	queries := r.URL.Query()
// 	searchedName := queries.Get("name")

// 	if searchedName != "" {
// 		likeValue := fmt.Sprintf("%%%s%%", searchedName)
// 		gormDb.DB.Where("name LIKE ?", likeValue).Find(&users)
// 	} else {
// 		gormDb.DB.Find(&users)
// 	}
// 	responseModel := models.ResponseModel{
// 		Message: "Success",
// 		Data:    users,
// 	}

// 	// c1, err := r.Cookie("userCookie")

// 	// fmt.Printf("\n\n%+v %+v\n\n", r.Cookies(), c1)
// 	// if err != nil {
// 	// 	http.SetCookie(w, &http.Cookie{
// 	// 		Name:  "userCookie",
// 	// 		Value: "Hagar",
// 	// 	})
// 	// 	println(w, "New COOKIE Created #1:", c1)
// 	// } else {
// 	// 	println(w, "YOUR COOKIE #1:", c1)
// 	// }

// 	responseWithJson(w, http.StatusOK, responseModel)
// }

// func (gormDb *GormDB) HandlerGetUserById(w http.ResponseWriter, r *http.Request) {
// 	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	user := models.UserTable{}
// 	result := gormDb.DB.Where("id = ?", idUint).Find(&user)
// 	if result.Error != nil {
// 		responseWithError(w, http.StatusBadRequest, result.Error.Error())
// 		return
// 	}

// 	if result.RowsAffected == 0 {
// 		responseWithError(w, http.StatusNotFound, "User not found")
// 		return
// 	}

// 	responseModel := models.ResponseModel{
// 		Message: "Success",
// 		Data:    user,
// 	}
// 	responseWithJson(w, http.StatusOK, responseModel)
// }

// func (gormDb *GormDB) HandlerEditUser(w http.ResponseWriter, r *http.Request) {
// 	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	user := models.UserTable{}
// 	result := gormDb.DB.Where("id = ?", idUint).Find(&user)

// 	updatedModel := models.UserTable{}
// 	error := json.NewDecoder(r.Body).Decode(&updatedModel)
// 	if error != nil {
// 		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
// 		return
// 	}

// 	isValid, errorr := validation.ValidateEditing(&user, &updatedModel)
// 	if !isValid && errorr != "" {
// 		responseWithError(w, http.StatusBadRequest, fmt.Sprint(errorr))
// 		return
// 	}
// 	result = gormDb.DB.Save(&user)
// 	if result != nil {
// 		responseWithJson(w, http.StatusOK, user)
// 	}
// }

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
