package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"com.test.users_api_test/models"
	"com.test.users_api_test/utils"
	"com.test.users_api_test/validation"
)

func (gormDb *GormDB) HandlerCreatePost(w http.ResponseWriter, r *http.Request) {
	isvalid, err, ctx := utils.ValidateToken(w, r)
	if !isvalid {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	// user check //
	userId := ctx.Value("userId").(uint)
	user := models.UserTable{}
	userResult := gormDb.DB.Where("id = ?", userId).Find(&user)
	print("id - ", user.ID)
	if userResult.Error != nil {
		responseWithError(w, http.StatusBadRequest, userResult.Error.Error())
		return
	}
	if userResult.RowsAffected == 0 {
		responseWithError(w, http.StatusNotFound, "User not found")
		return
	}
	// post check //
	post := models.Post{}
	error := json.NewDecoder(r.Body).Decode(&post)
	if error != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
		return
	}

	isValid, errMessage := validation.IsValidCreatePost(post)
	if !isValid {
		responseWithError(w, http.StatusBadRequest, fmt.Sprint(errMessage))
		return
	}

	post.UserID = userId
	post.CreatedAt = time.Now()
	result := gormDb.DB.Create(&post)
	if result.Error != nil {
		responseWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	responseWithJson(w, http.StatusOK, post)
}
