package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"com.test.users_api_test/models"
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
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	isValid, errMessage := IsValidUser(user)
	if isValid {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint(errMessage))
		return
	}

	user.CreatedAt = time.Now()
	fmt.Printf("user is: %v\n", user)

	result := gormDb.DB.Create(&user)
	if result != nil {
		responseWithJson(w, http.StatusOK, user)
	}
}

func IsValidUser(user models.UserTable) (bool, string) {
	if !validation.IsValidEmail(user.Email) {
		return false, "Enter a valid mail"
	}
	if !validation.IsValidPassword(user.Password) {
		return false, "Enter a valid password"
	}
	return true, ""
}

func (gormDb *GormDB) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.UserTable{}
	queries := r.URL.Query()
	searchedName := queries.Get("name")

	// Check if the query parameter exists
	if searchedName != "" {
		likeValue := fmt.Sprintf("%%%s%%", searchedName)
		gormDb.DB.Where("name LIKE ?", likeValue).Find(&users)
	} else {
		gormDb.DB.Find(&users)
	}
	responseWithJson(w, http.StatusOK, users)
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
		responseWithError(w, http.StatusBadRequest, "")
		return
	}

	if result.RowsAffected == 0 {
		responseWithError(w, http.StatusNotFound, "User not found")
		return
	}

	responseWithJson(w, http.StatusOK, user)

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

	if updatedModel.Name != "" {
		user.Name = updatedModel.Name
	}

	if updatedModel.Email != "" && !validation.IsValidEmail(updatedModel.Email) {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint("Enter a valid mail"))
		return
	} else {
		user.Email = updatedModel.Email
	}

	if updatedModel.Password != "" && !validation.IsValidPassword(updatedModel.Password) {
		user.Password = updatedModel.Password
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
