package validation

import "com.test.users_api_test/models"

func ValidateEditing(usr *models.UserTable, updatedModel *models.UserTable) (bool, string) {
	if updatedModel.Name != "" {
		usr.Name = updatedModel.Name
	}

	if updatedModel.Email != "" && !IsValidEmail(updatedModel.Email) {
		return false, "Enter a valid mail"
	} else {
		usr.Email = updatedModel.Email
	}

	if updatedModel.Password != "" && !IsValidPassword(updatedModel.Password) {
		usr.Password = updatedModel.Password
	}
	return true, ""
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
