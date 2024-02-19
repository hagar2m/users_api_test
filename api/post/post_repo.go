package post

import (
	"time"

	"com.test.users_api_test/api/models"
	"com.test.users_api_test/db"
)

func CreatePostQuery(post *models.Post) (*models.Post, error) {
	post.CreatedAt = time.Now()
	result := db.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}
