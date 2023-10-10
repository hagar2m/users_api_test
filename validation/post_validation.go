package validation

import (
	"com.test.users_api_test/models"
)

func IsValidCreatePost(post models.Post) (bool, string) {
	if !IsValidText(post.Body) || !IsValidText(post.Title) {
		return false, "Enter a valid Content"
	}
	if post.UserID == 0 {
		return false, "Enter user id"
	}
	return true, ""
}
