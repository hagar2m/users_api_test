package post

import (
	"time"

	"com.test.users_api_test/db"
	"com.test.users_api_test/api/models"
)

func CreatePostQuery(post *models.Post, userId uint) (*models.Post, error) {
	post.UserID = userId
	post.CreatedAt = time.Now()
	result := db.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}
