package validation

import (
	"com.test.users_api_test/models"
)

func IsValidCreatePost(post models.Post) (bool, string) {
	if !IsValidText(post.Body) || !IsValidText(post.Title) {
		return false, "Enter a valid Content"
	}

	return true, ""
}
