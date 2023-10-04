package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"com.test.users_api_test/models"
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
	user.CreatedAt = time.Now()
	fmt.Printf("user is: %v\n", user)

	result := gormDb.DB.Create(&user)
	if result != nil {
		responseWithJson(w, http.StatusOK, user)
	}
}

func (gormDb *GormDB) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.UserTable{}
	gormDb.DB.Find(&users)
	responseWithJson(w, http.StatusOK, users)
}
