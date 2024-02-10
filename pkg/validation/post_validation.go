package validation

import (
	"com.test.users_api_test/api/models"
)

func IsValidCreatePost(post models.Post) (bool, string) {
	if !IsValidText(post.Body) || !IsValidTitle(post.Title) {
		return false, "Enter a valid Contentttt"
	}

	return true, ""
}
