package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"com.test.users_api_test/models"
	"com.test.users_api_test/utils"
	"com.test.users_api_test/validation"
	"github.com/go-chi/chi"
)

func (gromDB *GormDB) HandlerCreateUserFromBrowser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("exampleInputEmail1")
	password := r.FormValue("exampleInputPassword1")
	name := r.FormValue("exampleInputName1")
	user := models.UserTable{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	result := gromDB.DB.Create(&user)
	if result != nil {
		responseWithJson(w, http.StatusOK, user)
	}
}

func (gormDb *GormDB) HandlerCreateUserFromAPi(w http.ResponseWriter, r *http.Request) {
	user := models.UserTable{}
	error := json.NewDecoder(r.Body).Decode(&user)
	if error != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	isValid, errMessage := validation.IsValidCreateUser(user)
	if !isValid {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint(errMessage))
		return
	}

	user.CreatedAt = time.Now()
	result := gormDb.DB.Create(&user)
	if result.Error != nil {
		responseWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	responseWithJson(w, http.StatusOK, user)

}

func (gromDB *GormDB) HandlerSignInAPi(w http.ResponseWriter, r *http.Request) {
	loginUserData := models.UserTable{}
	error := json.NewDecoder(r.Body).Decode(&loginUserData)
	if error != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	if !validation.IsValidEmail(loginUserData.Email) {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint("Enter a valid mail"))
	}

	result := gromDB.DB.Where("email = ? AND password = ?", loginUserData.Email, loginUserData.Password).Find(&loginUserData)
	if result.Error != nil {
		responseWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		responseWithError(w, http.StatusNotFound, "User not found")
		return
	}

	tokenString, err := utils.GenerateToken(loginUserData)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    map[string]interface{}{"token": tokenString},
	}
	responseWithJson(w, http.StatusOK, responseModel)
}

func (gormDb *GormDB) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.UserTable{}
	queries := r.URL.Query()
	searchedName := queries.Get("name")

	if searchedName != "" {
		likeValue := fmt.Sprintf("%%%s%%", searchedName)
		gormDb.DB.Where("name LIKE ?", likeValue).Find(&users)
	} else {
		gormDb.DB.Find(&users)
	}
	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    users,
	}

	// c1, err := r.Cookie("userCookie")

	// fmt.Printf("\n\n%+v %+v\n\n", r.Cookies(), c1)
	// if err != nil {
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:  "userCookie",
	// 		Value: "Hagar",
	// 	})
	// 	println(w, "New COOKIE Created #1:", c1)
	// } else {
	// 	println(w, "YOUR COOKIE #1:", c1)
	// }

	responseWithJson(w, http.StatusOK, responseModel)
}

func (gormDb *GormDB) HandlerGetUserById(w http.ResponseWriter, r *http.Request) {
	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user := models.UserTable{}
	result := gormDb.DB.Where("id = ?", idUint).Find(&user)
	if result.Error != nil {
		responseWithError(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		responseWithError(w, http.StatusNotFound, "User not found")
		return
	}

	responseModel := models.ResponseModel{
		Message: "Success",
		Data:    user,
	}
	responseWithJson(w, http.StatusOK, responseModel)
}

func (gormDb *GormDB) HandlerEditUser(w http.ResponseWriter, r *http.Request) {
	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user := models.UserTable{}
	result := gormDb.DB.Where("id = ?", idUint).Find(&user)

	updatedModel := models.UserTable{}
	error := json.NewDecoder(r.Body).Decode(&updatedModel)
	if error != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	isValid, errorr := validation.ValidateEditing(&user, &updatedModel)
	if !isValid && errorr != "" {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint(errorr))
		return
	}
	result = gormDb.DB.Save(&user)
	if result != nil {
		responseWithJson(w, http.StatusOK, user)
	}
}

func (gormDb *GormDB) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	idUint, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user := models.UserTable{
		ID: uint(idUint),
	}
	result := gormDb.DB.Where("id = ?", idUint).Delete(&user)
	if result.Error != nil {
		responseWithError(w, http.StatusNotFound, "failed to delete user: "+result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		responseWithError(w, http.StatusNotFound, "User not found")
		return
	}
	responseWithJson(w, http.StatusOK, models.ResponseModel{
		Message: "Successfully deleted user",
	})
}
