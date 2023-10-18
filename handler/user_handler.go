package handler

import (
	"net/http"

	"com.test.users_api_test/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

func (uh *UserHandler) CreateUserAPiHandler(ctx *gin.Context) {
	uh.UserService.Context = ctx
	if newUser, err := uh.UserService.CreateUserAPiService(); err != nil {
		ResponseWithError(ctx.Writer, http.StatusInternalServerError, err.Error())
	} else {
		ResponseWithJson(ctx.Writer, http.StatusOK, newUser)
	}
}

// func HandlerSignInAPi(appContext *config.AppContext) {
// 	loginUserData := models.UserTable{}
// 	error := json.NewDecoder(r.Body).Decode(&loginUserData)
// 	if error != nil {
// 		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", error))
// 		return
// 	}

// 	if !validation.IsValidEmail(loginUserData.Email) {
// 		responseWithError(w, http.StatusBadRequest, fmt.Sprint("Enter a valid mail"))
// 	}

// 	result := gromDB.DB.Where("email = ? AND password = ?", loginUserData.Email, loginUserData.Password).Find(&loginUserData)
// 	if result.Error != nil {
// 		responseWithError(w, http.StatusBadRequest, result.Error.Error())
// 		return
// 	}

// 	if result.RowsAffected == 0 {
// 		responseWithError(w, http.StatusNotFound, "User not found")
// 		return
// 	}

// 	tokenString, err := utils.GenerateToken(loginUserData)
// 	if err != nil {
// 		// If there is an error in creating the JWT return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

//	responseModel := models.ResponseModel{
//		Message: "Success",
//		Data:    map[string]interface{}{"token": tokenString},
//	}
//
// responseWithJson(w, http.StatusOK, responseModel)
// }
