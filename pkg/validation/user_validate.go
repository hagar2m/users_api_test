package validation

import (
	"fmt"

	"com.test.users_api_test/api/models"
)

func ValidateEditing(usr *models.UserTable, updatedModel *models.UserTable) (bool, error) {
	if updatedModel.Name != "" {
		usr.Name = updatedModel.Name
	}

	if updatedModel.Email != "" && !IsValidEmail(updatedModel.Email) {
		return false, fmt.Errorf("enter a valid mail")
	} else {
		usr.Email = updatedModel.Email
	}

	if updatedModel.Password != "" && !IsValidPassword(updatedModel.Password) {
		usr.Password = updatedModel.Password
	}
	return true, nil
}

func IsValidCreateUser(user models.UserTable) (bool, string) {
	if !IsValidEmail(user.Email) {
		return false, "Enter a valid mail"
	}
	if !IsValidPassword(user.Password) {
		return false, "Enter a valid password"
	}
	return true, ""
}
