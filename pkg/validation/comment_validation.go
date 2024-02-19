package validation

import (
	"com.test.users_api_test/api/models"
)

func IsValidCreateComment(comment models.Comment) (bool, string) {
	if !IsValidText(comment.Body) {
		return false, "Enter a valid Conten"
	}

	return true, ""
}
